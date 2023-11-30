package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecVolumeClaimTemplatesStatusCapacity interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecVolumeClaimTemplatesStatusCapacity
type jsiiProxy_WorkflowSpecVolumeClaimTemplatesStatusCapacity struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecVolumeClaimTemplatesStatusCapacity) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecVolumeClaimTemplatesStatusCapacity_FromNumber(value *float64) WorkflowSpecVolumeClaimTemplatesStatusCapacity {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumeClaimTemplatesStatusCapacity_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumeClaimTemplatesStatusCapacity

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecVolumeClaimTemplatesStatusCapacity",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecVolumeClaimTemplatesStatusCapacity_FromString(value *string) WorkflowSpecVolumeClaimTemplatesStatusCapacity {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumeClaimTemplatesStatusCapacity_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumeClaimTemplatesStatusCapacity

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecVolumeClaimTemplatesStatusCapacity",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

