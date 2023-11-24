package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor
type jsiiProxy_WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor_FromNumber(value *float64) WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor_FromString(value *string) WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

