package workflowtemplates _argoprojio


type WorkflowTemplateSpecVolumesProjectedSources struct {
	ConfigMap *WorkflowTemplateSpecVolumesProjectedSourcesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	DownwardApi *WorkflowTemplateSpecVolumesProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	Secret *WorkflowTemplateSpecVolumesProjectedSourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	ServiceAccountToken *WorkflowTemplateSpecVolumesProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

