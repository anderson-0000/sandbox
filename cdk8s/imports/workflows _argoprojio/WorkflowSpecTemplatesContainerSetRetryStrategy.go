package workflows _argoprojio


type WorkflowSpecTemplatesContainerSetRetryStrategy struct {
	Retries WorkflowSpecTemplatesContainerSetRetryStrategyRetries `field:"required" json:"retries" yaml:"retries"`
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
}

