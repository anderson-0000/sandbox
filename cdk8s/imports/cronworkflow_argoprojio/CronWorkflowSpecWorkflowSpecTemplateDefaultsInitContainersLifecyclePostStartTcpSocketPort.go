package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

