package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor
type jsiiProxy_WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor_FromNumber(value *float64) WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor_FromString(value *string) WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

