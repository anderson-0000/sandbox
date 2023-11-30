package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

