package cronworkflow_argoprojio

import (
	_init_ "example.com/cdk8s/imports/cronworkflow_argoprojio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort
type jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort_FromNumber(value *float64) CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort_FromString(value *string) CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort {
	_init_.Initialize()

	if err := validateCronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort

	_jsii_.StaticInvoke(
		"cronworkflow_argoprojio.CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

