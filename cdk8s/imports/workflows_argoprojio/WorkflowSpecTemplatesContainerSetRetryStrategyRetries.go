package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesContainerSetRetryStrategyRetries interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesContainerSetRetryStrategyRetries
type jsiiProxy_WorkflowSpecTemplatesContainerSetRetryStrategyRetries struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesContainerSetRetryStrategyRetries) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesContainerSetRetryStrategyRetries_FromNumber(value *float64) WorkflowSpecTemplatesContainerSetRetryStrategyRetries {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetRetryStrategyRetries_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetRetryStrategyRetries

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerSetRetryStrategyRetries",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesContainerSetRetryStrategyRetries_FromString(value *string) WorkflowSpecTemplatesContainerSetRetryStrategyRetries {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetRetryStrategyRetries_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetRetryStrategyRetries

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerSetRetryStrategyRetries",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

