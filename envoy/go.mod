module github.com/sefaphlvn/versioned-go-control-plane/envoy

go 1.22.0

toolchain go1.23.4

require (
	github.com/cncf/xds/go v0.0.0-20250121191232-2f005788dc42
	github.com/envoyproxy/protoc-gen-validate v1.2.0
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10
	github.com/prometheus/client_model v0.6.1
	go.opentelemetry.io/proto/otlp v1.5.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250122153221-138b5a5a4fd4
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f
	google.golang.org/grpc v1.69.4
	google.golang.org/protobuf v1.36.3
)

require (
	cel.dev/expr v0.16.2 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
