package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort
type jsiiProxy_WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort_FromString(value *string) WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

