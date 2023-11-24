//go:build no_runtime_type_checking

package workflows _argoprojio

// Building without runtime type checking enabled, so all the below just return nil

func validateWorkflow_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateWorkflow_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWorkflow_ManifestParameters(props *WorkflowProps) error {
	return nil
}

func validateWorkflow_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewWorkflowParameters(scope constructs.Construct, id *string, props *WorkflowProps) error {
	return nil
}

