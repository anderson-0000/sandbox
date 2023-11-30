package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecPodDisruptionBudgetMinAvailable interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecPodDisruptionBudgetMinAvailable
type jsiiProxy_WorkflowSpecPodDisruptionBudgetMinAvailable struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecPodDisruptionBudgetMinAvailable) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecPodDisruptionBudgetMinAvailable_FromNumber(value *float64) WorkflowSpecPodDisruptionBudgetMinAvailable {
	_init_.Initialize()

	if err := validateWorkflowSpecPodDisruptionBudgetMinAvailable_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecPodDisruptionBudgetMinAvailable

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecPodDisruptionBudgetMinAvailable",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecPodDisruptionBudgetMinAvailable_FromString(value *string) WorkflowSpecPodDisruptionBudgetMinAvailable {
	_init_.Initialize()

	if err := validateWorkflowSpecPodDisruptionBudgetMinAvailable_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecPodDisruptionBudgetMinAvailable

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecPodDisruptionBudgetMinAvailable",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

