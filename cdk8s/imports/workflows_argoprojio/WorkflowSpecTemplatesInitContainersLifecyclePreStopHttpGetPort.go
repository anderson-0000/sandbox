package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort
type jsiiProxy_WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort_FromString(value *string) WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

