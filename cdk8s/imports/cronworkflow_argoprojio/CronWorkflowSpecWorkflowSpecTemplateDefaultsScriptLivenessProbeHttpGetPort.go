package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

