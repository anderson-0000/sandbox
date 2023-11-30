package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort
type jsiiProxy_WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort_FromNumber(value *float64) WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort_FromString(value *string) WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

