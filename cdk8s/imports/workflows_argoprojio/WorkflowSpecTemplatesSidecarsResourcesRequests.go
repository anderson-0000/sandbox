package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesSidecarsResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesSidecarsResourcesRequests
type jsiiProxy_WorkflowSpecTemplatesSidecarsResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesSidecarsResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesSidecarsResourcesRequests_FromNumber(value *float64) WorkflowSpecTemplatesSidecarsResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesSidecarsResourcesRequests_FromString(value *string) WorkflowSpecTemplatesSidecarsResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

