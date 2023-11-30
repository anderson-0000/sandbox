package workflows_argoprojio


type WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

