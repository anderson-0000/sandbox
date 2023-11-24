package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsAffinity struct {
	NodeAffinity *WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	PodAffinity *WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	PodAntiAffinity *WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

