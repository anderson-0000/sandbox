package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecRetryStrategy struct {
	Affinity *CronWorkflowSpecWorkflowSpecRetryStrategyAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	Backoff *CronWorkflowSpecWorkflowSpecRetryStrategyBackoff `field:"optional" json:"backoff" yaml:"backoff"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Limit CronWorkflowSpecWorkflowSpecRetryStrategyLimit `field:"optional" json:"limit" yaml:"limit"`
	RetryPolicy *string `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
}

