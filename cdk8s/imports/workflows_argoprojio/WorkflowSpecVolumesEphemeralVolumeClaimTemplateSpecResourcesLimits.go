package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits
type jsiiProxy_WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromNumber(value *float64) WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromString(value *string) WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

