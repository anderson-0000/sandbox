package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

