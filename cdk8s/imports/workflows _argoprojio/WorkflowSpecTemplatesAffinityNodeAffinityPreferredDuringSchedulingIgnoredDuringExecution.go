package workflows _argoprojio


type WorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	Preference *WorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference `field:"required" json:"preference" yaml:"preference"`
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

