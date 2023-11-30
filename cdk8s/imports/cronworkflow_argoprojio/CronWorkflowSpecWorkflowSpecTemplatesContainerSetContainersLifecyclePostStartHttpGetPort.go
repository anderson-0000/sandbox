package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

