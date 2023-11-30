package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

