package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort
type jsiiProxy_WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort_FromString(value *string) WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

