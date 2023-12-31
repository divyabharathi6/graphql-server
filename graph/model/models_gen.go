// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Data struct {
	ID          string `json:"_id" bson:"_id"`
	UserID      string `json:"userId"`
	AccountID   string `json:"accountId"`
	ProjectID   string `json:"projectId"`
	EventName   string `json:"eventName"`
	ProjectName string `json:"projectName"`
	Properties  JSON   `json:"properties"`
}

type NewData struct {
	UserID      string `json:"userId"`
	AccountID   string `json:"accountId"`
	ProjectID   string `json:"projectId"`
	EventName   string `json:"eventName"`
	ProjectName string `json:"projectName"`
	Properties  JSON   `json:"properties"`
}
