package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

