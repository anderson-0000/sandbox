package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount
type jsiiProxy_WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount_FromNumber(value *float64) WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsDagTasksWithSequenceCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount_FromString(value *string) WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsDagTasksWithSequenceCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

