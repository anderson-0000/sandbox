package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort
type jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort_FromNumber(value *float64) WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort_FromString(value *string) WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

