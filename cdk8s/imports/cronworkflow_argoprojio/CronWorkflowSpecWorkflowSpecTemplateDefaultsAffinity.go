package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinity struct {
	NodeAffinity *CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	PodAffinity *CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	PodAntiAffinity *CronWorkflowSpecWorkflowSpecTemplateDefaultsAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

