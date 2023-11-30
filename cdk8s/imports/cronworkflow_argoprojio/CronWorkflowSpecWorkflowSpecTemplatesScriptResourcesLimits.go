package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesScriptResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

