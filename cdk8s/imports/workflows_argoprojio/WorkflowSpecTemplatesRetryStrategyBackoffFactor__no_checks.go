//go:build no_runtime_type_checking

package workflows_argoprojio

// Building without runtime type checking enabled, so all the below just return nil

func validateWorkflowSpecTemplatesRetryStrategyBackoffFactor_FromNumberParameters(value *float64) error {
	return nil
}

func validateWorkflowSpecTemplatesRetryStrategyBackoffFactor_FromStringParameters(value *string) error {
	return nil
}

