package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesAffinity struct {
	NodeAffinity *WorkflowTemplateSpecTemplatesAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	PodAffinity *WorkflowTemplateSpecTemplatesAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	PodAntiAffinity *WorkflowTemplateSpecTemplatesAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

