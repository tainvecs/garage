package esdao

import "encoding/json"

// QueryToString marshal elasticsearch query body to string
func QueryToString(inQuery interface{}) (string, error) {

	mar, err := json.Marshal(inQuery)
	if err != nil {
		return "", err
	}

	return string(mar), nil
}
