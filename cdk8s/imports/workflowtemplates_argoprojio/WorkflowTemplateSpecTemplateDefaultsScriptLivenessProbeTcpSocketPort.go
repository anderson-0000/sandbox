package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

