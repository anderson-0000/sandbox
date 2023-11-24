package cronworkflow _argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor
type jsiiProxy_CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor_FromString(value *string) CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

