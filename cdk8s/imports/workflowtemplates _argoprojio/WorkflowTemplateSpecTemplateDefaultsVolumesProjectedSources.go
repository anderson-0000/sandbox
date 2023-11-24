package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSources struct {
	ConfigMap *WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	DownwardApi *WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	Secret *WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	ServiceAccountToken *WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

