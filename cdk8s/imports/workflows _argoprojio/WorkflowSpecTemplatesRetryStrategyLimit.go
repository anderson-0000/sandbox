package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesRetryStrategyLimit interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesRetryStrategyLimit
type jsiiProxy_WorkflowSpecTemplatesRetryStrategyLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesRetryStrategyLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesRetryStrategyLimit_FromNumber(value *float64) WorkflowSpecTemplatesRetryStrategyLimit {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesRetryStrategyLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesRetryStrategyLimit

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesRetryStrategyLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesRetryStrategyLimit_FromString(value *string) WorkflowSpecTemplatesRetryStrategyLimit {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesRetryStrategyLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesRetryStrategyLimit

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesRetryStrategyLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

