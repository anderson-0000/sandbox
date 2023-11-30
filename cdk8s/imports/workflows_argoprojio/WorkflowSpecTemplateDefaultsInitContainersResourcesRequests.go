package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsInitContainersResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsInitContainersResourcesRequests
type jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsInitContainersResourcesRequests_FromNumber(value *float64) WorkflowSpecTemplateDefaultsInitContainersResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsInitContainersResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsInitContainersResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsInitContainersResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsInitContainersResourcesRequests_FromString(value *string) WorkflowSpecTemplateDefaultsInitContainersResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsInitContainersResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsInitContainersResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsInitContainersResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

