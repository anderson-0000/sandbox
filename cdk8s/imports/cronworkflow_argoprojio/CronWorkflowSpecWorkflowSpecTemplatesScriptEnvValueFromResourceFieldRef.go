package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

