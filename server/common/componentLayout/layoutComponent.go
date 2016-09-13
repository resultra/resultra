package componentLayout

type LayoutComponentCol struct {
	ComponentIDs []string `json:"componentIDs"`
}

type LayoutComponentRow struct {
	Columns []LayoutComponentCol `json:"columns"`
}

type ComponentLayout []LayoutComponentRow
