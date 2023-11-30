package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesContainerSetContainersResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesContainerSetContainersResourcesRequests
type jsiiProxy_WorkflowSpecTemplatesContainerSetContainersResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesContainerSetContainersResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesContainerSetContainersResourcesRequests_FromNumber(value *float64) WorkflowSpecTemplatesContainerSetContainersResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetContainersResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetContainersResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerSetContainersResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesContainerSetContainersResourcesRequests_FromString(value *string) WorkflowSpecTemplatesContainerSetContainersResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetContainersResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetContainersResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerSetContainersResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

