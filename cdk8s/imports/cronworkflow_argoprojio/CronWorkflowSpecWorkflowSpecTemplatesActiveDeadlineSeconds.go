package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesActiveDeadlineSeconds",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

