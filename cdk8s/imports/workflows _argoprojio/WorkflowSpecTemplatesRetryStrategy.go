package workflows _argoprojio


type WorkflowSpecTemplatesRetryStrategy struct {
	Affinity *WorkflowSpecTemplatesRetryStrategyAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	Backoff *WorkflowSpecTemplatesRetryStrategyBackoff `field:"optional" json:"backoff" yaml:"backoff"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Limit WorkflowSpecTemplatesRetryStrategyLimit `field:"optional" json:"limit" yaml:"limit"`
	RetryPolicy *string `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
}

