package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesRetryStrategy struct {
	Affinity *WorkflowTemplateSpecTemplatesRetryStrategyAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	Backoff *WorkflowTemplateSpecTemplatesRetryStrategyBackoff `field:"optional" json:"backoff" yaml:"backoff"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	Limit WorkflowTemplateSpecTemplatesRetryStrategyLimit `field:"optional" json:"limit" yaml:"limit"`
	RetryPolicy *string `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
}

