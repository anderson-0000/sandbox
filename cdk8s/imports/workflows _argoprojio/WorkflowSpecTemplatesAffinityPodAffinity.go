package workflows _argoprojio


type WorkflowSpecTemplatesAffinityPodAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution *[]*WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"preferredDuringSchedulingIgnoredDuringExecution" yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	RequiredDuringSchedulingIgnoredDuringExecution *[]*WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"requiredDuringSchedulingIgnoredDuringExecution" yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}

