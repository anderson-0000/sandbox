package workflowtemplates_argoprojio


type WorkflowTemplateSpecAffinityPodAntiAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution *[]*WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"preferredDuringSchedulingIgnoredDuringExecution" yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	RequiredDuringSchedulingIgnoredDuringExecution *[]*WorkflowTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"requiredDuringSchedulingIgnoredDuringExecution" yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}
