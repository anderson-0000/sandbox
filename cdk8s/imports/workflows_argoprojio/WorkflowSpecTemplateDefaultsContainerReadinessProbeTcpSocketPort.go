package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort
type jsiiProxy_WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort_FromString(value *string) WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

