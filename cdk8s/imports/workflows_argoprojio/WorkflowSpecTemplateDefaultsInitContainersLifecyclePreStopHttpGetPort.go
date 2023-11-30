package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort
type jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort_FromString(value *string) WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

