package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

