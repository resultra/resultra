package form

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

const layoutEntityKind string = "Layout"

type Layout struct {
	Name string `json:"name"`
}

type LayoutRef struct {
	LayoutID string `json"layoutID"`
	Layout   Layout `json"layout"`
}

func NewLayout(appEngContext appengine.Context, layoutName string) (string, error) {

	sanitizedLayoutName, sanitizeErr := generic.SanitizeName(layoutName)
	if sanitizeErr != nil {
		return "", sanitizeErr
	}

	var newLayout = Layout{sanitizedLayoutName}
	layoutID, insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, layoutEntityKind, &newLayout)
	if insertErr != nil {
		return "", insertErr
	}

	// TODO - verify IntID != 0
	log.Printf("NewLayout: Created new Layout: id= %v, name='%v'", layoutID, sanitizedLayoutName)

	return layoutID, nil

}

func GetAllLayoutRefs(appEngContext appengine.Context) ([]LayoutRef, error) {

	var allLayouts []Layout
	ids, err := datastoreWrapper.GetAllRootEntities(appEngContext, layoutEntityKind, &allLayouts)
	if err != nil {
		return nil, fmt.Errorf("GetAllLayouts: Unable to retrieve layouts from datastore: datastore error =%v", err)
	}

	layoutRefs := make([]LayoutRef, len(allLayouts))
	for i, currLayout := range allLayouts {
		layoutID := ids[i]
		layoutRefs[i] = LayoutRef{layoutID, currLayout}
	}
	return layoutRefs, nil

}

type GetLayoutParams struct {
	// TODO - More fields will go here once a layout is
	// tied to a database table
	LayoutID string `json:"layoutID"`
}

func GetLayoutRef(appEngContext appengine.Context, layoutParams GetLayoutParams) (*LayoutRef, error) {

	getLayout := Layout{}
	getErr := datastoreWrapper.GetEntity(appEngContext, layoutParams.LayoutID, &getLayout)
	if getErr != nil {
		return nil, fmt.Errorf("Can't get layout: Error retrieving existing layout: params=%+v, err = %v", layoutParams, getErr)
	}

	return &LayoutRef{layoutParams.LayoutID, getLayout}, nil

}
