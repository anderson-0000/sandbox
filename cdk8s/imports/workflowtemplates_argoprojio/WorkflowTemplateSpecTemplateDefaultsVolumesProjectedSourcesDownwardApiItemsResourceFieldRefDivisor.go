package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

