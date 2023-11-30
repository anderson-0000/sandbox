package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits
type jsiiProxy_WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromNumber(value *float64) WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromString(value *string) WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

