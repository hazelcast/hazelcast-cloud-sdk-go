package models

import "io"

type ListUploadedArtifactsInput struct {
	ClusterId string
}

type UploadArtifactInput struct {
	ClusterId string
	FileName  string
	Content   io.Reader
}

type UploadArtifactArgs struct {
	ClusterId string
}

type DeleteArtifactInput struct {
	ClusterId       string
	CustomClassesId string
}

type DownloadArtifactInput struct {
	ClusterId       string
	CustomClassesId string
}

type UploadedArtifact struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type UploadedArtifactLink struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
