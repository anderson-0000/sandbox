package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount
type jsiiProxy_WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount_FromNumber(value *float64) WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesDagTasksWithSequenceCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount_FromString(value *string) WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesDagTasksWithSequenceCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

