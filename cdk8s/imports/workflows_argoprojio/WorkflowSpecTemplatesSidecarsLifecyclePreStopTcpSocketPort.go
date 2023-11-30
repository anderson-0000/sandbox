package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort
type jsiiProxy_WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort_FromString(value *string) WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

