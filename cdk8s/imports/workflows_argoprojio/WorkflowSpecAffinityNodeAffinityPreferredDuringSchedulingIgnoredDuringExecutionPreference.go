package workflows_argoprojio


type WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference struct {
	MatchExpressions *[]*WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchFields *[]*WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields `field:"optional" json:"matchFields" yaml:"matchFields"`
}

