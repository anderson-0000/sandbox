package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor
type jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor_FromString(value *string) WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

