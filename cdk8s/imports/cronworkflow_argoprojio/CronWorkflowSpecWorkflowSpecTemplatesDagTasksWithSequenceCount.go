package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

