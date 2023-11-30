package workflows_argoprojio


type WorkflowSpecArtifactGcStrategy string

const (
	// OnWorkflowCompletion.
	WorkflowSpecArtifactGcStrategy_ON_WORKFLOW_COMPLETION WorkflowSpecArtifactGcStrategy = "ON_WORKFLOW_COMPLETION"
	// OnWorkflowDeletion.
	WorkflowSpecArtifactGcStrategy_ON_WORKFLOW_DELETION WorkflowSpecArtifactGcStrategy = "ON_WORKFLOW_DELETION"
	// Never.
	WorkflowSpecArtifactGcStrategy_NEVER WorkflowSpecArtifactGcStrategy = "NEVER"
)

