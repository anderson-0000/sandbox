package workflows _argoprojio


type WorkflowSpecRetryStrategy struct {
	Affinity *WorkflowSpecRetryStrategyAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	Backoff *WorkflowSpecRetryStrategyBackoff `field:"optional" json:"backoff" yaml:"backoff"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Limit WorkflowSpecRetryStrategyLimit `field:"optional" json:"limit" yaml:"limit"`
	RetryPolicy *string `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
}

