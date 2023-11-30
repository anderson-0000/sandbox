package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecVolumesDownwardApiItemsResourceFieldRef struct {
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	Divisor CronWorkflowSpecWorkflowSpecVolumesDownwardApiItemsResourceFieldRefDivisor `field:"optional" json:"divisor" yaml:"divisor"`
}

