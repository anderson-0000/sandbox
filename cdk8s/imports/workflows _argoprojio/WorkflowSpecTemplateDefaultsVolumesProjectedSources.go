package workflows _argoprojio


type WorkflowSpecTemplateDefaultsVolumesProjectedSources struct {
	ConfigMap *WorkflowSpecTemplateDefaultsVolumesProjectedSourcesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	DownwardApi *WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	Secret *WorkflowSpecTemplateDefaultsVolumesProjectedSourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	ServiceAccountToken *WorkflowSpecTemplateDefaultsVolumesProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

