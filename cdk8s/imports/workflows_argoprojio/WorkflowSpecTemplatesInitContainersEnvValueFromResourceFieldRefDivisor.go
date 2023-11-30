package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor
type jsiiProxy_WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor_FromString(value *string) WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

