package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecAffinityPodAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution *[]*CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"preferredDuringSchedulingIgnoredDuringExecution" yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	RequiredDuringSchedulingIgnoredDuringExecution *[]*CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"requiredDuringSchedulingIgnoredDuringExecution" yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}
