package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

