package workflows_argoprojio


type WorkflowSpecRetryStrategyBackoff struct {
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	Factor WorkflowSpecRetryStrategyBackoffFactor `field:"optional" json:"factor" yaml:"factor"`
	MaxDuration *string `field:"optional" json:"maxDuration" yaml:"maxDuration"`
}

