package workflowtemplates_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates_argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

