package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort
type jsiiProxy_WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort_FromString(value *string) WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

