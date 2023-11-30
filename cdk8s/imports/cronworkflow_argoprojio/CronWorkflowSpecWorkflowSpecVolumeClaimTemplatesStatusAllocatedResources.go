package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources
type jsiiProxy_CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources_FromString(value *string) CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusAllocatedResources",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

