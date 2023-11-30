package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesContainerResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesContainerResourcesRequests
type jsiiProxy_WorkflowTemplateSpecTemplatesContainerResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesContainerResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesContainerResourcesRequests_FromNumber(value *float64) WorkflowTemplateSpecTemplatesContainerResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerResourcesRequests

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesContainerResourcesRequests_FromString(value *string) WorkflowTemplateSpecTemplatesContainerResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerResourcesRequests

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

