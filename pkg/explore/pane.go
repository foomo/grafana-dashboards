package explore

type Pane struct {
	Datasource string `json:"datasource"`
	Queries    []any  `json:"queries"`
	Range      Range  `json:"range"`
	Compact    bool   `json:"compact"`
}
