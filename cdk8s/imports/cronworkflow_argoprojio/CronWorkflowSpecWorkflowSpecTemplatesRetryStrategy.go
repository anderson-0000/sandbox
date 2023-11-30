package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesRetryStrategy struct {
	Affinity *CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	Backoff *CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyBackoff `field:"optional" json:"backoff" yaml:"backoff"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Limit CronWorkflowSpecWorkflowSpecTemplatesRetryStrategyLimit `field:"optional" json:"limit" yaml:"limit"`
	RetryPolicy *string `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
}

