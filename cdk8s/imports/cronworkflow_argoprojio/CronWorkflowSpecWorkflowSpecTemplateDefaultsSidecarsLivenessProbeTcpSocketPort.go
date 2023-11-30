package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

