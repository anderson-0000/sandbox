package workflows_argoprojio


type WorkflowSpecVolumesProjectedSources struct {
	ConfigMap *WorkflowSpecVolumesProjectedSourcesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	DownwardApi *WorkflowSpecVolumesProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	Secret *WorkflowSpecVolumesProjectedSourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	ServiceAccountToken *WorkflowSpecVolumesProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

