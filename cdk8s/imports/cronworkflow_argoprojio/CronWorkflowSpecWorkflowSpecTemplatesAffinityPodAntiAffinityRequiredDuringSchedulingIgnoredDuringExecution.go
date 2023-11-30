package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	TopologyKey *string `field:"required" json:"topologyKey" yaml:"topologyKey"`
	LabelSelector *CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	NamespaceSelector *CronWorkflowSpecWorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
}

