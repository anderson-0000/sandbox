package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsDagTasksWithSequenceStart",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

