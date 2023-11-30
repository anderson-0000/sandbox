package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort
type jsiiProxy_WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort_FromString(value *string) WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

