package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution *[]*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"preferredDuringSchedulingIgnoredDuringExecution" yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	RequiredDuringSchedulingIgnoredDuringExecution *[]*CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"requiredDuringSchedulingIgnoredDuringExecution" yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}
