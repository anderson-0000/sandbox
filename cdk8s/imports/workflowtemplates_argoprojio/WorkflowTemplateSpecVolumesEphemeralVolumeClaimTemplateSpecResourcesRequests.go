package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests
type jsiiProxy_WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumber(value *float64) WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromString(value *string) WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

