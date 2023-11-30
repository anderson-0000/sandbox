package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoff struct {
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	Factor CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor `field:"optional" json:"factor" yaml:"factor"`
	MaxDuration *string `field:"optional" json:"maxDuration" yaml:"maxDuration"`
}

