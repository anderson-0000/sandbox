package cronworkflow _argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

