package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests
type jsiiProxy_WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumber(value *float64) WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromString(value *string) WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

