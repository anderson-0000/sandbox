package workflowtemplates_argoprojio


type WorkflowTemplateSpecAffinityNodeAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution *[]*WorkflowTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"preferredDuringSchedulingIgnoredDuringExecution" yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	RequiredDuringSchedulingIgnoredDuringExecution *WorkflowTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"requiredDuringSchedulingIgnoredDuringExecution" yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}

