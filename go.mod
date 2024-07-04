module github.com/loopholelabs/drafter-cri

go 1.21.6

require (
	github.com/mdlayher/vsock v1.2.1
	google.golang.org/grpc v1.60.1
	k8s.io/cri-api v0.29.2
)

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/mdlayher/socket v0.4.1 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)

replace github.com/loopholelabs/drafter => ../drafter

replace github.com/pojntfx/panrpc/go => ../panrpc/go

replace github.com/pojntfx/r3map => ../r3map
