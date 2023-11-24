package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSources struct {
	ConfigMap *CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	DownwardApi *CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	Secret *CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	ServiceAccountToken *CronWorkflowSpecWorkflowSpecTemplatesVolumesProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

