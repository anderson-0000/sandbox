package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesContainerSetContainersResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesContainerSetContainersResourcesLimits
type jsiiProxy_WorkflowSpecTemplatesContainerSetContainersResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesContainerSetContainersResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesContainerSetContainersResourcesLimits_FromNumber(value *float64) WorkflowSpecTemplatesContainerSetContainersResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetContainersResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetContainersResourcesLimits

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesContainerSetContainersResourcesLimits_FromString(value *string) WorkflowSpecTemplatesContainerSetContainersResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetContainersResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetContainersResourcesLimits

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

