package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

