package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecPodDisruptionBudgetMinAvailable interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecPodDisruptionBudgetMinAvailable
type jsiiProxy_WorkflowTemplateSpecPodDisruptionBudgetMinAvailable struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecPodDisruptionBudgetMinAvailable) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecPodDisruptionBudgetMinAvailable_FromNumber(value *float64) WorkflowTemplateSpecPodDisruptionBudgetMinAvailable {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecPodDisruptionBudgetMinAvailable_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecPodDisruptionBudgetMinAvailable

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecPodDisruptionBudgetMinAvailable",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecPodDisruptionBudgetMinAvailable_FromString(value *string) WorkflowTemplateSpecPodDisruptionBudgetMinAvailable {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecPodDisruptionBudgetMinAvailable_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecPodDisruptionBudgetMinAvailable

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecPodDisruptionBudgetMinAvailable",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

