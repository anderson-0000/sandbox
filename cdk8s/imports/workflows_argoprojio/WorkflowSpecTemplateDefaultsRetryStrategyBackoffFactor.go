package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor
type jsiiProxy_WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor_FromNumber(value *float64) WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor_FromString(value *string) WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

