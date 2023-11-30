package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort
type jsiiProxy_WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort_FromString(value *string) WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

