package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

