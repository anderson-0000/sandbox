package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

