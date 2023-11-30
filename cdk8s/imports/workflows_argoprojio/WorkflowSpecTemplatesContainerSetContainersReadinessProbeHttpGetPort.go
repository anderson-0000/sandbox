package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort
type jsiiProxy_WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort_FromString(value *string) WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

