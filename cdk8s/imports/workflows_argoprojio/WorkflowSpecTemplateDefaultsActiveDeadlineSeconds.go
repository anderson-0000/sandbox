package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsActiveDeadlineSeconds interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsActiveDeadlineSeconds
type jsiiProxy_WorkflowSpecTemplateDefaultsActiveDeadlineSeconds struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsActiveDeadlineSeconds) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsActiveDeadlineSeconds_FromNumber(value *float64) WorkflowSpecTemplateDefaultsActiveDeadlineSeconds {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsActiveDeadlineSeconds_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsActiveDeadlineSeconds

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsActiveDeadlineSeconds",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsActiveDeadlineSeconds_FromString(value *string) WorkflowSpecTemplateDefaultsActiveDeadlineSeconds {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsActiveDeadlineSeconds_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsActiveDeadlineSeconds

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplateDefaultsActiveDeadlineSeconds",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

