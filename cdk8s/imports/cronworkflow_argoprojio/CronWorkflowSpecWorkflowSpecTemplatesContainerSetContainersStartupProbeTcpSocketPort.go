package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

