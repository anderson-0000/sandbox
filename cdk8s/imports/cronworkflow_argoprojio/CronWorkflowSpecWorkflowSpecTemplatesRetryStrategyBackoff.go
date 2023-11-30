package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoff struct {
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	Factor CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoffFactor `field:"optional" json:"factor" yaml:"factor"`
	MaxDuration *string `field:"optional" json:"maxDuration" yaml:"maxDuration"`
}

