package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort
type jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort_FromString(value *string) WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

