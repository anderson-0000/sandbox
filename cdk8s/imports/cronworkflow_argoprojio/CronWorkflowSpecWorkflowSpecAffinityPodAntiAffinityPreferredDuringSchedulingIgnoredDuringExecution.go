package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	PodAffinityTerm *CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

