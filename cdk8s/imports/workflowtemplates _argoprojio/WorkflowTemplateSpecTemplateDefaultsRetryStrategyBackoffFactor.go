package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

