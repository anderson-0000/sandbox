package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort
type jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort_FromString(value *string) WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

