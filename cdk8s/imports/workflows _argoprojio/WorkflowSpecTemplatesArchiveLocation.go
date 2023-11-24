package workflows _argoprojio


type WorkflowSpecTemplatesArchiveLocation struct {
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	Artifactory *WorkflowSpecTemplatesArchiveLocationArtifactory `field:"optional" json:"artifactory" yaml:"artifactory"`
	Azure *WorkflowSpecTemplatesArchiveLocationAzure `field:"optional" json:"azure" yaml:"azure"`
	Gcs *WorkflowSpecTemplatesArchiveLocationGcs `field:"optional" json:"gcs" yaml:"gcs"`
	Git *WorkflowSpecTemplatesArchiveLocationGit `field:"optional" json:"git" yaml:"git"`
	Hdfs *WorkflowSpecTemplatesArchiveLocationHdfs `field:"optional" json:"hdfs" yaml:"hdfs"`
	Http *WorkflowSpecTemplatesArchiveLocationHttp `field:"optional" json:"http" yaml:"http"`
	Oss *WorkflowSpecTemplatesArchiveLocationOss `field:"optional" json:"oss" yaml:"oss"`
	Raw *WorkflowSpecTemplatesArchiveLocationRaw `field:"optional" json:"raw" yaml:"raw"`
	S3 *WorkflowSpecTemplatesArchiveLocationS3 `field:"optional" json:"s3" yaml:"s3"`
}

