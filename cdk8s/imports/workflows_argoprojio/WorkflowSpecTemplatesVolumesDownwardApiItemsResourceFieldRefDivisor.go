package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor
type jsiiProxy_WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor_FromNumber(value *float64) WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor_FromString(value *string) WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

