package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	TopologyKey *string `field:"required" json:"topologyKey" yaml:"topologyKey"`
	LabelSelector *WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	NamespaceSelector *WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
}

