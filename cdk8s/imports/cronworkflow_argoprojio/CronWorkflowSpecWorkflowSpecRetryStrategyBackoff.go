package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecRetryStrategyBackoff struct {
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	Factor CronWorkflowSpecWorkflowSpecRetryStrategyBackoffFactor `field:"optional" json:"factor" yaml:"factor"`
	MaxDuration *string `field:"optional" json:"maxDuration" yaml:"maxDuration"`
}

