package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart
type jsiiProxy_WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart_FromNumber(value *float64) WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart_FromString(value *string) WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

