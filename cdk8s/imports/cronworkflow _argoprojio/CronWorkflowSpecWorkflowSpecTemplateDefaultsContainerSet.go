package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSet struct {
	Containers *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainers `field:"required" json:"containers" yaml:"containers"`
	RetryStrategy *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	VolumeMounts *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}

