package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesArchiveLocation struct {
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	Artifactory *CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationArtifactory `field:"optional" json:"artifactory" yaml:"artifactory"`
	Azure *CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationAzure `field:"optional" json:"azure" yaml:"azure"`
	Gcs *CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGcs `field:"optional" json:"gcs" yaml:"gcs"`
	Git *CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationGit `field:"optional" json:"git" yaml:"git"`
	Hdfs *CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHdfs `field:"optional" json:"hdfs" yaml:"hdfs"`
	Http *CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationHttp `field:"optional" json:"http" yaml:"http"`
	Oss *CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationOss `field:"optional" json:"oss" yaml:"oss"`
	Raw *CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationRaw `field:"optional" json:"raw" yaml:"raw"`
	S3 *CronWorkflowSpecWorkflowSpecTemplatesArchiveLocationS3 `field:"optional" json:"s3" yaml:"s3"`
}

