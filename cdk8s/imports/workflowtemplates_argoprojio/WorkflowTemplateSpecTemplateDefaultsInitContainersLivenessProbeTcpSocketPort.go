package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

