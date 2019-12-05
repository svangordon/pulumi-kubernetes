// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Describe a container image
type ContainerImage struct {
	// Names by which this image is known. e.g. ["k8s.gcr.io/hyperkube:v1.0.7",
	// "dockerhub.io/google_containers/hyperkube:v1.0.7"]
	Names []string `pulumi:"names"`

	// The size of the image in bytes.
	SizeBytes *int `pulumi:"sizeBytes"`

}

var _ContainerImageType = reflect.TypeOf((*ContainerImage)(nil)).Elem()

// ContainerImageInput represents an input type that resolves to a ContainerImage.
type ContainerImageInput interface {
	ElementType() reflect.Type

	ToContainerImageOutput() ContainerImageOutput
	ToContainerImageOutputWithContext(ctx context.Context) ContainerImageOutput
}

// ContainerImageArgs is a ContainerImageInput whose fields are all Input types.
type ContainerImageArgs struct {
	// Names by which this image is known. e.g. ["k8s.gcr.io/hyperkube:v1.0.7",
	// "dockerhub.io/google_containers/hyperkube:v1.0.7"]
	Names pulumi.StringArrayInput `pulumi:"names"`

	// The size of the image in bytes.
	SizeBytes pulumi.IntInput `pulumi:"sizeBytes"`

}

func (a ContainerImageArgs) ElementType() reflect.Type {
	return _ContainerImageType
}

func (a ContainerImageArgs) ToContainerImageOutput() ContainerImageOutput {
	return pulumi.ToOutput(a).(ContainerImageOutput)
}

func (a ContainerImageArgs) ToContainerImageOutputWithContext(ctx context.Context) ContainerImageOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ContainerImageOutput)
}

// ContainerImageOutput is an output type that resolves to a Input.
type ContainerImageOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ContainerImageOutput{}) }

func (ContainerImageOutput) ElementType() reflect.Type {
	return _ContainerImageType
}

func (o ContainerImageOutput) Names() pulumi.StringArrayOutput {
	return o.Apply(func(v ContainerImage) []string {
		return v.Names
	}).(pulumi.StringArrayOutput)
}

func (o ContainerImageOutput) SizeBytes() pulumi.IntOutput {
	return o.Apply(func(v ContainerImage) *int {
		return v.SizeBytes
	}).(pulumi.IntOutput)
}

var _ContainerImageArrayType = reflect.TypeOf((*[]ContainerImage)(nil)).Elem()

type ContainerImageArrayInput interface {
	ElementType() reflect.Type

	ToContainerImageArrayOutput() ContainerImageArrayOutput
	ToContainerImageArrayOutputWithContext(ctx context.Context) ContainerImageArrayOutput
}

type ContainerImageArray []ContainerImageInput

func (a ContainerImageArray) ElementType() reflect.Type {
	return _ContainerImageArrayType
}

func (a ContainerImageArray) ToContainerImageArrayOutput() ContainerImageArrayOutput {
	return pulumi.ToOutput(a).(ContainerImageArrayOutput)
}

func (a ContainerImageArray) ToContainerImageArrayOutputWithContext(ctx context.Context) ContainerImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ContainerImageArrayOutput)
}

type ContainerImageArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ContainerImageArrayOutput{}) }

func (ContainerImageArrayOutput) ElementType() reflect.Type {
	return _ContainerImageArrayType
}

func (o ContainerImageArrayOutput) Index(i pulumi.IntInput) ContainerImageOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) ContainerImage {
		return vs[0].([]ContainerImage)[vs[1].(int)]
	}).(ContainerImageOutput)
}