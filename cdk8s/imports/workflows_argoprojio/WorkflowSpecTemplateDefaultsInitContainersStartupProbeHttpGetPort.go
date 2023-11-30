package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort
type jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort_FromString(value *string) WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

