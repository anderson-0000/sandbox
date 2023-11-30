package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecVolumesEmptyDirSizeLimit interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecVolumesEmptyDirSizeLimit
type jsiiProxy_WorkflowSpecVolumesEmptyDirSizeLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecVolumesEmptyDirSizeLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecVolumesEmptyDirSizeLimit_FromNumber(value *float64) WorkflowSpecVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumesEmptyDirSizeLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecVolumesEmptyDirSizeLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecVolumesEmptyDirSizeLimit_FromString(value *string) WorkflowSpecVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumesEmptyDirSizeLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecVolumesEmptyDirSizeLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

