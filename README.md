# Versioned Go Control Plane

This repository creates versioned releases of Envoy’s Go Control Plane pinned to specific Envoy versions. For each selected Go Control Plane version (e.g., v0.13.4), we generate protos from a chosen Envoy version (e.g., v1.32.3) and publish the resulting code as a combined release (v0.13.4-envoy1.32.3). This ensures that each release is tied to an exact Envoy version, making it easier to maintain stability and compatibility across different environments.

## How to use

```
go get github.com/sefaphlvn/versioned-control-plane@v0.13.4-envoy1.32.3
```

or

```
go get github.com/sefaphlvn/versioned-control-plane@v0.13.4-envoy1.33.0
```
