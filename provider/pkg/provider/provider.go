package provider

import (
	helmbase "github.com/pulumi/pulumi-go-helmbase"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pp "github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

const (
	ProviderName  = "impart-sidecar-injector"
	ComponentName = ProviderName + ":index:IngressController"
)

// Serve launches the gRPC server for the resource provider.
func Serve(version string, schema []byte) {
	if err := provider.ComponentMain(ProviderName, version, schema, Construct); err != nil {
		cmdutil.ExitError(err.Error())
	}
}

// Construct is the RPC call that initiates the creation of a new component resource. It
// creates, registers, and returns the resulting object.
func Construct(ctx *pulumi.Context, typ, name string, inputs pp.ConstructInputs,
	opts pulumi.ResourceOption) (*pp.ConstructResult, error) {
	return helmbase.Construct(ctx, &SidecarInjector{}, typ, name, &SidecarInjectorArgs{}, inputs, opts)
}
