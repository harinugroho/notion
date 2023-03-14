package properties

type Filter struct {
	Logic string
	Items []FilterItem
}

type FilterItem struct {
	Property string
	Type     string
	Logic    string
	Value    interface{}
}
