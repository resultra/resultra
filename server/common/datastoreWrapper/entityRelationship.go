package datastoreWrapper

// Datastore objects for entities with a parent require not only a parent key and child key, but a parent entity
// kind and child entity kind. The keys are encoded as string ID's and passed back to the client of this pacakge,
// but the entity kinds are constant for any given child entity kind. This struct should be used to define this
// constant child->parent entity relationship; this constant can then be passed to the datastoreWrapper
// functions.
type ChildParentEntityRel struct {
	ParentEntityKind string
	ChildEntityKind  string
}
