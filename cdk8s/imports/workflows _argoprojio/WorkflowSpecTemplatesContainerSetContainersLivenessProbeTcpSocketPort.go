package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort
type jsiiProxy_WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort_FromString(value *string) WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

