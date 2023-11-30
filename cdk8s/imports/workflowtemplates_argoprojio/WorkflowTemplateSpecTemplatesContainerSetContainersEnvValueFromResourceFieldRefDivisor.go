package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor
type jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor_FromString(value *string) WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

