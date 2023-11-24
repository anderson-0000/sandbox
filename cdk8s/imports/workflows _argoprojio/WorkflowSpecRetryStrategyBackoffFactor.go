package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecRetryStrategyBackoffFactor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecRetryStrategyBackoffFactor
type jsiiProxy_WorkflowSpecRetryStrategyBackoffFactor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecRetryStrategyBackoffFactor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecRetryStrategyBackoffFactor_FromNumber(value *float64) WorkflowSpecRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateWorkflowSpecRetryStrategyBackoffFactor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecRetryStrategyBackoffFactor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecRetryStrategyBackoffFactor_FromString(value *string) WorkflowSpecRetryStrategyBackoffFactor {
	_init_.Initialize()

	if err := validateWorkflowSpecRetryStrategyBackoffFactor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecRetryStrategyBackoffFactor

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecRetryStrategyBackoffFactor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

