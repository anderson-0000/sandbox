package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests
type jsiiProxy_WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumber(value *float64) WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromString(value *string) WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

