package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsRetryStrategyLimit interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsRetryStrategyLimit
type jsiiProxy_WorkflowSpecTemplateDefaultsRetryStrategyLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsRetryStrategyLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsRetryStrategyLimit_FromNumber(value *float64) WorkflowSpecTemplateDefaultsRetryStrategyLimit {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsRetryStrategyLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsRetryStrategyLimit

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsRetryStrategyLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsRetryStrategyLimit_FromString(value *string) WorkflowSpecTemplateDefaultsRetryStrategyLimit {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsRetryStrategyLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsRetryStrategyLimit

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsRetryStrategyLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

