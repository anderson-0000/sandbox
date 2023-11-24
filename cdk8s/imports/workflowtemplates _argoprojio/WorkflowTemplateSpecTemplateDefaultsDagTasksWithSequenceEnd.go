package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

