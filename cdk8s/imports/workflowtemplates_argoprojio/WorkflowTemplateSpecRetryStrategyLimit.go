package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecRetryStrategyLimit interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecRetryStrategyLimit
type jsiiProxy_WorkflowTemplateSpecRetryStrategyLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecRetryStrategyLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecRetryStrategyLimit_FromNumber(value *float64) WorkflowTemplateSpecRetryStrategyLimit {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecRetryStrategyLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecRetryStrategyLimit

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecRetryStrategyLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecRetryStrategyLimit_FromString(value *string) WorkflowTemplateSpecRetryStrategyLimit {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecRetryStrategyLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecRetryStrategyLimit

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecRetryStrategyLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

