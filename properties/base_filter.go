package properties

type Filter struct {
	Logic string       `json:"logic"`
	Items []FilterItem `json:"items"`
}

type FilterItem struct {
	Property string      `json:"property"`
	Type     string      `json:"type"`
	Logic    string      `json:"logic"`
	Value    interface{} `json:"value"`
}
