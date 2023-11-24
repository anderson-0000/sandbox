package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

