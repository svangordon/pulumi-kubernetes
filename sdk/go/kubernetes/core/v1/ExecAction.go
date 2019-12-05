// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ExecAction describes a "run in container" action.
type ExecAction struct {
	// Command is the command line to execute inside the container, the working directory for the
	// command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not
	// run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you
	// need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and
	// non-zero is unhealthy.
	Command []string `pulumi:"command"`

}

var _ExecActionType = reflect.TypeOf((*ExecAction)(nil)).Elem()

// ExecActionInput represents an input type that resolves to a ExecAction.
type ExecActionInput interface {
	ElementType() reflect.Type

	ToExecActionOutput() ExecActionOutput
	ToExecActionOutputWithContext(ctx context.Context) ExecActionOutput
}

// ExecActionArgs is a ExecActionInput whose fields are all Input types.
type ExecActionArgs struct {
	// Command is the command line to execute inside the container, the working directory for the
	// command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not
	// run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you
	// need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and
	// non-zero is unhealthy.
	Command pulumi.StringArrayInput `pulumi:"command"`

}

func (a ExecActionArgs) ElementType() reflect.Type {
	return _ExecActionType
}

func (a ExecActionArgs) ToExecActionOutput() ExecActionOutput {
	return pulumi.ToOutput(a).(ExecActionOutput)
}

func (a ExecActionArgs) ToExecActionOutputWithContext(ctx context.Context) ExecActionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ExecActionOutput)
}

// ExecActionOutput is an output type that resolves to a Input.
type ExecActionOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ExecActionOutput{}) }

func (ExecActionOutput) ElementType() reflect.Type {
	return _ExecActionType
}

func (o ExecActionOutput) Command() pulumi.StringArrayOutput {
	return o.Apply(func(v ExecAction) []string {
		return v.Command
	}).(pulumi.StringArrayOutput)
}
