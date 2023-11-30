package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

