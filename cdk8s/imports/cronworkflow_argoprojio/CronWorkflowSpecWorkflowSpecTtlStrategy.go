package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTtlStrategy struct {
	SecondsAfterCompletion *float64 `field:"optional" json:"secondsAfterCompletion" yaml:"secondsAfterCompletion"`
	SecondsAfterFailure *float64 `field:"optional" json:"secondsAfterFailure" yaml:"secondsAfterFailure"`
	SecondsAfterSuccess *float64 `field:"optional" json:"secondsAfterSuccess" yaml:"secondsAfterSuccess"`
}

