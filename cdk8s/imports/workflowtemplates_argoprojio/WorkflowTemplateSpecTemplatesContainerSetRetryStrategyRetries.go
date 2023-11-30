package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries
type jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries_FromNumber(value *float64) WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries_FromString(value *string) WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

