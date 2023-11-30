package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits
type jsiiProxy_WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits_FromNumber(value *float64) WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits_FromString(value *string) WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

