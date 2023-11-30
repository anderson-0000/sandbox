package workflows_argoprojio


type WorkflowSpecTemplatesAffinity struct {
	NodeAffinity *WorkflowSpecTemplatesAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	PodAffinity *WorkflowSpecTemplatesAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	PodAntiAffinity *WorkflowSpecTemplatesAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

