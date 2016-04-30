package datastoreWrapper

// UniqueID stores the opaque/encoded database keys for an entity.
type UniqueID struct {
	ParentID string `json:"parentID"`
	ObjectID string `json:"objectID"`
}

type UniqueIDHeader struct {
	UniqueID UniqueID `json:"uniqueID"`
}

func NewUniqueIDHeader(parentID string, objectID string) UniqueIDHeader {
	uniqueID := UniqueID{parentID, objectID}
	uniqueIDHeader := UniqueIDHeader{uniqueID}
	return uniqueIDHeader
}

func (header UniqueIDHeader) GetUniqueID() UniqueID {
	return header.UniqueID
}

type UniqueIDInterface interface {
	GetUniqueID() UniqueID
}

type UniqueRootID struct {
	ObjectID string `json:"objectID"`
}

type UniqueRootIDInterface interface {
	GetUniqueRootID() UniqueRootID
}

type UniqueRootIDHeader struct {
	UniqueID UniqueRootID `json:"uniqueID"`
}

func NewUniqueRootIDHeader(objectID string) UniqueRootIDHeader {
	uniqueID := UniqueRootID{objectID}
	uniqueIDHeader := UniqueRootIDHeader{uniqueID}
	return uniqueIDHeader
}

func (header UniqueRootIDHeader) GetUniqueRootID() UniqueRootID {
	return header.UniqueID
}
