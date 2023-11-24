package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsVolumes struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	AwsElasticBlockStore *WorkflowTemplateSpecTemplateDefaultsVolumesAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	AzureDisk *WorkflowTemplateSpecTemplateDefaultsVolumesAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	AzureFile *WorkflowTemplateSpecTemplateDefaultsVolumesAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	Cephfs *WorkflowTemplateSpecTemplateDefaultsVolumesCephfs `field:"optional" json:"cephfs" yaml:"cephfs"`
	Cinder *WorkflowTemplateSpecTemplateDefaultsVolumesCinder `field:"optional" json:"cinder" yaml:"cinder"`
	ConfigMap *WorkflowTemplateSpecTemplateDefaultsVolumesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	Csi *WorkflowTemplateSpecTemplateDefaultsVolumesCsi `field:"optional" json:"csi" yaml:"csi"`
	DownwardApi *WorkflowTemplateSpecTemplateDefaultsVolumesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	EmptyDir *WorkflowTemplateSpecTemplateDefaultsVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	Ephemeral *WorkflowTemplateSpecTemplateDefaultsVolumesEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	Fc *WorkflowTemplateSpecTemplateDefaultsVolumesFc `field:"optional" json:"fc" yaml:"fc"`
	FlexVolume *WorkflowTemplateSpecTemplateDefaultsVolumesFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	Flocker *WorkflowTemplateSpecTemplateDefaultsVolumesFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	GcePersistentDisk *WorkflowTemplateSpecTemplateDefaultsVolumesGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	GitRepo *WorkflowTemplateSpecTemplateDefaultsVolumesGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	Glusterfs *WorkflowTemplateSpecTemplateDefaultsVolumesGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	HostPath *WorkflowTemplateSpecTemplateDefaultsVolumesHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	Iscsi *WorkflowTemplateSpecTemplateDefaultsVolumesIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	Nfs *WorkflowTemplateSpecTemplateDefaultsVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	PersistentVolumeClaim *WorkflowTemplateSpecTemplateDefaultsVolumesPersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	PhotonPersistentDisk *WorkflowTemplateSpecTemplateDefaultsVolumesPhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	PortworxVolume *WorkflowTemplateSpecTemplateDefaultsVolumesPortworxVolume `field:"optional" json:"portworxVolume" yaml:"portworxVolume"`
	Projected *WorkflowTemplateSpecTemplateDefaultsVolumesProjected `field:"optional" json:"projected" yaml:"projected"`
	Quobyte *WorkflowTemplateSpecTemplateDefaultsVolumesQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	Rbd *WorkflowTemplateSpecTemplateDefaultsVolumesRbd `field:"optional" json:"rbd" yaml:"rbd"`
	ScaleIo *WorkflowTemplateSpecTemplateDefaultsVolumesScaleIo `field:"optional" json:"scaleIo" yaml:"scaleIo"`
	Secret *WorkflowTemplateSpecTemplateDefaultsVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
	Storageos *WorkflowTemplateSpecTemplateDefaultsVolumesStorageos `field:"optional" json:"storageos" yaml:"storageos"`
	VsphereVolume *WorkflowTemplateSpecTemplateDefaultsVolumesVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

