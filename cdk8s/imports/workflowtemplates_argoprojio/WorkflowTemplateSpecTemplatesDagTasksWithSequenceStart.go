package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart
type jsiiProxy_WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart_FromNumber(value *float64) WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesDagTasksWithSequenceStart_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart_FromString(value *string) WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesDagTasksWithSequenceStart_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

