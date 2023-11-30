package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort
type jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort_FromNumber(value *float64) WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort_FromString(value *string) WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

