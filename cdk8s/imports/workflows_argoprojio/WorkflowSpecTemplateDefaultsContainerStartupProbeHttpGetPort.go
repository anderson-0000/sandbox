package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort
type jsiiProxy_WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort_FromString(value *string) WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

