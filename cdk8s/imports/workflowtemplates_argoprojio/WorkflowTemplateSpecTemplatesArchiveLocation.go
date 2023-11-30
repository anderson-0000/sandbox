package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesArchiveLocation struct {
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	Artifactory *WorkflowTemplateSpecTemplatesArchiveLocationArtifactory `field:"optional" json:"artifactory" yaml:"artifactory"`
	Azure *WorkflowTemplateSpecTemplatesArchiveLocationAzure `field:"optional" json:"azure" yaml:"azure"`
	Gcs *WorkflowTemplateSpecTemplatesArchiveLocationGcs `field:"optional" json:"gcs" yaml:"gcs"`
	Git *WorkflowTemplateSpecTemplatesArchiveLocationGit `field:"optional" json:"git" yaml:"git"`
	Hdfs *WorkflowTemplateSpecTemplatesArchiveLocationHdfs `field:"optional" json:"hdfs" yaml:"hdfs"`
	Http *WorkflowTemplateSpecTemplatesArchiveLocationHttp `field:"optional" json:"http" yaml:"http"`
	Oss *WorkflowTemplateSpecTemplatesArchiveLocationOss `field:"optional" json:"oss" yaml:"oss"`
	Raw *WorkflowTemplateSpecTemplatesArchiveLocationRaw `field:"optional" json:"raw" yaml:"raw"`
	S3 *WorkflowTemplateSpecTemplatesArchiveLocationS3 `field:"optional" json:"s3" yaml:"s3"`
}

