package workflowtemplates_argoprojio


type WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	PodAffinityTerm *WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

