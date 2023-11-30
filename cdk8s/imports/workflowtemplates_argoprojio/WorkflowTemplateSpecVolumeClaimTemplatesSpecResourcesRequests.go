package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests
type jsiiProxy_WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests_FromNumber(value *float64) WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests_FromString(value *string) WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

