package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesAffinity struct {
	NodeAffinity *CronWorkflowSpecWorkflowSpecTemplatesAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	PodAffinity *CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	PodAntiAffinity *CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

