// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// Namespace provides a scope for Names. Use of multiple namespaces is optional.
type Namespace struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// Spec defines the behavior of the Namespace. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec NamespaceSpecOutput `pulumi:"spec"`

	// Status describes the current status of a Namespace. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status NamespaceStatusOutput `pulumi:"status"`

}

// NamespaceArgs is the set of arguments needed to create a Namespace resource.
type NamespaceArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Spec defines the behavior of the Namespace. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec NamespaceSpecInput `pulumi:"spec"`

}

// NewNamespace creates a Namespace resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context, name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		args.ApiVersion = pulumi.String("v1")
		args.Kind = pulumi.String("Namespace")
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToObjectMetaOutput()
		}
		if i := args.Spec; i != nil {
			inputs["spec"] = i.ToNamespaceSpecOutput()
		}
	}
	var resource Namespace
	err := ctx.RegisterResource("kubernetes:core/v1:Namespace", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name and ID.
func GetNamespace(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("kubernetes:core/v1:Namespace", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Namespace provides a scope for Names. Use of multiple namespaces is optional.
type NamespaceProperty struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

	// Spec defines the behavior of the Namespace. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *NamespaceSpec `pulumi:"spec"`

	// Status describes the current status of a Namespace. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status NamespaceStatus `pulumi:"status"`

}

var _NamespacePropertyType = reflect.TypeOf((*NamespaceProperty)(nil)).Elem()

// NamespacePropertyInput represents an input type that resolves to a NamespaceProperty.
type NamespacePropertyInput interface {
	ElementType() reflect.Type

	ToNamespacePropertyOutput() NamespacePropertyOutput
	ToNamespacePropertyOutputWithContext(ctx context.Context) NamespacePropertyOutput
}

// NamespacePropertyArgs is a NamespacePropertyInput whose fields are all Input types.
type NamespacePropertyArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Spec defines the behavior of the Namespace. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec NamespaceSpecInput `pulumi:"spec"`

}

func (a NamespacePropertyArgs) ElementType() reflect.Type {
	return _NamespacePropertyType
}

func (a NamespacePropertyArgs) ToNamespacePropertyOutput() NamespacePropertyOutput {
	return pulumi.ToOutput(a).(NamespacePropertyOutput)
}

func (a NamespacePropertyArgs) ToNamespacePropertyOutputWithContext(ctx context.Context) NamespacePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NamespacePropertyOutput)
}

// NamespacePropertyOutput is an output type that resolves to a Input.
type NamespacePropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NamespacePropertyOutput{}) }

func (NamespacePropertyOutput) ElementType() reflect.Type {
	return _NamespacePropertyType
}

func (o NamespacePropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v NamespaceProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o NamespacePropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v NamespaceProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o NamespacePropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v NamespaceProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o NamespacePropertyOutput) Spec() NamespaceSpecOutput {
	return o.Apply(func(v NamespaceProperty) *NamespaceSpec {
		return v.Spec
	}).(NamespaceSpecOutput)
}

func (o NamespacePropertyOutput) Status() NamespaceStatusOutput {
	return o.Apply(func(v NamespaceProperty) NamespaceStatus {
		return v.Status
	}).(NamespaceStatusOutput)
}

var _NamespacePropertyArrayType = reflect.TypeOf((*[]NamespaceProperty)(nil)).Elem()

type NamespacePropertyArrayInput interface {
	ElementType() reflect.Type

	ToNamespacePropertyArrayOutput() NamespacePropertyArrayOutput
	ToNamespacePropertyArrayOutputWithContext(ctx context.Context) NamespacePropertyArrayOutput
}

type NamespacePropertyArray []NamespacePropertyInput

func (a NamespacePropertyArray) ElementType() reflect.Type {
	return _NamespacePropertyArrayType
}

func (a NamespacePropertyArray) ToNamespacePropertyArrayOutput() NamespacePropertyArrayOutput {
	return pulumi.ToOutput(a).(NamespacePropertyArrayOutput)
}

func (a NamespacePropertyArray) ToNamespacePropertyArrayOutputWithContext(ctx context.Context) NamespacePropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NamespacePropertyArrayOutput)
}

type NamespacePropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NamespacePropertyArrayOutput{}) }

func (NamespacePropertyArrayOutput) ElementType() reflect.Type {
	return _NamespacePropertyArrayType
}

func (o NamespacePropertyArrayOutput) Index(i pulumi.IntInput) NamespacePropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) NamespaceProperty {
		return vs[0].([]NamespaceProperty)[vs[1].(int)]
	}).(NamespacePropertyOutput)
}