package tools

import (
	ps "github.com/mitchellh/go-ps"
	"servermonitor/pkg/schemas"
)

func GetProcessList() ([]schemas.Process, error) {
	processList, err := ps.Processes()
	if err != nil {
		return nil, err
	}
	var activeProcess []schemas.Process
	for _, process := range processList {
		activeProcess = append(activeProcess, schemas.Process{
			Pid:  process.Pid(),
			Name: process.Executable(),
		})

	}
	return activeProcess, nil
}
