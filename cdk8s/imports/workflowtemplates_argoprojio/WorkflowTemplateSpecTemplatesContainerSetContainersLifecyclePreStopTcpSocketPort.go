package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort
type jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort_FromNumber(value *float64) WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort_FromString(value *string) WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

