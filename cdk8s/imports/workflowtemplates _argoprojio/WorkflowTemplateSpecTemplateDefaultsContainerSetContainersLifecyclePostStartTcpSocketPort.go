package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

