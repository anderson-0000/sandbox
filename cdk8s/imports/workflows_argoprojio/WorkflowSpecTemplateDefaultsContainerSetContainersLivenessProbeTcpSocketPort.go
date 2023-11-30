package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort
type jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort_FromString(value *string) WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

