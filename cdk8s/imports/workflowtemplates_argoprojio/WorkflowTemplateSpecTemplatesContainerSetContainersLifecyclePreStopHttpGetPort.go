package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort
type jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort_FromNumber(value *float64) WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort_FromString(value *string) WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

