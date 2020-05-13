package writing

import (
	"encoding/json"
	"rocketlibs/preflector/writing/line_managers"
)

type Metadata struct {
	Template   string
	OutputFile string
}

func Write(jsonContent string, lineProvider line_managers.ILineProvider, lineReceiver line_managers.ILineReceiver) (err error) {
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonContent), &jsonMap)
	metadata, err := GetMetadata(jsonMap)
	if err != nil {
		return err
	}
	lineProvider.Initialize(metadata.OutputFile)
	work(lineProvider, lineReceiver)
	return nil
}

func work(lineProvider line_managers.ILineProvider, lineReceiver line_managers.ILineReceiver) (err error) {
	line, err := lineProvider.Get()
	if err != nil {
		return err
	}
	if line != "" {
		lineReceiver.Send(line)
		work(lineProvider, lineReceiver)
	}
	return nil
}
