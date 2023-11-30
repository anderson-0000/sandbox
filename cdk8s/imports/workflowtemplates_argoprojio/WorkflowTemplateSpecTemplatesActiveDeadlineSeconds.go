package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesActiveDeadlineSeconds interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesActiveDeadlineSeconds
type jsiiProxy_WorkflowTemplateSpecTemplatesActiveDeadlineSeconds struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesActiveDeadlineSeconds) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesActiveDeadlineSeconds_FromNumber(value *float64) WorkflowTemplateSpecTemplatesActiveDeadlineSeconds {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesActiveDeadlineSeconds_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesActiveDeadlineSeconds

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesActiveDeadlineSeconds",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesActiveDeadlineSeconds_FromString(value *string) WorkflowTemplateSpecTemplatesActiveDeadlineSeconds {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesActiveDeadlineSeconds_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesActiveDeadlineSeconds

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesActiveDeadlineSeconds",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

