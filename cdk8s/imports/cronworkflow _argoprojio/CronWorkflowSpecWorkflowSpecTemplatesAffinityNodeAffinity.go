package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution *[]*CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"preferredDuringSchedulingIgnoredDuringExecution" yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	RequiredDuringSchedulingIgnoredDuringExecution *CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"requiredDuringSchedulingIgnoredDuringExecution" yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}

