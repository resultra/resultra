package form

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"log"
	"resultra/datasheet/server/dataModel"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
)

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
	layoutID, insertErr := datastoreWrapper.InsertNewRootEntity(appEngContext, dataModel.LayoutEntityKind, &newLayout)
	if insertErr != nil {
		return "", insertErr
	}

	// TODO - verify IntID != 0
	log.Printf("NewLayout: Created new Layout: id= %v, name='%v'", layoutID, sanitizedLayoutName)

	return layoutID, nil

}

func GetAllLayoutRefs(appEngContext appengine.Context) ([]LayoutRef, error) {
	var allLayouts []Layout
	layoutQuery := datastore.NewQuery(dataModel.LayoutEntityKind)
	keys, err := layoutQuery.GetAll(appEngContext, &allLayouts)

	if err != nil {
		return nil, fmt.Errorf("GetAllLayouts: Unable to retrieve layouts from datastore: datastore error =%v", err)
	}

	layoutRefs := make([]LayoutRef, len(allLayouts))
	for i, currLayout := range allLayouts {
		layoutKey := keys[i]
		layoutID, encodeErr := datastoreWrapper.EncodeUniqueEntityIDToStr(layoutKey)
		if encodeErr != nil {
			return nil, fmt.Errorf("Failed to encode unique ID for layout: key=%+v, encode err=%v", layoutKey, encodeErr)
		}
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
	getErr := datastoreWrapper.GetRootEntity(appEngContext, dataModel.LayoutEntityKind, layoutParams.LayoutID, &getLayout)
	if getErr != nil {
		return nil, fmt.Errorf("Can't get layout: Error retrieving existing layout: params=%+v, err = %v", layoutParams, getErr)
	}

	return &LayoutRef{layoutParams.LayoutID, getLayout}, nil

}
