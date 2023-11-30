package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsContainerResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsContainerResourcesLimits
type jsiiProxy_WorkflowSpecTemplateDefaultsContainerResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsContainerResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsContainerResourcesLimits_FromNumber(value *float64) WorkflowSpecTemplateDefaultsContainerResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsContainerResourcesLimits_FromString(value *string) WorkflowSpecTemplateDefaultsContainerResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

