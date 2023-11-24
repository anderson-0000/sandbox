package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplateDefaultsScriptResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplateDefaultsScriptResourcesRequests
type jsiiProxy_WorkflowSpecTemplateDefaultsScriptResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplateDefaultsScriptResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplateDefaultsScriptResourcesRequests_FromNumber(value *float64) WorkflowSpecTemplateDefaultsScriptResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsScriptResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsScriptResourcesRequests

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplateDefaultsScriptResourcesRequests_FromString(value *string) WorkflowSpecTemplateDefaultsScriptResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplateDefaultsScriptResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplateDefaultsScriptResourcesRequests

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

