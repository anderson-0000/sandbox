package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort
type jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort_FromString(value *string) WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

