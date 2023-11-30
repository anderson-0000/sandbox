package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesSidecarsResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesSidecarsResourcesLimits
type jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesSidecarsResourcesLimits_FromNumber(value *float64) WorkflowTemplateSpecTemplatesSidecarsResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesSidecarsResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesSidecarsResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesSidecarsResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesSidecarsResourcesLimits_FromString(value *string) WorkflowTemplateSpecTemplatesSidecarsResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesSidecarsResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesSidecarsResourcesLimits

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesSidecarsResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

