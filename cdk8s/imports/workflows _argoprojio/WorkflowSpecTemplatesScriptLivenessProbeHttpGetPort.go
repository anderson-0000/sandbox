package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort
type jsiiProxy_WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort_FromString(value *string) WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesScriptLivenessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

