module github.com/sefaphlvn/versioned-go-control-plane/contrib

go 1.22.0

toolchain go1.23.4

require (
	github.com/cncf/xds/go v0.0.0-20250121191232-2f005788dc42
	github.com/envoyproxy/protoc-gen-validate v1.2.0
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10
	github.com/sefaphlvn/versioned-go-control-plane/envoy v1.32.3
	google.golang.org/grpc v1.69.4
	google.golang.org/protobuf v1.36.3
)

require (
	cel.dev/expr v0.16.2 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250122153221-138b5a5a4fd4 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
)

replace github.com/sefaphlvn/versioned-go-control-plane/envoy => ../envoy
