package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategyRetries",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

