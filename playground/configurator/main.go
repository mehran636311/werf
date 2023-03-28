package main

import (
	"encoding/json"
	"fmt"
)

const (
	nameUsage       = "Usage"
	nameVersion     = "Version"
	nameChannel     = "Channel"
	nameBackend     = "Backend"
	nameCI          = "CI"
	nameExecutor    = "Executor"
	nameOS          = "OS"
	nameProjectType = "Project type"

	usageLocal           = "Local dev"
	usageCI              = "CI"
	version11            = "1.1"
	version12            = "1.2"
	channelAlpha         = "Alpha"
	channelBeta          = "Beta"
	channelEA            = "Early-access"
	channelStable        = "Stable"
	channelRockSolid     = "Rock-solid"
	backendDocker        = "Docker"
	backendBuildah       = "Buildah"
	ciGL                 = "GitLab CI/CD"
	ciGHActions          = "GitHub Actions"
	executorShell        = "Shell"
	executorDocker       = "Docker"
	executorKubernetes   = "Kubernetes"
	osLinux              = "Linux"
	osWindows            = "Windows"
	osMacOS              = "Mac OS"
	projectTypeSimple    = "Simple"
	projectTypeMonorepo  = "Monorepo"
	projectTypeMultiRepo = "Multi-repo"
)

var (
	valuesUsage       = []string{usageLocal, usageCI}
	valuesVersion     = []string{version11, version12}
	valuesChannel     = []string{channelAlpha, channelBeta, channelEA, channelStable, channelRockSolid}
	valuesBackend     = []string{backendDocker, backendBuildah}
	valuesCI          = []string{ciGL, ciGHActions}
	valuesExecutor    = []string{executorShell, executorDocker, executorKubernetes}
	valuesOS          = []string{osLinux, osWindows, osMacOS}
	valuesProjectType = []string{projectTypeSimple, projectTypeMonorepo, projectTypeMultiRepo}
)

func main() {
	optionUsage := generateOption(nameUsage, valuesUsage, func(vUsage string) interface{} {
		return generateOption(nameVersion, valuesVersion, func(vVersion string) interface{} {
			if vUsage == usageLocal && vVersion == version11 {
				return nil
			}

			return generateOption(nameChannel, valuesChannel, func(string) interface{} {
				return generateOption(nameBackend, valuesBackend, func(string) interface{} {
					return generateOption(nameCI, valuesCI, func(string) interface{} {
						return generateOption(nameExecutor, valuesExecutor, func(string) interface{} {
							return generateOption(nameOS, valuesOS, func(string) interface{} {
								return generateOption(nameProjectType, valuesProjectType, nil)
							})
						})
					})
				})
			})
		})
	})

	jsonString, _ := json.MarshalIndent(optionUsage, "", "    ")
	fmt.Println(string(jsonString))
}

func generateOption(name string, values []string, optionValuesFunc func(value string) interface{}) interface{} {
	option := map[string]interface{}{}
	optionValues := map[string]interface{}{}

	for _, value := range values {
		if optionValuesFunc != nil {
			res := optionValuesFunc(value)
			if res == nil {
				continue
			}

			optionValues[value] = optionValuesFunc(value)
		} else {
			optionValues[value] = ":END:"
		}
	}

	option["option"] = name
	option["values"] = optionValues
	return option
}
