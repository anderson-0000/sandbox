package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

