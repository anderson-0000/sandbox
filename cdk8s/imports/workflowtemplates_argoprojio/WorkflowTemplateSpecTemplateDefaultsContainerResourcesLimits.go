package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

