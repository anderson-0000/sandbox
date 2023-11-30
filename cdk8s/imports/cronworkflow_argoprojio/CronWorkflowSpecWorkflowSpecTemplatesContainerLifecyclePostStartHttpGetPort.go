package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

