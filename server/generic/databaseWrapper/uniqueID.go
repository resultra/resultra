package databaseWrapper

import "github.com/twinj/uuid"

func GlobalUniqueID() string {
	return uuid.NewV4().String()
}