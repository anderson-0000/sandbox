package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort
type jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort_FromString(value *string) WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

