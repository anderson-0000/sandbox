package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesScriptResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesScriptResourcesLimits
type jsiiProxy_WorkflowTemplateSpecTemplatesScriptResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesScriptResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesScriptResourcesLimits_FromNumber(value *float64) WorkflowTemplateSpecTemplatesScriptResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesScriptResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesScriptResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesScriptResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesScriptResourcesLimits_FromString(value *string) WorkflowTemplateSpecTemplatesScriptResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesScriptResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesScriptResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesScriptResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

