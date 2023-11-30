package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort
type jsiiProxy_WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort_FromString(value *string) WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

