module github.com/sefaphlvn/versioned-go-control-plane/ratelimit

go 1.22

require (
  github.com/sefaphlvn/versioned-go-control-plane/envoy v1.33.0
  google.golang.org/protobuf v1.36.3
)

replace github.com/sefaphlvn/versioned-go-control-plane/envoy => ../envoy
