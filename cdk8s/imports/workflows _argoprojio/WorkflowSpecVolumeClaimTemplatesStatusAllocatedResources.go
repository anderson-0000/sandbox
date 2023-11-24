package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources
type jsiiProxy_WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources_FromNumber(value *float64) WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources_FromString(value *string) WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

