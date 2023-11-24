package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesRetryStrategyBackoffFactor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesRetryStrategyBackoffFactor
type jsiiProxy_WorkflowSpecTemplatesRetryStrategyBackoffFactor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesRetryStrategyBackoffFactor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesRetryStrategyBackoffFactor_FromNumber(value *float64) WorkflowSpecTemplatesRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesRetryStrategyBackoffFactor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesRetryStrategyBackoffFactor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesRetryStrategyBackoffFactor_FromString(value *string) WorkflowSpecTemplatesRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesRetryStrategyBackoffFactor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesRetryStrategyBackoffFactor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

