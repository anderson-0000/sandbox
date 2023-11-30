package workflowtemplates_argoprojio


type WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	TopologyKey *string `field:"required" json:"topologyKey" yaml:"topologyKey"`
	LabelSelector *WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	NamespaceSelector *WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
}

