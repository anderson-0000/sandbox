package workflowtemplates_argoprojio


type WorkflowTemplateSpecPodMetadata struct {
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

