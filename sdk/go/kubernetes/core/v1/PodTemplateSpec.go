// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// PodTemplateSpec describes the data a pod should have when created from a template
type PodTemplateSpec struct {
	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

	// Specification of the desired behavior of the pod. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *PodSpec `pulumi:"spec"`

}

var _PodTemplateSpecType = reflect.TypeOf((*PodTemplateSpec)(nil)).Elem()

// PodTemplateSpecInput represents an input type that resolves to a PodTemplateSpec.
type PodTemplateSpecInput interface {
	ElementType() reflect.Type

	ToPodTemplateSpecOutput() PodTemplateSpecOutput
	ToPodTemplateSpecOutputWithContext(ctx context.Context) PodTemplateSpecOutput
}

// PodTemplateSpecArgs is a PodTemplateSpecInput whose fields are all Input types.
type PodTemplateSpecArgs struct {
	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Specification of the desired behavior of the pod. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec PodSpecInput `pulumi:"spec"`

}

func (a PodTemplateSpecArgs) ElementType() reflect.Type {
	return _PodTemplateSpecType
}

func (a PodTemplateSpecArgs) ToPodTemplateSpecOutput() PodTemplateSpecOutput {
	return pulumi.ToOutput(a).(PodTemplateSpecOutput)
}

func (a PodTemplateSpecArgs) ToPodTemplateSpecOutputWithContext(ctx context.Context) PodTemplateSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PodTemplateSpecOutput)
}

// PodTemplateSpecOutput is an output type that resolves to a Input.
type PodTemplateSpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(PodTemplateSpecOutput{}) }

func (PodTemplateSpecOutput) ElementType() reflect.Type {
	return _PodTemplateSpecType
}

func (o PodTemplateSpecOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v PodTemplateSpec) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o PodTemplateSpecOutput) Spec() PodSpecOutput {
	return o.Apply(func(v PodTemplateSpec) *PodSpec {
		return v.Spec
	}).(PodSpecOutput)
}
