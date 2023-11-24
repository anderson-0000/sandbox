package cronworkflow _argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesDagTasksWithSequenceEnd",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

