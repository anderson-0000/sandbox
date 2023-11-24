package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecPodDisruptionBudgetMaxUnavailable interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecPodDisruptionBudgetMaxUnavailable
type jsiiProxy_WorkflowSpecPodDisruptionBudgetMaxUnavailable struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecPodDisruptionBudgetMaxUnavailable) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecPodDisruptionBudgetMaxUnavailable_FromNumber(value *float64) WorkflowSpecPodDisruptionBudgetMaxUnavailable {
	_init_.Initialize()

	if err := validateWorkflowSpecPodDisruptionBudgetMaxUnavailable_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecPodDisruptionBudgetMaxUnavailable

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecPodDisruptionBudgetMaxUnavailable",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecPodDisruptionBudgetMaxUnavailable_FromString(value *string) WorkflowSpecPodDisruptionBudgetMaxUnavailable {
	_init_.Initialize()

	if err := validateWorkflowSpecPodDisruptionBudgetMaxUnavailable_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecPodDisruptionBudgetMaxUnavailable

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecPodDisruptionBudgetMaxUnavailable",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

