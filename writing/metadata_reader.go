package writing

import "errors"

func GetMetadata(jsonMap map[string]interface{}) (metadata Metadata, err error) {
	metadataMap := jsonMap["Metadata"].(map[string]interface{})
	metadata = Metadata{
		OutputFile: metadataMap["OutputFile"].(string),
		Template:   metadataMap["Template"].(string),
	}

	if metadata.Template == "" {
		return metadata, errors.New("No template was specified in the metadata section of your json. Template is mandatory")
	}
	return metadata, nil
}
