package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesContainerResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesContainerResourcesRequests
type jsiiProxy_WorkflowSpecTemplatesContainerResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesContainerResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesContainerResourcesRequests_FromNumber(value *float64) WorkflowSpecTemplatesContainerResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesContainerResourcesRequests_FromString(value *string) WorkflowSpecTemplatesContainerResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

