package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesVolumesEmptyDirSizeLimit interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesVolumesEmptyDirSizeLimit
type jsiiProxy_WorkflowSpecTemplatesVolumesEmptyDirSizeLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesVolumesEmptyDirSizeLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesVolumesEmptyDirSizeLimit_FromNumber(value *float64) WorkflowSpecTemplatesVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesVolumesEmptyDirSizeLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesVolumesEmptyDirSizeLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesVolumesEmptyDirSizeLimit_FromString(value *string) WorkflowSpecTemplatesVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesVolumesEmptyDirSizeLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesVolumesEmptyDirSizeLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

