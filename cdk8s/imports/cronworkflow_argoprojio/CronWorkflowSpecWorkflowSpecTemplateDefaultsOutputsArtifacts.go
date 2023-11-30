package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifacts struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Archive *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArchive `field:"optional" json:"archive" yaml:"archive"`
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	ArtifactGc *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGc `field:"optional" json:"artifactGc" yaml:"artifactGc"`
	Artifactory *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsArtifactory `field:"optional" json:"artifactory" yaml:"artifactory"`
	Azure *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsAzure `field:"optional" json:"azure" yaml:"azure"`
	Deleted *bool `field:"optional" json:"deleted" yaml:"deleted"`
	From *string `field:"optional" json:"from" yaml:"from"`
	FromExpression *string `field:"optional" json:"fromExpression" yaml:"fromExpression"`
	Gcs *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	Git *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsGit `field:"optional" json:"git" yaml:"git"`
	GlobalName *string `field:"optional" json:"globalName" yaml:"globalName"`
	Hdfs *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHdfs `field:"optional" json:"hdfs" yaml:"hdfs"`
	Http *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsHttp `field:"optional" json:"http" yaml:"http"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
	Oss *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsOss `field:"optional" json:"oss" yaml:"oss"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Raw *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsRaw `field:"optional" json:"raw" yaml:"raw"`
	RecurseMode *bool `field:"optional" json:"recurseMode" yaml:"recurseMode"`
	S3 *CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifactsS3 `field:"optional" json:"s3" yaml:"s3"`
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
}

