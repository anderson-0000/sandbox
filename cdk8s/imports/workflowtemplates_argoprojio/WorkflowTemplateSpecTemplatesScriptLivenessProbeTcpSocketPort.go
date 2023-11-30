package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort
type jsiiProxy_WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort_FromNumber(value *float64) WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort_FromString(value *string) WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

