package deploy

import "fmt"

type LintOptions struct {
	ProjectDir   string
	Values       []string
	SecretValues []string
	Set          []string
}

func RunLint(opts LintOptions) error {
	if debug() {
		fmt.Printf("Lint options: %#v\n", opts)
	}

	s, err := getOptionalSecret(opts.ProjectDir, opts.SecretValues)
	if err != nil {
		return fmt.Errorf("cannot get project secret: %s", err)
	}

	serviceValues, err := GetServiceValues("PROJECT_NAME", "REPO", "NAMESPACE", "DOCKER_TAG", nil, nil, ServiceValuesOptions{
		Fake:            true,
		WithoutRegistry: true,
	})

	dappChart, err := getDappChart(opts.ProjectDir, s, opts.Values, opts.SecretValues, opts.Set, serviceValues)
	if err != nil {
		return err
	}

	return dappChart.Lint()
}
