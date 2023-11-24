package cronworkflow _argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable
type jsiiProxy_CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable_FromString(value *string) CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMaxUnavailable",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

