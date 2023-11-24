package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests
type jsiiProxy_WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests_FromNumber(value *float64) WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests_FromString(value *string) WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecVolumeClaimTemplatesSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

