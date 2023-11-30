package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort
type jsiiProxy_WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort_FromString(value *string) WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

