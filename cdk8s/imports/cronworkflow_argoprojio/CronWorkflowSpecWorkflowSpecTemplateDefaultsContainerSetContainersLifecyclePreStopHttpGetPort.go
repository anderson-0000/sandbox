package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

