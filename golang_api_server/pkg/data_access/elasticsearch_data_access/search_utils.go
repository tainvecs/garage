package elasticsearch_data_access

import "encoding/json"

func ESQueryToString(inQuery interface{}) (string, error) {

	mar, err := json.Marshal(inQuery)
	if err != nil {
		return "", err
	}

	return string(mar), nil
}
