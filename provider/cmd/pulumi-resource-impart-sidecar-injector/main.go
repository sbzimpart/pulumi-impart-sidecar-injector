//go:generate go run ./generate.go

package main

import (
	"github.com/impart-security/pulumi-impart-sidecar-injector/pkg/provider"
	"github.com/impart-security/pulumi-impart-sidecar-injector/pkg/version"
)

// go:embed schema-embed.json
// var pulumiSchema []byte

func main() {
	provider.Serve(version.Version, pulumiSchema)
}
