package cronworkflow _argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceStart",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

