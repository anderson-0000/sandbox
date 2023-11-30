package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesDagTasksWithSequenceStart interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesDagTasksWithSequenceStart
type jsiiProxy_WorkflowSpecTemplatesDagTasksWithSequenceStart struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesDagTasksWithSequenceStart) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesDagTasksWithSequenceStart_FromNumber(value *float64) WorkflowSpecTemplatesDagTasksWithSequenceStart {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesDagTasksWithSequenceStart_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesDagTasksWithSequenceStart

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesDagTasksWithSequenceStart",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesDagTasksWithSequenceStart_FromString(value *string) WorkflowSpecTemplatesDagTasksWithSequenceStart {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesDagTasksWithSequenceStart_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesDagTasksWithSequenceStart

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesDagTasksWithSequenceStart",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

