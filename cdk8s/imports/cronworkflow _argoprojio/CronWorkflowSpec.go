package cronworkflow _argoprojio


type CronWorkflowSpec struct {
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	WorkflowSpec *CronWorkflowSpecWorkflowSpec `field:"required" json:"workflowSpec" yaml:"workflowSpec"`
	ConcurrencyPolicy *string `field:"optional" json:"concurrencyPolicy" yaml:"concurrencyPolicy"`
	FailedJobsHistoryLimit *float64 `field:"optional" json:"failedJobsHistoryLimit" yaml:"failedJobsHistoryLimit"`
	StartingDeadlineSeconds *float64 `field:"optional" json:"startingDeadlineSeconds" yaml:"startingDeadlineSeconds"`
	SuccessfulJobsHistoryLimit *float64 `field:"optional" json:"successfulJobsHistoryLimit" yaml:"successfulJobsHistoryLimit"`
	Suspend *bool `field:"optional" json:"suspend" yaml:"suspend"`
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	WorkflowMetadata interface{} `field:"optional" json:"workflowMetadata" yaml:"workflowMetadata"`
}

