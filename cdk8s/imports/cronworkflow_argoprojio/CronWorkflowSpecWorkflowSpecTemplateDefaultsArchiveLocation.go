package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocation struct {
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	Artifactory *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationArtifactory `field:"optional" json:"artifactory" yaml:"artifactory"`
	Azure *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationAzure `field:"optional" json:"azure" yaml:"azure"`
	Gcs *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGcs `field:"optional" json:"gcs" yaml:"gcs"`
	Git *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationGit `field:"optional" json:"git" yaml:"git"`
	Hdfs *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHdfs `field:"optional" json:"hdfs" yaml:"hdfs"`
	Http *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationHttp `field:"optional" json:"http" yaml:"http"`
	Oss *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationOss `field:"optional" json:"oss" yaml:"oss"`
	Raw *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationRaw `field:"optional" json:"raw" yaml:"raw"`
	S3 *CronWorkflowSpecWorkflowSpecTemplateDefaultsArchiveLocationS3 `field:"optional" json:"s3" yaml:"s3"`
}

