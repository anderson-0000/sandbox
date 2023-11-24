package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecAffinity struct {
	NodeAffinity *CronWorkflowSpecWorkflowSpecAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	PodAffinity *CronWorkflowSpecWorkflowSpecAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	PodAntiAffinity *CronWorkflowSpecWorkflowSpecAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

