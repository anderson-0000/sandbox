package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDirSizeLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

