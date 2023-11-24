package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd
type jsiiProxy_WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd_FromNumber(value *float64) WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd_FromString(value *string) WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

