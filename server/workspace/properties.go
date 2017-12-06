package workspace

type WorkspaceProperties struct {
	AllowUserRegistration bool `json:"allowUserRegistration"`
}

func newDefaultWorkspaceProperties() WorkspaceProperties {
	props := WorkspaceProperties{
		AllowUserRegistration: true}
	return props
}
