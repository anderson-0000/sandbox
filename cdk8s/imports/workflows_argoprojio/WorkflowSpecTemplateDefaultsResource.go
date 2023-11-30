package workflows_argoprojio


type WorkflowSpecTemplateDefaultsResource struct {
	Action *string `field:"required" json:"action" yaml:"action"`
	FailureCondition *string `field:"optional" json:"failureCondition" yaml:"failureCondition"`
	Flags *[]*string `field:"optional" json:"flags" yaml:"flags"`
	Manifest *string `field:"optional" json:"manifest" yaml:"manifest"`
	ManifestFrom *WorkflowSpecTemplateDefaultsResourceManifestFrom `field:"optional" json:"manifestFrom" yaml:"manifestFrom"`
	MergeStrategy *string `field:"optional" json:"mergeStrategy" yaml:"mergeStrategy"`
	SetOwnerReference *bool `field:"optional" json:"setOwnerReference" yaml:"setOwnerReference"`
	SuccessCondition *string `field:"optional" json:"successCondition" yaml:"successCondition"`
}

