package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits
type jsiiProxy_CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits_FromString(value *string) CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

