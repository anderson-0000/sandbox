package workflows_argoprojio


type WorkflowSpecHooksTemplateRef struct {
	ClusterScope *bool `field:"optional" json:"clusterScope" yaml:"clusterScope"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Template *string `field:"optional" json:"template" yaml:"template"`
}

