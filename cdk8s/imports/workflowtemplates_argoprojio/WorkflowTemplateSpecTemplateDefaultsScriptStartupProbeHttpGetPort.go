package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

