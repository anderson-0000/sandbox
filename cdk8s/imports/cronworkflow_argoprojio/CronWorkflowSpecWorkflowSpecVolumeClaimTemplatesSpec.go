package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpec struct {
	AccessModes *[]*string `field:"optional" json:"accessModes" yaml:"accessModes"`
	DataSource *CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	DataSourceRef *CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecDataSourceRef `field:"optional" json:"dataSourceRef" yaml:"dataSourceRef"`
	Resources *CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecResources `field:"optional" json:"resources" yaml:"resources"`
	Selector *CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	VolumeMode *string `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

