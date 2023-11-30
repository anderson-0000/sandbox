package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor
type jsiiProxy_WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor_FromString(value *string) WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

