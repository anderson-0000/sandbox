package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsContainerResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsContainerResourcesRequests
type jsiiProxy_WorkflowSpecTemplateDefaultsContainerResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsContainerResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsContainerResourcesRequests_FromNumber(value *float64) WorkflowSpecTemplateDefaultsContainerResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsContainerResourcesRequests_FromString(value *string) WorkflowSpecTemplateDefaultsContainerResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

