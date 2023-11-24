package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesInitContainersResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesInitContainersResourcesRequests
type jsiiProxy_WorkflowSpecTemplatesInitContainersResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesInitContainersResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesInitContainersResourcesRequests_FromNumber(value *float64) WorkflowSpecTemplatesInitContainersResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersResourcesRequests

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesInitContainersResourcesRequests_FromString(value *string) WorkflowSpecTemplatesInitContainersResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersResourcesRequests

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

