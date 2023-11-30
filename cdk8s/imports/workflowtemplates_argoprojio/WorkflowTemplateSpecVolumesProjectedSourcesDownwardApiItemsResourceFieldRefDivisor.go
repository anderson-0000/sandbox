package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor
type jsiiProxy_WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor_FromNumber(value *float64) WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor_FromString(value *string) WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

