package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort
type jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort_FromString(value *string) WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

