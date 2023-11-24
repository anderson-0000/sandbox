package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesRetryStrategyLimit interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesRetryStrategyLimit
type jsiiProxy_WorkflowTemplateSpecTemplatesRetryStrategyLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesRetryStrategyLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesRetryStrategyLimit_FromNumber(value *float64) WorkflowTemplateSpecTemplatesRetryStrategyLimit {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesRetryStrategyLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesRetryStrategyLimit

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesRetryStrategyLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesRetryStrategyLimit_FromString(value *string) WorkflowTemplateSpecTemplatesRetryStrategyLimit {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesRetryStrategyLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesRetryStrategyLimit

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesRetryStrategyLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

