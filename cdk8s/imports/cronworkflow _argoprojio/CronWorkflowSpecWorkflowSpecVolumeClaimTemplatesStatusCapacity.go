package cronworkflow _argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity
type jsiiProxy_CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity_FromString(value *string) CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatusCapacity",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

