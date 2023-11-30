package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

