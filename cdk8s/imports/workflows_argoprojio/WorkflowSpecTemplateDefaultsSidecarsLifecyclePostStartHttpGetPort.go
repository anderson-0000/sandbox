package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort
type jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort_FromString(value *string) WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

