package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

