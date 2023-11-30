package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

