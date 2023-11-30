package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

