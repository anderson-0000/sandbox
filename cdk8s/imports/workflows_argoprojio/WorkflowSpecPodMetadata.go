package workflows_argoprojio


type WorkflowSpecPodMetadata struct {
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

