package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

