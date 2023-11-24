package workflows _argoprojio


type WorkflowSpecTemplateDefaultsAffinity struct {
	NodeAffinity *WorkflowSpecTemplateDefaultsAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	PodAffinity *WorkflowSpecTemplateDefaultsAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	PodAntiAffinity *WorkflowSpecTemplateDefaultsAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

