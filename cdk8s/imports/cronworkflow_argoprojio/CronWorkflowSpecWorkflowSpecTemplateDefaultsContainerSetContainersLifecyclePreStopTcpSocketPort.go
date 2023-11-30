package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

