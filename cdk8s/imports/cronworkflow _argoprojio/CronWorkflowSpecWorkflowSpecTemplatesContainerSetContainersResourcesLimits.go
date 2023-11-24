package cronworkflow _argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow _argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow _argoprojio.CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

