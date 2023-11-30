package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategy struct {
	Affinity *CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	Backoff *CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyBackoff `field:"optional" json:"backoff" yaml:"backoff"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Limit CronWorkflowSpecWorkflowSpecTemplateDefaultsRetryStrategyLimit `field:"optional" json:"limit" yaml:"limit"`
	RetryPolicy *string `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
}

