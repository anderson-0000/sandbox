package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort
type jsiiProxy_WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort_FromString(value *string) WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

