package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

