package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort
type jsiiProxy_WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort_FromString(value *string) WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

