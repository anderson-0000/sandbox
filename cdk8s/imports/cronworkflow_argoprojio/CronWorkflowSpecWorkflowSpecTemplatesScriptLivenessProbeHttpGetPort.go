package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

