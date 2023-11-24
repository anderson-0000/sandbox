package workflowtemplates _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflowtemplates _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests
type jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests_FromNumber(value *float64) WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests_FromString(value *string) WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests

	_jsii_.StaticInvoke(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

