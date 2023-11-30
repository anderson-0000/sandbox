package workflowtemplates_argoprojio


type WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	PodAffinityTerm *WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

