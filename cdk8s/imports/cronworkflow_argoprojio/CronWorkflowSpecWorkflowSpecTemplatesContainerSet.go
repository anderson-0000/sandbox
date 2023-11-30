package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerSet struct {
	Containers *[]*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainers `field:"required" json:"containers" yaml:"containers"`
	RetryStrategy *CronWorkflowSpecWorkflowSpecTemplatesContainerSetRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	VolumeMounts *[]*CronWorkflowSpecWorkflowSpecTemplatesContainerSetVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}

