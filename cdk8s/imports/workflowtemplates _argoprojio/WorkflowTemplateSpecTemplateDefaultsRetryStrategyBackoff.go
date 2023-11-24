package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoff struct {
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	Factor WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor `field:"optional" json:"factor" yaml:"factor"`
	MaxDuration *string `field:"optional" json:"maxDuration" yaml:"maxDuration"`
}

