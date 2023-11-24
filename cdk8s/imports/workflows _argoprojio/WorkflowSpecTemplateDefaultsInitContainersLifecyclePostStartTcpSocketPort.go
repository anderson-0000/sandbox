package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort
type jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort_FromString(value *string) WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

