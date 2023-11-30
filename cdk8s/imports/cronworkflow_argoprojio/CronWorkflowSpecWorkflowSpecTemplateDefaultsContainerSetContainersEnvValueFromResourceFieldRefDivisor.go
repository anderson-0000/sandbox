package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

