package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecRetryStrategyLimit interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecRetryStrategyLimit
type jsiiProxy_WorkflowSpecRetryStrategyLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecRetryStrategyLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecRetryStrategyLimit_FromNumber(value *float64) WorkflowSpecRetryStrategyLimit {
	_init_.Initialize()

	if err := validateWorkflowSpecRetryStrategyLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecRetryStrategyLimit

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecRetryStrategyLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecRetryStrategyLimit_FromString(value *string) WorkflowSpecRetryStrategyLimit {
	_init_.Initialize()

	if err := validateWorkflowSpecRetryStrategyLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecRetryStrategyLimit

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecRetryStrategyLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

