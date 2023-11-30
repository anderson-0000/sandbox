//go:build !no_runtime_type_checking

package cronworkflow_argoprojio

import (
	"fmt"
)

func validateCronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocketPort_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateCronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbeTcpSocketPort_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

