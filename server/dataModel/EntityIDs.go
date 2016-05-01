package dataModel

import (
	"resultra/datasheet/server/generic/datastoreWrapper"
)

// The entity IDs defined here are used as parent entities by many of the packages. Most of these
// packages only depend on the parent for their entity ID, so defining the parent entity IDs here
// allows the packages with child entities to avoid a circular dependency.

const DatabaseEntityKind string = "Database"

const TableEntityKind string = "Table"

const LayoutEntityKind string = "Layout"
const FormEntityKind string = "Form"

const DashboardEntityKind string = "Dashboard"

var TableChildParentEntityRel = datastoreWrapper.ChildParentEntityRel{
	ParentEntityKind: DatabaseEntityKind,
	ChildEntityKind:  TableEntityKind}
