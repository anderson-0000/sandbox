package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort
type jsiiProxy_WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort_FromNumber(value *float64) WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort_FromString(value *string) WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

