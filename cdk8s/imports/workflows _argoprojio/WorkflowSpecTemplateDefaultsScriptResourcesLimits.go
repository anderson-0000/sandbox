package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsScriptResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsScriptResourcesLimits
type jsiiProxy_WorkflowSpecTemplateDefaultsScriptResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsScriptResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsScriptResourcesLimits_FromNumber(value *float64) WorkflowSpecTemplateDefaultsScriptResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsScriptResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsScriptResourcesLimits

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsScriptResourcesLimits_FromString(value *string) WorkflowSpecTemplateDefaultsScriptResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsScriptResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsScriptResourcesLimits

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

