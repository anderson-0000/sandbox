package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort
type jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort_FromString(value *string) WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

