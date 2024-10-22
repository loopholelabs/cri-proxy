package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/loopholelabs/drafter-cri/pkg/utils"
	"github.com/mdlayher/vsock"
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

	log.Println("Proxying connections from", upstreamURL, "to", downstreamURL)

	clients := 0
	for {
		func() {
			upstreamConn, err := upstreamLis.Accept()
			if err != nil {
				log.Println("could not accept connection, continuing:", err)

				return
			}

			go func() {
				clients++

				log.Printf("%v clients connected", clients)

				defer func() {
					_ = upstreamConn.Close()

					if err := recover(); err != nil {
						log.Printf("Client disconnected with error: %v", err)
					}

					clients--

					log.Printf("%v clients connected", clients)
				}()

				var downstreamConn net.Conn
				switch downstreamURL.Scheme {
				case SchemeUnix:
					downstreamConn, err = net.Dial("unix", downstreamURL.Path)
					if err != nil {
						panic(err)
					}

				case SchemeVSock:
					downstreamConn, err = net.Dial("unix", downstreamURL.Path)
					if err != nil {
						panic(err)
					}

					port, err := strconv.Atoi(downstreamURL.Port())
					if err != nil {
						panic(err)
					}

					if _, err = downstreamConn.Write([]byte(fmt.Sprintf("CONNECT %d\n", port))); err != nil {
						panic(err)
					}

					line, err := utils.ReadLineNoBuffer(downstreamConn)
					if err != nil {
						panic(err)
					}

					if !strings.HasPrefix(line, "OK ") {
						panic(errors.New("could not connect to VSock"))
					}

				default:
					panic(ErrUnsupportedScheme)
				}
				defer downstreamConn.Close()

				var copyErr error
				defer func() {
					if err := copyErr; err != nil {
						panic(err)
					}
				}()

				copyDone := make(chan struct{})
				setCopyDone := sync.OnceFunc(func() {
					close(copyDone)
				})

				go func() {
					defer setCopyDone()

					if _, err := io.Copy(downstreamConn, upstreamConn); err != nil {
						copyErr = err
					}
				}()

				go func() {
					defer setCopyDone()

					if _, err := io.Copy(upstreamConn, downstreamConn); err != nil {
						copyErr = err
					}
				}()

				<-copyDone
			}()
		}()
	}
}
