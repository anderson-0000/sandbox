package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

