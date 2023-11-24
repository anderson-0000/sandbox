package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit
type jsiiProxy_WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit_FromNumber(value *float64) WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit_FromString(value *string) WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

