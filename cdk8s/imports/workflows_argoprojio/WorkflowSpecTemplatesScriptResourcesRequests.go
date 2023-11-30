package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesScriptResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesScriptResourcesRequests
type jsiiProxy_WorkflowSpecTemplatesScriptResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesScriptResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesScriptResourcesRequests_FromNumber(value *float64) WorkflowSpecTemplatesScriptResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesScriptResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesScriptResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesScriptResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesScriptResourcesRequests_FromString(value *string) WorkflowSpecTemplatesScriptResourcesRequests {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesScriptResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesScriptResourcesRequests

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesScriptResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

