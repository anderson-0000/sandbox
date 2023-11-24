package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesVolumesProjectedSources struct {
	ConfigMap *WorkflowTemplateSpecTemplatesVolumesProjectedSourcesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	DownwardApi *WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	Secret *WorkflowTemplateSpecTemplatesVolumesProjectedSourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	ServiceAccountToken *WorkflowTemplateSpecTemplatesVolumesProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

