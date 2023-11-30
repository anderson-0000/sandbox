package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests
type jsiiProxy_CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromString(value *string) CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

