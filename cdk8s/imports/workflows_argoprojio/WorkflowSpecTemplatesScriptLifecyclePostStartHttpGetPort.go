package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort
type jsiiProxy_WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort_FromString(value *string) WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

