package utils

import "encoding/base64"

func ConvertStringToJSON(env_details string) []byte {
	decoded_json, err := base64.StdEncoding.DecodeString(env_details)
	if err != nil {
		panic(err)
	}

	return decoded_json
}
