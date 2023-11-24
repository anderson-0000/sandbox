package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

