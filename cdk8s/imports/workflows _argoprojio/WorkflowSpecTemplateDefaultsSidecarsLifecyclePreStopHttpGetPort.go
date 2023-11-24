package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort
type jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort_FromString(value *string) WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

