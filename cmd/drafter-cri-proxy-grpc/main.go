package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/loopholelabs/drafter-cri/pkg/services"
	"github.com/loopholelabs/drafter-cri/pkg/utils"
	"github.com/mdlayher/vsock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	v1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

const (
	SchemeUnix  = "unix"
	SchemeVSock = "vsock"
)

var (
	ErrUnsupportedScheme = errors.New("unsupported scheme")
)

func main() {
	rawUpstreamURL := flag.String("upstream-url", "unix:///tmp/upstream.sock", "Upstream URL to listen on (formatted as unix://<path> or vsock://<cid>:<port>)")
	rawDownstreamURL := flag.String("downstream-url", "unix:///tmp/downstream.sock", "Downstream URL to dial (formatted as unix://<path> or vsock://localhost:<port>/<path>:)")

	flag.Parse()

	upstreamURL, err := url.Parse(*rawUpstreamURL)
	if err != nil {
		panic(err)
	}

	downstreamURL, err := url.Parse(*rawDownstreamURL)
	if err != nil {
		panic(err)
	}

	var upstreamLis net.Listener
	switch upstreamURL.Scheme {
	case SchemeUnix:
		_ = os.Remove(upstreamURL.Path)

		upstreamLis, err = net.Listen("unix", upstreamURL.Path)
		if err != nil {
			panic(err)
		}

	case SchemeVSock:
		cid, err := strconv.Atoi(upstreamURL.Hostname())
		if err != nil {
			panic(err)
		}

		port, err := strconv.Atoi(upstreamURL.Port())
		if err != nil {
			panic(err)
		}

		upstreamLis, err = vsock.ListenContextID(uint32(cid), uint32(port), nil)
		if err != nil {
			panic(err)
		}

	default:
		panic(ErrUnsupportedScheme)
	}
	defer upstreamLis.Close()

	downstreamConn, err := grpc.Dial(
		downstreamURL.Path,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			var rawDownstreamConn net.Conn
			switch downstreamURL.Scheme {
			case SchemeUnix:
				rawDownstreamConn, err = net.Dial("unix", downstreamURL.Path)
				if err != nil {
					return nil, err
				}

			case SchemeVSock:
				rawDownstreamConn, err = net.Dial("unix", downstreamURL.Path)
				if err != nil {
					return nil, err
				}

				port, err := strconv.Atoi(downstreamURL.Port())
				if err != nil {
					return nil, err
				}

				if _, err = rawDownstreamConn.Write([]byte(fmt.Sprintf("CONNECT %d\n", port))); err != nil {
					return nil, err
				}

				line, err := utils.ReadLineNoBuffer(rawDownstreamConn)
				if err != nil {
					return nil, err
				}

				if !strings.HasPrefix(line, "OK ") {
					return nil, errors.New("could not connect to VSock")
				}

			default:
				return nil, ErrUnsupportedScheme
			}

			return rawDownstreamConn, err
		}),
	)
	if err != nil {
		panic(err)
	}

	downstreamImageServiceClient := v1.NewImageServiceClient(downstreamConn)
	downstreamRuntimeServiceClient := v1.NewRuntimeServiceClient(downstreamConn)

	server := grpc.NewServer()
	v1.RegisterImageServiceServer(server, services.NewImageGrpc(downstreamImageServiceClient))
	v1.RegisterRuntimeServiceServer(server, services.NewRuntimeGrpc(downstreamRuntimeServiceClient))

	log.Println("Proxying requests from", upstreamURL, "to", downstreamURL)

	if err := server.Serve(upstreamLis); err != nil {
		panic(err)
	}
}
