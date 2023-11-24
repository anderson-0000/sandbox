package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

