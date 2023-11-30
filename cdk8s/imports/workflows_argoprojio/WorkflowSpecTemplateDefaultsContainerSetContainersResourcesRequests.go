package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests
type jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests_FromNumber(value *float64) WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests_FromString(value *string) WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

