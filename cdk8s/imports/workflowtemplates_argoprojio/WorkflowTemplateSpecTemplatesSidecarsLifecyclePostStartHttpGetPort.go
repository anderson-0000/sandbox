package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort
type jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort_FromNumber(value *float64) WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort_FromString(value *string) WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

