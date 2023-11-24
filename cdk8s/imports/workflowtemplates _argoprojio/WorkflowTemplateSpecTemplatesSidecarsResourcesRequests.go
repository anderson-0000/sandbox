package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesSidecarsResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesSidecarsResourcesRequests
type jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesSidecarsResourcesRequests_FromNumber(value *float64) WorkflowTemplateSpecTemplatesSidecarsResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesSidecarsResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesSidecarsResourcesRequests

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesSidecarsResourcesRequests_FromString(value *string) WorkflowTemplateSpecTemplatesSidecarsResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesSidecarsResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesSidecarsResourcesRequests

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

