package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

