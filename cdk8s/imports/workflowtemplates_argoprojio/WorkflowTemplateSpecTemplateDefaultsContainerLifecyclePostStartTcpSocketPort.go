package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

