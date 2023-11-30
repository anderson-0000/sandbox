package workflowtemplates_argoprojio


type WorkflowTemplateSpecRetryStrategyBackoff struct {
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	Factor WorkflowTemplateSpecRetryStrategyBackoffFactor `field:"optional" json:"factor" yaml:"factor"`
	MaxDuration *string `field:"optional" json:"maxDuration" yaml:"maxDuration"`
}

