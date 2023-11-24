package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort
type jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort_FromString(value *string) WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

