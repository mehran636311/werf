package main

import (
	"encoding/json"
	"fmt"
)

const (
	nameUsage = "Usage"
	nameVersion = "Version"
	nameChannel= "Channel"
	nameBackend ="Backend"
	nameCI ="CI"
	nameExecutor="Executor"
	nameOS="OS"
	nameProjectType="Project type"

	usageLocal = "Local dev"
	usageCI = "CI"
	version11 = "1.1"
	version12 = "1.2"
	channelAlpha= "Alpha"
	channelBeta= "Beta"
	channelEA= "Early-access"
	channelStable= "Stable"
	channelRockSolid= "Rock-solid"
	backendDocker="Docker"
	backendBuildah="Buildah"
	ciGL="GitLab CI/CD"
	ciGHActions="GitHub Actions"
	executorShell="Shell"
	executorDocker="Docker"
	executorKubernetes="Kubernetes"
	osLinux="Linux"
	osWindows="Windows"
	osMacOS="Mac OS"
	projectTypeSimple="Simple"
	projectTypeMonorepo="Monorepo"
	projectTypeMultiRepo="Multi-repo"
)

var (
	valuesUsage   = []string{usageLocal, usageCI}
	valuesVersion = []string{version11, version12}
	valuesChannel = []string{channelAlpha, channelBeta, channelEA, channelStable, channelRockSolid}
	valuesBackend  =[]string{backendDocker,backendBuildah}
	valuesCI       =[]string{ciGL, ciGHActions}
	valuesExecutor =[]string{executorShell, executorDocker, executorKubernetes}
	valuesOS       =[]string{osLinux,osWindows,osMacOS}
	valuesProjectType=[]string{projectTypeSimple,projectTypeMonorepo,projectTypeMultiRepo}
)

func main()  {
	valuesMap := map[string]interface{}{}
	for _, usage := range valuesUsage {
		usageMap := map[string]interface{}{}
		for _, version := range valuesVersion {
			// инструкции CI для 1.1 не будет
			if version == version11 && usage == usageCI {
				continue
			}

			versionMap := map[string]interface{}{}
			for _, channel := range valuesChannel {
				channelMap := map[string]interface{}{}
				for _, backend := range valuesBackend {
					backendMap := map[string]interface{}{}
					for _, ci := range valuesCI {
						ciMap := map[string]interface{}{}
						for _, executor := range valuesExecutor {
							executorMap := map[string]interface{}{}
							for _, os := range valuesOS {
								osMap := map[string]interface{}{}
								for _, projectType := range valuesProjectType {
									projectTypeMap := map[string]interface{}{}
									osMap[projectType] = projectTypeMap
								}

								executorMap[os] = osMap
							}

							ciMap[executor] = executorMap
						}

						backendMap[ci] = ciMap
					}

					channelMap[backend] = backendMap
				}

				versionMap[channel] = channelMap
			}

			usageMap[version] = versionMap
		}

		valuesMap[usage] = usageMap
	}

	jsonString, _ := json.MarshalIndent(valuesMap, "", "    ")
	fmt.Println(string(jsonString))
}
