package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort
type jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort_FromNumber(value *float64) WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort_FromString(value *string) WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

