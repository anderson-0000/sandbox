package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecVolumesProjectedSources struct {
	ConfigMap *CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	DownwardApi *CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	Secret *CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	ServiceAccountToken *CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

