package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesActiveDeadlineSeconds interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesActiveDeadlineSeconds
type jsiiProxy_WorkflowSpecTemplatesActiveDeadlineSeconds struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesActiveDeadlineSeconds) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesActiveDeadlineSeconds_FromNumber(value *float64) WorkflowSpecTemplatesActiveDeadlineSeconds {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesActiveDeadlineSeconds_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesActiveDeadlineSeconds

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesActiveDeadlineSeconds",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesActiveDeadlineSeconds_FromString(value *string) WorkflowSpecTemplatesActiveDeadlineSeconds {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesActiveDeadlineSeconds_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesActiveDeadlineSeconds

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesActiveDeadlineSeconds",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

