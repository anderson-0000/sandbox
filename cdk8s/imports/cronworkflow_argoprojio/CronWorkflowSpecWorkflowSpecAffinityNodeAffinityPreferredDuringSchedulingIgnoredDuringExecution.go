package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	Preference *CronWorkflowSpecWorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference `field:"required" json:"preference" yaml:"preference"`
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

