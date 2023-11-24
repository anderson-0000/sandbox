package cronworkflow _argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

