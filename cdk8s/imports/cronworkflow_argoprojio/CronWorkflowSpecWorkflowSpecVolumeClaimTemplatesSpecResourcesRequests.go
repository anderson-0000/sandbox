package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests
type jsiiProxy_CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests_FromString(value *string) CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

