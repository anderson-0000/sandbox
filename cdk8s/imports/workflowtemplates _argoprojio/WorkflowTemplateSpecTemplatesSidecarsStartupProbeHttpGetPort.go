package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort
type jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort_FromNumber(value *float64) WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort_FromString(value *string) WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

