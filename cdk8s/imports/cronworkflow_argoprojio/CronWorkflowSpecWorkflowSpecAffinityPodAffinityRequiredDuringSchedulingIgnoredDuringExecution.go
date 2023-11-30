package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	TopologyKey *string `field:"required" json:"topologyKey" yaml:"topologyKey"`
	LabelSelector *CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	NamespaceSelector *CronWorkflowSpecWorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
}

