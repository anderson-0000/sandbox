package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsRetryStrategy struct {
	Affinity *WorkflowTemplateSpecTemplateDefaultsRetryStrategyAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	Backoff *WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoff `field:"optional" json:"backoff" yaml:"backoff"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Limit WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit `field:"optional" json:"limit" yaml:"limit"`
	RetryPolicy *string `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
}

