package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

