package workflows_argoprojio


type WorkflowSpecTemplateDefaultsAffinityNodeAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution *[]*WorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"preferredDuringSchedulingIgnoredDuringExecution" yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	RequiredDuringSchedulingIgnoredDuringExecution *WorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"requiredDuringSchedulingIgnoredDuringExecution" yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}

