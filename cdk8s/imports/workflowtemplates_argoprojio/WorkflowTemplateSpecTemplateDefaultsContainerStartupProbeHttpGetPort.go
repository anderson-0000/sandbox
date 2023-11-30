package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

