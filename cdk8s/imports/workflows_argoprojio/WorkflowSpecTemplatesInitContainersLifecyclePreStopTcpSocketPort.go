package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort
type jsiiProxy_WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort_FromString(value *string) WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

