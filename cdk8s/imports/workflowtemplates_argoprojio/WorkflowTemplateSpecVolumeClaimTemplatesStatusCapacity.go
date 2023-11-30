package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity
type jsiiProxy_WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity_FromNumber(value *float64) WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity_FromString(value *string) WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

