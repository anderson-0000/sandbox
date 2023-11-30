package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

