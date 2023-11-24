package cronworkflow _argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable
type jsiiProxy_CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable_FromString(value *string) CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecPodDisruptionBudgetMinAvailable",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

