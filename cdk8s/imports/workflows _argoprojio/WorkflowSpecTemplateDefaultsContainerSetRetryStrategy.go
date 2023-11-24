package workflows _argoprojio


type WorkflowSpecTemplateDefaultsContainerSetRetryStrategy struct {
	Retries WorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries `field:"required" json:"retries" yaml:"retries"`
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
}

