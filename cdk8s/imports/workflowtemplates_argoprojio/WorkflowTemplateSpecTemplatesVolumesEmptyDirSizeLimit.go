package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit
type jsiiProxy_WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit_FromNumber(value *float64) WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit_FromString(value *string) WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

