package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor
type jsiiProxy_WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor_FromString(value *string) WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

