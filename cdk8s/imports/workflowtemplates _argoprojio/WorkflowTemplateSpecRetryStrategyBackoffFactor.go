package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecRetryStrategyBackoffFactor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecRetryStrategyBackoffFactor
type jsiiProxy_WorkflowTemplateSpecRetryStrategyBackoffFactor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecRetryStrategyBackoffFactor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecRetryStrategyBackoffFactor_FromNumber(value *float64) WorkflowTemplateSpecRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecRetryStrategyBackoffFactor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecRetryStrategyBackoffFactor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecRetryStrategyBackoffFactor_FromString(value *string) WorkflowTemplateSpecRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecRetryStrategyBackoffFactor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecRetryStrategyBackoffFactor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

