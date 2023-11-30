package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd
type jsiiProxy_WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd_FromNumber(value *float64) WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd_FromString(value *string) WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

