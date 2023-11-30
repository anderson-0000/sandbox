package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesDagTasksWithSequenceCount interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesDagTasksWithSequenceCount
type jsiiProxy_WorkflowSpecTemplatesDagTasksWithSequenceCount struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesDagTasksWithSequenceCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesDagTasksWithSequenceCount_FromNumber(value *float64) WorkflowSpecTemplatesDagTasksWithSequenceCount {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesDagTasksWithSequenceCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesDagTasksWithSequenceCount

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesDagTasksWithSequenceCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesDagTasksWithSequenceCount_FromString(value *string) WorkflowSpecTemplatesDagTasksWithSequenceCount {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesDagTasksWithSequenceCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesDagTasksWithSequenceCount

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesDagTasksWithSequenceCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

