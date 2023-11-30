package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

