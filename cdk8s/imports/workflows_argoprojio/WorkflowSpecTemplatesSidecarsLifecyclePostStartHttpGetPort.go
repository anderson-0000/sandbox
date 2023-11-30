package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort
type jsiiProxy_WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort_FromNumber(value *float64) WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort_FromString(value *string) WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

