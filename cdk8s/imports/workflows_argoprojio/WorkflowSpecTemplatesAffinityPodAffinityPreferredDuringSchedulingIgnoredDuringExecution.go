package workflows_argoprojio


type WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	PodAffinityTerm *WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

