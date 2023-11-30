package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort
type jsiiProxy_WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesScriptReadinessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort_FromString(value *string) WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesScriptReadinessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

