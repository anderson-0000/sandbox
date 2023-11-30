package workflows_argoprojio


type WorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	Preference *WorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference `field:"required" json:"preference" yaml:"preference"`
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

