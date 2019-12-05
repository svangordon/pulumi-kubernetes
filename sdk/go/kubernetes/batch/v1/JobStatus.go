// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// JobStatus represents the current state of a Job.
type JobStatus struct {
	// The number of actively running pods.
	Active *int `pulumi:"active"`

	// Represents time when the job was completed. It is not guaranteed to be set in happens-before
	// order across separate operations. It is represented in RFC3339 form and is in UTC.
	CompletionTime *string `pulumi:"completionTime"`

	// The latest available observations of an object's current state. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Conditions []JobCondition `pulumi:"conditions"`

	// The number of pods which reached phase Failed.
	Failed *int `pulumi:"failed"`

	// Represents time when the job was acknowledged by the job controller. It is not guaranteed to be
	// set in happens-before order across separate operations. It is represented in RFC3339 form and is
	// in UTC.
	StartTime *string `pulumi:"startTime"`

	// The number of pods which reached phase Succeeded.
	Succeeded *int `pulumi:"succeeded"`

}

var _JobStatusType = reflect.TypeOf((*JobStatus)(nil)).Elem()

// JobStatusInput represents an input type that resolves to a JobStatus.
type JobStatusInput interface {
	ElementType() reflect.Type

	ToJobStatusOutput() JobStatusOutput
	ToJobStatusOutputWithContext(ctx context.Context) JobStatusOutput
}

// JobStatusArgs is a JobStatusInput whose fields are all Input types.
type JobStatusArgs struct {
	// The number of actively running pods.
	Active pulumi.IntInput `pulumi:"active"`

	// Represents time when the job was completed. It is not guaranteed to be set in happens-before
	// order across separate operations. It is represented in RFC3339 form and is in UTC.
	CompletionTime pulumi.StringInput `pulumi:"completionTime"`

	// The latest available observations of an object's current state. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Conditions JobConditionArrayInput `pulumi:"conditions"`

	// The number of pods which reached phase Failed.
	Failed pulumi.IntInput `pulumi:"failed"`

	// Represents time when the job was acknowledged by the job controller. It is not guaranteed to be
	// set in happens-before order across separate operations. It is represented in RFC3339 form and is
	// in UTC.
	StartTime pulumi.StringInput `pulumi:"startTime"`

	// The number of pods which reached phase Succeeded.
	Succeeded pulumi.IntInput `pulumi:"succeeded"`

}

func (a JobStatusArgs) ElementType() reflect.Type {
	return _JobStatusType
}

func (a JobStatusArgs) ToJobStatusOutput() JobStatusOutput {
	return pulumi.ToOutput(a).(JobStatusOutput)
}

func (a JobStatusArgs) ToJobStatusOutputWithContext(ctx context.Context) JobStatusOutput {
	return pulumi.ToOutputWithContext(ctx, a).(JobStatusOutput)
}

// JobStatusOutput is an output type that resolves to a Input.
type JobStatusOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(JobStatusOutput{}) }

func (JobStatusOutput) ElementType() reflect.Type {
	return _JobStatusType
}

func (o JobStatusOutput) Active() pulumi.IntOutput {
	return o.Apply(func(v JobStatus) *int {
		return v.Active
	}).(pulumi.IntOutput)
}

func (o JobStatusOutput) CompletionTime() pulumi.StringOutput {
	return o.Apply(func(v JobStatus) *string {
		return v.CompletionTime
	}).(pulumi.StringOutput)
}

func (o JobStatusOutput) Conditions() JobConditionArrayOutput {
	return o.Apply(func(v JobStatus) []JobCondition {
		return v.Conditions
	}).(JobConditionArrayOutput)
}

func (o JobStatusOutput) Failed() pulumi.IntOutput {
	return o.Apply(func(v JobStatus) *int {
		return v.Failed
	}).(pulumi.IntOutput)
}

func (o JobStatusOutput) StartTime() pulumi.StringOutput {
	return o.Apply(func(v JobStatus) *string {
		return v.StartTime
	}).(pulumi.StringOutput)
}

func (o JobStatusOutput) Succeeded() pulumi.IntOutput {
	return o.Apply(func(v JobStatus) *int {
		return v.Succeeded
	}).(pulumi.IntOutput)
}
