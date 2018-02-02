package main

type specification struct {
	Name       string `json:"name"`
	Contains   bool   `json:"contains"`
	Pattern    string `json:"pattern"`
	Filename   string `json:"filename"`
	Error      string `json:"error"`
	Resolution string `json:"resolution"`
}
