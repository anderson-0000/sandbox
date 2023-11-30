package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

