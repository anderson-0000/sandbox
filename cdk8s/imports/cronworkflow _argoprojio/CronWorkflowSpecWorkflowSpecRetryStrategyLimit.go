package cronworkflow _argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecRetryStrategyLimit interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecRetryStrategyLimit
type jsiiProxy_CronWorkflowSpecWorkflowSpecRetryStrategyLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecRetryStrategyLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecRetryStrategyLimit_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecRetryStrategyLimit {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecRetryStrategyLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecRetryStrategyLimit

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecRetryStrategyLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecRetryStrategyLimit_FromString(value *string) CronWorkflowSpecWorkflowSpecRetryStrategyLimit {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecRetryStrategyLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecRetryStrategyLimit

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecRetryStrategyLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

