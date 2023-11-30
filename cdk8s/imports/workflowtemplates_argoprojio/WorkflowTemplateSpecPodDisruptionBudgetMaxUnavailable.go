package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable
type jsiiProxy_WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable_FromNumber(value *float64) WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable_FromString(value *string) WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

