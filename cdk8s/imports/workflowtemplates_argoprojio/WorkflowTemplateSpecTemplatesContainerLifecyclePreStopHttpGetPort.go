package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort
type jsiiProxy_WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort_FromNumber(value *float64) WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort_FromString(value *string) WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

