package workflowtemplates _argoprojio


type WorkflowTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	Preference *WorkflowTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference `field:"required" json:"preference" yaml:"preference"`
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

