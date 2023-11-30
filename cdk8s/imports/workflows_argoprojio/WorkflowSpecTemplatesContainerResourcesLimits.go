package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesContainerResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesContainerResourcesLimits
type jsiiProxy_WorkflowSpecTemplatesContainerResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesContainerResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesContainerResourcesLimits_FromNumber(value *float64) WorkflowSpecTemplatesContainerResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesContainerResourcesLimits_FromString(value *string) WorkflowSpecTemplatesContainerResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

