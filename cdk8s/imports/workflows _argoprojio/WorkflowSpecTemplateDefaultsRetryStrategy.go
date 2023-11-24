package workflows _argoprojio


type WorkflowSpecTemplateDefaultsRetryStrategy struct {
	Affinity *WorkflowSpecTemplateDefaultsRetryStrategyAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	Backoff *WorkflowSpecTemplateDefaultsRetryStrategyBackoff `field:"optional" json:"backoff" yaml:"backoff"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Limit WorkflowSpecTemplateDefaultsRetryStrategyLimit `field:"optional" json:"limit" yaml:"limit"`
	RetryPolicy *string `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
}

