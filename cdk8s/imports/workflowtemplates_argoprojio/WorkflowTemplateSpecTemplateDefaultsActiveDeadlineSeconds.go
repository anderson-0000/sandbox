package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

