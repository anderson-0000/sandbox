package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesDagTasksWithSequenceEnd interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesDagTasksWithSequenceEnd
type jsiiProxy_WorkflowSpecTemplatesDagTasksWithSequenceEnd struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesDagTasksWithSequenceEnd) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesDagTasksWithSequenceEnd_FromNumber(value *float64) WorkflowSpecTemplatesDagTasksWithSequenceEnd {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesDagTasksWithSequenceEnd_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesDagTasksWithSequenceEnd

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksWithSequenceEnd",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesDagTasksWithSequenceEnd_FromString(value *string) WorkflowSpecTemplatesDagTasksWithSequenceEnd {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesDagTasksWithSequenceEnd_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesDagTasksWithSequenceEnd

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksWithSequenceEnd",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

