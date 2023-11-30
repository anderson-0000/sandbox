package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

