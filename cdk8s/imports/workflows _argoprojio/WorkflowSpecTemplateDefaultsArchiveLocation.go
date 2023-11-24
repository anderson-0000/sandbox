package workflows _argoprojio


type WorkflowSpecTemplateDefaultsArchiveLocation struct {
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	Artifactory *WorkflowSpecTemplateDefaultsArchiveLocationArtifactory `field:"optional" json:"artifactory" yaml:"artifactory"`
	Azure *WorkflowSpecTemplateDefaultsArchiveLocationAzure `field:"optional" json:"azure" yaml:"azure"`
	Gcs *WorkflowSpecTemplateDefaultsArchiveLocationGcs `field:"optional" json:"gcs" yaml:"gcs"`
	Git *WorkflowSpecTemplateDefaultsArchiveLocationGit `field:"optional" json:"git" yaml:"git"`
	Hdfs *WorkflowSpecTemplateDefaultsArchiveLocationHdfs `field:"optional" json:"hdfs" yaml:"hdfs"`
	Http *WorkflowSpecTemplateDefaultsArchiveLocationHttp `field:"optional" json:"http" yaml:"http"`
	Oss *WorkflowSpecTemplateDefaultsArchiveLocationOss `field:"optional" json:"oss" yaml:"oss"`
	Raw *WorkflowSpecTemplateDefaultsArchiveLocationRaw `field:"optional" json:"raw" yaml:"raw"`
	S3 *WorkflowSpecTemplateDefaultsArchiveLocationS3 `field:"optional" json:"s3" yaml:"s3"`
}

