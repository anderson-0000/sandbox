package workflows _argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesInitContainersResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesInitContainersResourcesLimits
type jsiiProxy_WorkflowSpecTemplatesInitContainersResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesInitContainersResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesInitContainersResourcesLimits_FromNumber(value *float64) WorkflowSpecTemplatesInitContainersResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersResourcesLimits

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesInitContainersResourcesLimits_FromString(value *string) WorkflowSpecTemplatesInitContainersResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesInitContainersResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesInitContainersResourcesLimits

	_jsii_.StaticInvoke(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

