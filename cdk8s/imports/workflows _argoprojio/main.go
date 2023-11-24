// workflows _argoprojio
package workflows _argoprojio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"workflows _argoprojio.Workflow",
		reflect.TypeOf((*Workflow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Workflow{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowProps",
		reflect.TypeOf((*WorkflowProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpec",
		reflect.TypeOf((*WorkflowSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinity",
		reflect.TypeOf((*WorkflowSpecAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityNodeAffinity",
		reflect.TypeOf((*WorkflowSpecAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*WorkflowSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*WorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*WorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*WorkflowSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAffinity",
		reflect.TypeOf((*WorkflowSpecAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*WorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*WorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*WorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAntiAffinity",
		reflect.TypeOf((*WorkflowSpecAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*WorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*WorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*WorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*WorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*WorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArguments",
		reflect.TypeOf((*WorkflowSpecArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifacts",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsArchive",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsAzure",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsGcs",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsGit",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHdfs",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttp",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsOss",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsRaw",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsS3",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsParameters",
		reflect.TypeOf((*WorkflowSpecArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsParametersValueFrom",
		reflect.TypeOf((*WorkflowSpecArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArtifactGc",
		reflect.TypeOf((*WorkflowSpecArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecArtifactRepositoryRef",
		reflect.TypeOf((*WorkflowSpecArtifactRepositoryRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecDnsConfig",
		reflect.TypeOf((*WorkflowSpecDnsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecDnsConfigOptions",
		reflect.TypeOf((*WorkflowSpecDnsConfigOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecExecutor",
		reflect.TypeOf((*WorkflowSpecExecutor)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooks",
		reflect.TypeOf((*WorkflowSpecHooks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArguments",
		reflect.TypeOf((*WorkflowSpecHooksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifacts",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsArchive",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecHooksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsAzure",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsGcs",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsGit",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHdfs",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttp",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsOss",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsRaw",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsS3",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsParameters",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsParametersValueFrom",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecHooksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHooksTemplateRef",
		reflect.TypeOf((*WorkflowSpecHooksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecHostAliases",
		reflect.TypeOf((*WorkflowSpecHostAliases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecImagePullSecrets",
		reflect.TypeOf((*WorkflowSpecImagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecMetrics",
		reflect.TypeOf((*WorkflowSpecMetrics)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecMetricsPrometheus",
		reflect.TypeOf((*WorkflowSpecMetricsPrometheus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecMetricsPrometheusCounter",
		reflect.TypeOf((*WorkflowSpecMetricsPrometheusCounter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecMetricsPrometheusGauge",
		reflect.TypeOf((*WorkflowSpecMetricsPrometheusGauge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecMetricsPrometheusHistogram",
		reflect.TypeOf((*WorkflowSpecMetricsPrometheusHistogram)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecMetricsPrometheusLabels",
		reflect.TypeOf((*WorkflowSpecMetricsPrometheusLabels)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecPodDisruptionBudget",
		reflect.TypeOf((*WorkflowSpecPodDisruptionBudget)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecPodDisruptionBudgetMaxUnavailable",
		reflect.TypeOf((*WorkflowSpecPodDisruptionBudgetMaxUnavailable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecPodDisruptionBudgetMaxUnavailable{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecPodDisruptionBudgetMinAvailable",
		reflect.TypeOf((*WorkflowSpecPodDisruptionBudgetMinAvailable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecPodDisruptionBudgetMinAvailable{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecPodDisruptionBudgetSelector",
		reflect.TypeOf((*WorkflowSpecPodDisruptionBudgetSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecPodDisruptionBudgetSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecPodDisruptionBudgetSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecPodGc",
		reflect.TypeOf((*WorkflowSpecPodGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecPodGcLabelSelector",
		reflect.TypeOf((*WorkflowSpecPodGcLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecPodGcLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecPodGcLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecPodMetadata",
		reflect.TypeOf((*WorkflowSpecPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecRetryStrategy",
		reflect.TypeOf((*WorkflowSpecRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecRetryStrategyAffinity",
		reflect.TypeOf((*WorkflowSpecRetryStrategyAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecRetryStrategyBackoff",
		reflect.TypeOf((*WorkflowSpecRetryStrategyBackoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecRetryStrategyBackoffFactor",
		reflect.TypeOf((*WorkflowSpecRetryStrategyBackoffFactor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecRetryStrategyBackoffFactor{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecRetryStrategyLimit",
		reflect.TypeOf((*WorkflowSpecRetryStrategyLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecRetryStrategyLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecSecurityContext",
		reflect.TypeOf((*WorkflowSpecSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecSecurityContextSysctls",
		reflect.TypeOf((*WorkflowSpecSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecSynchronization",
		reflect.TypeOf((*WorkflowSpecSynchronization)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecSynchronizationMutex",
		reflect.TypeOf((*WorkflowSpecSynchronizationMutex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecSynchronizationSemaphore",
		reflect.TypeOf((*WorkflowSpecSynchronizationSemaphore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecSynchronizationSemaphoreConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecSynchronizationSemaphoreConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaults",
		reflect.TypeOf((*WorkflowSpecTemplateDefaults)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsActiveDeadlineSeconds",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsActiveDeadlineSeconds)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsActiveDeadlineSeconds{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinity",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityNodeAffinity",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAffinity",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAntiAffinity",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocation",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationAzure",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationGcs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationGit",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHdfs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttp",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationOss",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationRaw",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationS3",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsArchiveLocationS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsArchiveLocationS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainer",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerEnv",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerEnvFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerEnvFromSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerEnvValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecycle",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePostStart",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePreStop",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLivenessProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLivenessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerPorts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerReadinessProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerReadinessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerResources",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerResourcesLimits",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerResourcesRequests",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSecurityContext",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainers",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersEnv",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersEnvFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersEnvFromSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecycle",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStart",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStop",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersPorts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersResources",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersResourcesLimits",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersSecurityContext",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersVolumeDevices",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetContainersVolumeMounts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetRetryStrategy",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerSetRetryStrategyRetries{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerSetVolumeMounts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerSetVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerStartupProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerStartupProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerStartupProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsContainerStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerVolumeDevices",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsContainerVolumeMounts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsContainerVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDag",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDag)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasks",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArguments",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifacts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArchive",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsAzure",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGcs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGit",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttp",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOss",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsRaw",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsParameters",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsParametersValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksContinueOn",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksContinueOn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooks",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArguments",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifacts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArchive",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzure",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGcs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGit",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttp",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOss",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsRaw",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsParameters",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsParametersValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksHooksTemplateRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksHooksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksTemplateRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksWithSequence",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksWithSequence)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsDagTasksWithSequenceCount{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsDagTasksWithSequenceEnd{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsDagTasksWithSequenceStart{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsData",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSource",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPaths",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPaths)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArchive",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArchiveTar",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsAzure",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGcs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGit",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttp",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsOss",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsRaw",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsDataTransformation",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsDataTransformation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsExecutor",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsExecutor)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsHostAliases",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsHostAliases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsHttp",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsHttpBodyFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsHttpBodyFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsHttpHeadersValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsHttpHeadersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsHttpHeadersValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsHttpHeadersValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainers",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersEnv",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersEnvFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersEnvFromSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersEnvValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecycle",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStart",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStop",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLivenessProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLivenessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersPorts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersReadinessProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersReadinessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersResources",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersResourcesLimits",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersResourcesRequests",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersSecurityContext",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersStartupProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersStartupProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersStartupProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsInitContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersVolumeDevices",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInitContainersVolumeMounts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInitContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifacts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsArchive",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsAzure",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsGcs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsGit",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHdfs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttp",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsOss",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsRaw",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsS3",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsParameters",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsParametersValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsInputsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsInputsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsMemoize",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsMemoize)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsMemoizeCache",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsMemoizeCache)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsMemoizeCacheConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsMemoizeCacheConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsMetadata",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsMetrics",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsMetrics)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsMetricsPrometheus",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsMetricsPrometheus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsMetricsPrometheusCounter",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsMetricsPrometheusCounter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsMetricsPrometheusGauge",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsMetricsPrometheusGauge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsMetricsPrometheusHistogram",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsMetricsPrometheusHistogram)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsMetricsPrometheusLabels",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsMetricsPrometheusLabels)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifacts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsArchive",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsAzure",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsGcs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsGit",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHdfs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttp",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsOss",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsRaw",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsS3",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsParameters",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsParametersValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsOutputsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsOutputsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResource",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifact",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifact)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArchive",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArchiveTar",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactAzure",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactGcs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactGit",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttp",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactOss",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactRaw",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsRetryStrategy",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsRetryStrategyAffinity",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsRetryStrategyAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsRetryStrategyBackoff",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsRetryStrategyBackoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsRetryStrategyBackoffFactor{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsRetryStrategyLimit",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsRetryStrategyLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsRetryStrategyLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScript",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScript)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptEnv",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptEnvFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptEnvFromSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptEnvValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecycle",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePostStart",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePreStop",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLivenessProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLivenessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptPorts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptReadinessProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptReadinessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptResources",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptResourcesLimits",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptResourcesRequests",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptSecurityContext",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptStartupProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptStartupProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptStartupProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptVolumeDevices",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsScriptVolumeMounts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsScriptVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSecurityContext",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSecurityContextSysctls",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecars",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecars)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsEnv",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsEnvFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsEnvFromSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsEnvValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecycle",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStart",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStop",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLivenessProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLivenessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsPorts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsReadinessProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsReadinessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsResources",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsResourcesLimits",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsResourcesRequests",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsSecurityContext",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsStartupProbe",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsStartupProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsStartupProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsVolumeDevices",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSidecarsVolumeMounts",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSidecarsVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSuspend",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSuspend)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSynchronization",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSynchronization)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSynchronizationMutex",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSynchronizationMutex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSynchronizationSemaphore",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSynchronizationSemaphore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsSynchronizationSemaphoreConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsSynchronizationSemaphoreConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsTolerations",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumes",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesAwsElasticBlockStore",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesAzureDisk",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesAzureFile",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesCephfs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesCephfsSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesCinder",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesCinderSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesConfigMapItems",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesCsi",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesDownwardApi",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesDownwardApiItems",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEmptyDir",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeral",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesEphemeral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplate",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpec",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesFc",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesFlexVolume",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesFlocker",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesGcePersistentDisk",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesGitRepo",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesGlusterfs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesHostPath",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesIscsi",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesIscsiSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesNfs",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesPersistentVolumeClaim",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesPhotonPersistentDisk",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesPortworxVolume",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesProjected",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesProjectedSources",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesProjectedSourcesSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesQuobyte",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesRbd",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesRbdSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesScaleIo",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesScaleIoSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesSecret",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesSecretItems",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesStorageos",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesStorageosSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplateDefaultsVolumesVsphereVolume",
		reflect.TypeOf((*WorkflowSpecTemplateDefaultsVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplates",
		reflect.TypeOf((*WorkflowSpecTemplates)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesActiveDeadlineSeconds",
		reflect.TypeOf((*WorkflowSpecTemplatesActiveDeadlineSeconds)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesActiveDeadlineSeconds{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinity",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityNodeAffinity",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAffinity",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAntiAffinity",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocation",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationAzure",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationGcs",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationGit",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHdfs",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttp",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationOss",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationRaw",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationS3",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesArchiveLocationS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesArchiveLocationS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainer",
		reflect.TypeOf((*WorkflowSpecTemplatesContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerEnv",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerEnvFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerEnvFromSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerEnvValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecycle",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePostStart",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePreStop",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLivenessProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLivenessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerPorts",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerReadinessProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerReadinessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerResources",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerResourcesLimits",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerResourcesRequests",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSecurityContext",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSet",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainers",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersEnv",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersEnvFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersEnvFromSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersEnvValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecycle",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePostStart",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePreStop",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLivenessProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLivenessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersPorts",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersReadinessProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersReadinessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersResources",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersResourcesLimits",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersResourcesRequests",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersSecurityContext",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersStartupProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersStartupProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersStartupProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersVolumeDevices",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetContainersVolumeMounts",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetRetryStrategy",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetRetryStrategyRetries",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetRetryStrategyRetries)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerSetRetryStrategyRetries{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerSetVolumeMounts",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerSetVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerStartupProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerStartupProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerStartupProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesContainerStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerVolumeDevices",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesContainerVolumeMounts",
		reflect.TypeOf((*WorkflowSpecTemplatesContainerVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDag",
		reflect.TypeOf((*WorkflowSpecTemplatesDag)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasks",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArguments",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifacts",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsArchive",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsAzure",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsGcs",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsGit",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfs",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttp",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsOss",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsRaw",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsParameters",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsParametersValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksContinueOn",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksContinueOn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooks",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArguments",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifacts",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArchive",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsAzure",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGcs",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGit",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfs",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttp",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOss",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsRaw",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsParameters",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsParametersValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksHooksTemplateRef",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksHooksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksTemplateRef",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksWithSequence",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksWithSequence)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksWithSequenceCount",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksWithSequenceCount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesDagTasksWithSequenceCount{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksWithSequenceEnd",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksWithSequenceEnd)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesDagTasksWithSequenceEnd{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesDagTasksWithSequenceStart",
		reflect.TypeOf((*WorkflowSpecTemplatesDagTasksWithSequenceStart)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesDagTasksWithSequenceStart{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesData",
		reflect.TypeOf((*WorkflowSpecTemplatesData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSource",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPaths",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPaths)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsArchive",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsArchiveTar",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsArtifactGc",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsAzure",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsGcs",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsGit",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHdfs",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttp",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsOss",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsRaw",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsS3",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataSourceArtifactPathsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesDataSourceArtifactPathsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesDataTransformation",
		reflect.TypeOf((*WorkflowSpecTemplatesDataTransformation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesExecutor",
		reflect.TypeOf((*WorkflowSpecTemplatesExecutor)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesHostAliases",
		reflect.TypeOf((*WorkflowSpecTemplatesHostAliases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesHttp",
		reflect.TypeOf((*WorkflowSpecTemplatesHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesHttpBodyFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesHttpBodyFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesHttpHeadersValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesHttpHeadersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesHttpHeadersValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesHttpHeadersValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainers",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersEnv",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersEnvFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersEnvFromSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersEnvValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecycle",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePostStart",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePreStop",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLivenessProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLivenessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersPorts",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersReadinessProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersReadinessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersResources",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersResourcesLimits",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersResourcesRequests",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersSecurityContext",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersStartupProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersStartupProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersStartupProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesInitContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersVolumeDevices",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInitContainersVolumeMounts",
		reflect.TypeOf((*WorkflowSpecTemplatesInitContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputs",
		reflect.TypeOf((*WorkflowSpecTemplatesInputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifacts",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsArchive",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecTemplatesInputsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecTemplatesInputsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecTemplatesInputsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsAzure",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsGcs",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsGit",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHdfs",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttp",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsOss",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsRaw",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsS3",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsParameters",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsParametersValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesInputsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesInputsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesMemoize",
		reflect.TypeOf((*WorkflowSpecTemplatesMemoize)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesMemoizeCache",
		reflect.TypeOf((*WorkflowSpecTemplatesMemoizeCache)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesMemoizeCacheConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplatesMemoizeCacheConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesMetadata",
		reflect.TypeOf((*WorkflowSpecTemplatesMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesMetrics",
		reflect.TypeOf((*WorkflowSpecTemplatesMetrics)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesMetricsPrometheus",
		reflect.TypeOf((*WorkflowSpecTemplatesMetricsPrometheus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesMetricsPrometheusCounter",
		reflect.TypeOf((*WorkflowSpecTemplatesMetricsPrometheusCounter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesMetricsPrometheusGauge",
		reflect.TypeOf((*WorkflowSpecTemplatesMetricsPrometheusGauge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesMetricsPrometheusHistogram",
		reflect.TypeOf((*WorkflowSpecTemplatesMetricsPrometheusHistogram)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesMetricsPrometheusLabels",
		reflect.TypeOf((*WorkflowSpecTemplatesMetricsPrometheusLabels)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputs",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifacts",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsArchive",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecTemplatesOutputsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecTemplatesOutputsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecTemplatesOutputsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsAzure",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsGcs",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsGit",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHdfs",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttp",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsOss",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsRaw",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsS3",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsParameters",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsParametersValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesOutputsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesOutputsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResource",
		reflect.TypeOf((*WorkflowSpecTemplatesResource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifact",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifact)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactArchive",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactArchiveTar",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactArtifactGc",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactArtifactory",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactAzure",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactGcs",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactGit",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactGitPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactGitUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHdfs",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttp",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCert",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactOss",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactOssAccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactOssLifecycleRule",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactOssSecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactRaw",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactS3",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactS3AccessKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactS3CaSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactS3EncryptionOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesResourceManifestFromArtifactS3SecretKeySecret",
		reflect.TypeOf((*WorkflowSpecTemplatesResourceManifestFromArtifactS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesRetryStrategy",
		reflect.TypeOf((*WorkflowSpecTemplatesRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesRetryStrategyAffinity",
		reflect.TypeOf((*WorkflowSpecTemplatesRetryStrategyAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesRetryStrategyBackoff",
		reflect.TypeOf((*WorkflowSpecTemplatesRetryStrategyBackoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesRetryStrategyBackoffFactor",
		reflect.TypeOf((*WorkflowSpecTemplatesRetryStrategyBackoffFactor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesRetryStrategyBackoffFactor{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesRetryStrategyLimit",
		reflect.TypeOf((*WorkflowSpecTemplatesRetryStrategyLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesRetryStrategyLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScript",
		reflect.TypeOf((*WorkflowSpecTemplatesScript)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptEnv",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptEnvFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptEnvFromSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptEnvValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecycle",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePostStart",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePreStop",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLivenessProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLivenessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptPorts",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptReadinessProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptReadinessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptResources",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptResourcesLimits",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptResourcesRequests",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptSecurityContext",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptStartupProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptStartupProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptStartupProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesScriptStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptVolumeDevices",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesScriptVolumeMounts",
		reflect.TypeOf((*WorkflowSpecTemplatesScriptVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSecurityContext",
		reflect.TypeOf((*WorkflowSpecTemplatesSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecTemplatesSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSecurityContextSysctls",
		reflect.TypeOf((*WorkflowSpecTemplatesSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecars",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecars)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsEnv",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsEnvFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsEnvFromSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsEnvValueFrom",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecycle",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePostStart",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePreStop",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLivenessProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLivenessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsPorts",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsReadinessProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsReadinessProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsResources",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsResourcesLimits",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsResourcesRequests",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsSecurityContext",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsStartupProbe",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsStartupProbeExec",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsStartupProbeGrpc",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesSidecarsStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsVolumeDevices",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSidecarsVolumeMounts",
		reflect.TypeOf((*WorkflowSpecTemplatesSidecarsVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSuspend",
		reflect.TypeOf((*WorkflowSpecTemplatesSuspend)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSynchronization",
		reflect.TypeOf((*WorkflowSpecTemplatesSynchronization)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSynchronizationMutex",
		reflect.TypeOf((*WorkflowSpecTemplatesSynchronizationMutex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSynchronizationSemaphore",
		reflect.TypeOf((*WorkflowSpecTemplatesSynchronizationSemaphore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesSynchronizationSemaphoreConfigMapKeyRef",
		reflect.TypeOf((*WorkflowSpecTemplatesSynchronizationSemaphoreConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesTolerations",
		reflect.TypeOf((*WorkflowSpecTemplatesTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumes",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesAwsElasticBlockStore",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesAzureDisk",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesAzureFile",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesCephfs",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesCephfsSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesCinder",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesCinderSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesConfigMapItems",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesCsi",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesDownwardApi",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesDownwardApiItems",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesEmptyDir",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesEmptyDirSizeLimit",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesEmptyDirSizeLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesVolumesEmptyDirSizeLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesEphemeral",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesEphemeral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplate",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpec",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesFc",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesFlexVolume",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesFlocker",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesGcePersistentDisk",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesGitRepo",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesGlusterfs",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesHostPath",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesIscsi",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesIscsiSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesNfs",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesPersistentVolumeClaim",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesPhotonPersistentDisk",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesPortworxVolume",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesProjected",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesProjectedSources",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesProjectedSourcesSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesQuobyte",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesRbd",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesRbdSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesScaleIo",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesScaleIoSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesSecret",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesSecretItems",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesStorageos",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesStorageosSecretRef",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTemplatesVolumesVsphereVolume",
		reflect.TypeOf((*WorkflowSpecTemplatesVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTolerations",
		reflect.TypeOf((*WorkflowSpecTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecTtlStrategy",
		reflect.TypeOf((*WorkflowSpecTtlStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumeClaimGc",
		reflect.TypeOf((*WorkflowSpecVolumeClaimGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplates",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplates)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesSpec",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplatesSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesSpecDataSource",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplatesSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesSpecDataSourceRef",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplatesSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesSpecResources",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplatesSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecVolumeClaimTemplatesSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecVolumeClaimTemplatesSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesSpecSelector",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplatesSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesSpecSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplatesSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesStatus",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplatesStatus)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecVolumeClaimTemplatesStatusAllocatedResources{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesStatusCapacity",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplatesStatusCapacity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecVolumeClaimTemplatesStatusCapacity{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumeClaimTemplatesStatusConditions",
		reflect.TypeOf((*WorkflowSpecVolumeClaimTemplatesStatusConditions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumes",
		reflect.TypeOf((*WorkflowSpecVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesAwsElasticBlockStore",
		reflect.TypeOf((*WorkflowSpecVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesAzureDisk",
		reflect.TypeOf((*WorkflowSpecVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesAzureFile",
		reflect.TypeOf((*WorkflowSpecVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesCephfs",
		reflect.TypeOf((*WorkflowSpecVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesCephfsSecretRef",
		reflect.TypeOf((*WorkflowSpecVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesCinder",
		reflect.TypeOf((*WorkflowSpecVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesCinderSecretRef",
		reflect.TypeOf((*WorkflowSpecVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesConfigMap",
		reflect.TypeOf((*WorkflowSpecVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesConfigMapItems",
		reflect.TypeOf((*WorkflowSpecVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesCsi",
		reflect.TypeOf((*WorkflowSpecVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*WorkflowSpecVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesDownwardApi",
		reflect.TypeOf((*WorkflowSpecVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesDownwardApiItems",
		reflect.TypeOf((*WorkflowSpecVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*WorkflowSpecVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecVolumesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecVolumesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecVolumesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesEmptyDir",
		reflect.TypeOf((*WorkflowSpecVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecVolumesEmptyDirSizeLimit",
		reflect.TypeOf((*WorkflowSpecVolumesEmptyDirSizeLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecVolumesEmptyDirSizeLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesEphemeral",
		reflect.TypeOf((*WorkflowSpecVolumesEphemeral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplate",
		reflect.TypeOf((*WorkflowSpecVolumesEphemeralVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpec",
		reflect.TypeOf((*WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesFc",
		reflect.TypeOf((*WorkflowSpecVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesFlexVolume",
		reflect.TypeOf((*WorkflowSpecVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*WorkflowSpecVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesFlocker",
		reflect.TypeOf((*WorkflowSpecVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesGcePersistentDisk",
		reflect.TypeOf((*WorkflowSpecVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesGitRepo",
		reflect.TypeOf((*WorkflowSpecVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesGlusterfs",
		reflect.TypeOf((*WorkflowSpecVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesHostPath",
		reflect.TypeOf((*WorkflowSpecVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesIscsi",
		reflect.TypeOf((*WorkflowSpecVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesIscsiSecretRef",
		reflect.TypeOf((*WorkflowSpecVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesNfs",
		reflect.TypeOf((*WorkflowSpecVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesPersistentVolumeClaim",
		reflect.TypeOf((*WorkflowSpecVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesPhotonPersistentDisk",
		reflect.TypeOf((*WorkflowSpecVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesPortworxVolume",
		reflect.TypeOf((*WorkflowSpecVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesProjected",
		reflect.TypeOf((*WorkflowSpecVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesProjectedSources",
		reflect.TypeOf((*WorkflowSpecVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*WorkflowSpecVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*WorkflowSpecVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*WorkflowSpecVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*WorkflowSpecVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*WorkflowSpecVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*WorkflowSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflows _argoprojio.WorkflowSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesProjectedSourcesSecret",
		reflect.TypeOf((*WorkflowSpecVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*WorkflowSpecVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*WorkflowSpecVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesQuobyte",
		reflect.TypeOf((*WorkflowSpecVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesRbd",
		reflect.TypeOf((*WorkflowSpecVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesRbdSecretRef",
		reflect.TypeOf((*WorkflowSpecVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesScaleIo",
		reflect.TypeOf((*WorkflowSpecVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesScaleIoSecretRef",
		reflect.TypeOf((*WorkflowSpecVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesSecret",
		reflect.TypeOf((*WorkflowSpecVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesSecretItems",
		reflect.TypeOf((*WorkflowSpecVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesStorageos",
		reflect.TypeOf((*WorkflowSpecVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesStorageosSecretRef",
		reflect.TypeOf((*WorkflowSpecVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecVolumesVsphereVolume",
		reflect.TypeOf((*WorkflowSpecVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecWorkflowMetadata",
		reflect.TypeOf((*WorkflowSpecWorkflowMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecWorkflowMetadataLabelsFrom",
		reflect.TypeOf((*WorkflowSpecWorkflowMetadataLabelsFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflows _argoprojio.WorkflowSpecWorkflowTemplateRef",
		reflect.TypeOf((*WorkflowSpecWorkflowTemplateRef)(nil)).Elem(),
	)
}
