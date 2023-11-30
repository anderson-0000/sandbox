package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsInitContainersResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsInitContainersResourcesLimits
type jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsInitContainersResourcesLimits_FromNumber(value *float64) WorkflowSpecTemplateDefaultsInitContainersResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsInitContainersResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsInitContainersResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsInitContainersResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsInitContainersResourcesLimits_FromString(value *string) WorkflowSpecTemplateDefaultsInitContainersResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsInitContainersResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsInitContainersResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsInitContainersResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

