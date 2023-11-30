package workflowtemplates_argoprojio


type WorkflowTemplateSpecTolerations struct {
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	Key *string `field:"optional" json:"key" yaml:"key"`
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	TolerationSeconds *float64 `field:"optional" json:"tolerationSeconds" yaml:"tolerationSeconds"`
	Value *string `field:"optional" json:"value" yaml:"value"`
}

