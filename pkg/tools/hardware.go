package tools

import (
	ps "github.com/mitchellh/go-ps"
	"servermonitor/pkg/types"
)

func GetProcessList() ([]types.Process, error) {
	processList, err := ps.Processes()
	if err != nil {
		return nil, err
	}
	var activeProcess []types.Process
	for _, process := range processList {
		activeProcess = append(activeProcess, types.Process{
			Pid:  process.Pid(),
			Name: process.Executable(),
		})

	}
	return activeProcess, nil
}
