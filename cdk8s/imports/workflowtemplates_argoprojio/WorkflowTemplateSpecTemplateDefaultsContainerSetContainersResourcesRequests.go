package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

