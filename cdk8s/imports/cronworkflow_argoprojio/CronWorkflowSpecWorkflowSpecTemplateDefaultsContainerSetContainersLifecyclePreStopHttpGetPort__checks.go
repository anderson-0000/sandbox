//go:build !no_runtime_type_checking

package cronworkflow_argoprojio

import (
	"fmt"
)

func validateCronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

