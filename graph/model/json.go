package model

import (
	"encoding/json"
	"io"
)

// JSON is a custom scalar type representing JSON data.
type JSON map[string]interface{}

// MarshalGQL converts the Go type to its GraphQL representation.
func (j JSON) MarshalGQL(w io.Writer) {
	data, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	w.Write(data)
}

// UnmarshalGQL converts the GraphQL representation to the Go type.
func (j *JSON) UnmarshalGQL(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, j)
}
