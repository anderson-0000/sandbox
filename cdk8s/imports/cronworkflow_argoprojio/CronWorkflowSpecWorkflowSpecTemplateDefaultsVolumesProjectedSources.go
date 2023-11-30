package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSources struct {
	ConfigMap *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	DownwardApi *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	Secret *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesSecret `field:"optional" json:"secret" yaml:"secret"`
	ServiceAccountToken *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

