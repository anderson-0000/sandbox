package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits
type jsiiProxy_WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits_FromNumber(value *float64) WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits_FromString(value *string) WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

