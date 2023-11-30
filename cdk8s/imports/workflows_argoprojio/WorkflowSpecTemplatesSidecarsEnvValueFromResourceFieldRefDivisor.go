package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor
type jsiiProxy_WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor_FromString(value *string) WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

