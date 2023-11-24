package workflows _argoprojio


type WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	Preference *WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference `field:"required" json:"preference" yaml:"preference"`
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

