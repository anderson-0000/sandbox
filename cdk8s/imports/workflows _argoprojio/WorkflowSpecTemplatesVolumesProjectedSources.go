package workflows _argoprojio


type WorkflowSpecTemplatesVolumesProjectedSources struct {
	ConfigMap *WorkflowSpecTemplatesVolumesProjectedSourcesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	DownwardApi *WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	Secret *WorkflowSpecTemplatesVolumesProjectedSourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	ServiceAccountToken *WorkflowSpecTemplatesVolumesProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

