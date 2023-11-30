package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecVolumesEmptyDirSizeLimit interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecVolumesEmptyDirSizeLimit
type jsiiProxy_WorkflowTemplateSpecVolumesEmptyDirSizeLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecVolumesEmptyDirSizeLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecVolumesEmptyDirSizeLimit_FromNumber(value *float64) WorkflowTemplateSpecVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumesEmptyDirSizeLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecVolumesEmptyDirSizeLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecVolumesEmptyDirSizeLimit_FromString(value *string) WorkflowTemplateSpecVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumesEmptyDirSizeLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecVolumesEmptyDirSizeLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

