package workflows_argoprojio

import (
	_init_ "example.com/cdk8s/imports/workflows_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type WorkflowSpecTemplatesSidecarsResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkflowSpecTemplatesSidecarsResourcesLimits
type jsiiProxy_WorkflowSpecTemplatesSidecarsResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowSpecTemplatesSidecarsResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkflowSpecTemplatesSidecarsResourcesLimits_FromNumber(value *float64) WorkflowSpecTemplatesSidecarsResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkflowSpecTemplatesSidecarsResourcesLimits_FromString(value *string) WorkflowSpecTemplatesSidecarsResourcesLimits {
	_init_.Initialize()

	if err := validateWorkflowSpecTemplatesSidecarsResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowSpecTemplatesSidecarsResourcesLimits

	_jsii_.StaticInvoke(
		"workflows_argoprojio.WorkflowSpecTemplatesSidecarsResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

