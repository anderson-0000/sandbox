package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

