package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

