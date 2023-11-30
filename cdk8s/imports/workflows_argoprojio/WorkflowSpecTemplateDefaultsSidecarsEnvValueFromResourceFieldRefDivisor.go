package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor
type jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor_FromString(value *string) WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

