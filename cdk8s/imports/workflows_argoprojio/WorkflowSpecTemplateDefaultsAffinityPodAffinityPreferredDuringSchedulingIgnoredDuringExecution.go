package workflows_argoprojio


type WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	PodAffinityTerm *WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

