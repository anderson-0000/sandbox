package workflows_argoprojio


type WorkflowSpecTemplateDefaultsDataSourceArtifactPaths struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Archive *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArchive `field:"optional" json:"archive" yaml:"archive"`
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	ArtifactGc *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGc `field:"optional" json:"artifactGc" yaml:"artifactGc"`
	Artifactory *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactory `field:"optional" json:"artifactory" yaml:"artifactory"`
	Azure *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsAzure `field:"optional" json:"azure" yaml:"azure"`
	Deleted *bool `field:"optional" json:"deleted" yaml:"deleted"`
	From *string `field:"optional" json:"from" yaml:"from"`
	FromExpression *string `field:"optional" json:"fromExpression" yaml:"fromExpression"`
	Gcs *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	Git *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGit `field:"optional" json:"git" yaml:"git"`
	GlobalName *string `field:"optional" json:"globalName" yaml:"globalName"`
	Hdfs *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfs `field:"optional" json:"hdfs" yaml:"hdfs"`
	Http *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttp `field:"optional" json:"http" yaml:"http"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
	Oss *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsOss `field:"optional" json:"oss" yaml:"oss"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Raw *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsRaw `field:"optional" json:"raw" yaml:"raw"`
	RecurseMode *bool `field:"optional" json:"recurseMode" yaml:"recurseMode"`
	S3 *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3 `field:"optional" json:"s3" yaml:"s3"`
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
}

