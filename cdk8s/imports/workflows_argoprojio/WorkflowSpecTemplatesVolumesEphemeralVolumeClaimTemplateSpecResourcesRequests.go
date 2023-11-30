package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests
type jsiiProxy_WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumber(value *float64) WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromString(value *string) WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

