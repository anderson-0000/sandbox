package workflows_argoprojio


type WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	PodAffinityTerm *WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

