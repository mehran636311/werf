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
	optionUsage := map[string]interface{}{}
	optionValuesUsage := map[string]interface{}{}
	for _, usage := range valuesUsage {
		optionVersion := map[string]interface{}{}
		optionValuesVersion := map[string]interface{}{}
		for _, version := range valuesVersion {
			// инструкции CI для 1.1 не будет
			if version == version11 && usage == usageCI {
				continue
			}

			optionChannel := map[string]interface{}{}
			optionValuesChannel := map[string]interface{}{}
			for _, channel := range valuesChannel {

				optionBackend := map[string]interface{}{}
				optionValuesBackend := map[string]interface{}{}
				for _, backend := range valuesBackend {

					optionCI := map[string]interface{}{}
					optionValuesCI := map[string]interface{}{}
					for _, ci := range valuesCI {
						optionExecutor := generateOption(nameExecutor, valuesExecutor, func() map[string]interface{} {
							return generateOption(nameOS, valuesOS, func() map[string]interface{} {
								return generateOption(nameProjectType, valuesProjectType, nil)
							})
						})

						optionValuesCI[ci] = optionExecutor
					}

					optionCI["option"] = nameCI
					optionCI["values"] = optionValuesCI
					optionValuesBackend[backend] = optionCI
				}

				optionBackend["option"] = nameBackend
				optionBackend["values"] = optionValuesBackend
				optionValuesChannel[channel] = optionBackend
			}

			optionChannel["option"] = nameChannel
			optionChannel["values"] = optionValuesChannel
			optionValuesVersion[version] = optionChannel
		}

		optionVersion["option"] = nameVersion
		optionVersion["values"] = optionValuesVersion
		optionValuesUsage[usage] = optionVersion
	}
	optionUsage["option"] = nameUsage
	optionUsage["values"] = optionValuesUsage

	jsonString, _ := json.MarshalIndent(optionUsage, "", "    ")
	fmt.Println(string(jsonString))
}

func generateOption(name string, values []string, optionValuesFunc func() map[string]interface{}) map[string]interface{} {
	option := map[string]interface{}{}
	optionValues := map[string]interface{}{}

	for _, value := range values {
		if optionValuesFunc != nil {
			optionValues[value] = optionValuesFunc()
		} else {
			optionValues[value] = nil
		}
	}

	option["option"] = name
	option["values"] = optionValues
	return option
}
