package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsArchiveLocation struct {
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	Artifactory *WorkflowTemplateSpecTemplateDefaultsArchiveLocationArtifactory `field:"optional" json:"artifactory" yaml:"artifactory"`
	Azure *WorkflowTemplateSpecTemplateDefaultsArchiveLocationAzure `field:"optional" json:"azure" yaml:"azure"`
	Gcs *WorkflowTemplateSpecTemplateDefaultsArchiveLocationGcs `field:"optional" json:"gcs" yaml:"gcs"`
	Git *WorkflowTemplateSpecTemplateDefaultsArchiveLocationGit `field:"optional" json:"git" yaml:"git"`
	Hdfs *WorkflowTemplateSpecTemplateDefaultsArchiveLocationHdfs `field:"optional" json:"hdfs" yaml:"hdfs"`
	Http *WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttp `field:"optional" json:"http" yaml:"http"`
	Oss *WorkflowTemplateSpecTemplateDefaultsArchiveLocationOss `field:"optional" json:"oss" yaml:"oss"`
	Raw *WorkflowTemplateSpecTemplateDefaultsArchiveLocationRaw `field:"optional" json:"raw" yaml:"raw"`
	S3 *WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3 `field:"optional" json:"s3" yaml:"s3"`
}

