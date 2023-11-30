package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort
type jsiiProxy_WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort_FromNumber(value *float64) WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort_FromString(value *string) WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

