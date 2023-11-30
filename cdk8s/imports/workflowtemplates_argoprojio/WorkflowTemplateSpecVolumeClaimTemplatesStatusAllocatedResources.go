package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources
type jsiiProxy_WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources_FromNumber(value *float64) WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources_FromString(value *string) WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

