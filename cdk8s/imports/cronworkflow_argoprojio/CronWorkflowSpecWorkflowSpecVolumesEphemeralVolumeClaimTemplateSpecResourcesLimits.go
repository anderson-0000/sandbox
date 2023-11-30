package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits
type jsiiProxy_CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromString(value *string) CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

