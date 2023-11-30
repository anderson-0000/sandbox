package workflows_argoprojio


type WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	TopologyKey *string `field:"required" json:"topologyKey" yaml:"topologyKey"`
	LabelSelector *WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	NamespaceSelector *WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
}

