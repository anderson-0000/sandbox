package workflows_argoprojio


type WorkflowSpecAffinity struct {
	NodeAffinity *WorkflowSpecAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	PodAffinity *WorkflowSpecAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	PodAntiAffinity *WorkflowSpecAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

