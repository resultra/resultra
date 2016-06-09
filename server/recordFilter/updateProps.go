package recordFilter

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/generic"
)

type FilterIDInterface interface {
	GetFilterID() string
	GetParentTableID() string
}

type FilterIDHeader struct {
	FilterID      string `json:"filterID"`
	ParentTableID string `json:"parentTableID"`
}

func (idHeader FilterIDHeader) GetFilterID() string {
	return idHeader.FilterID
}

func (idHeader FilterIDHeader) GetParentTableID() string {
	return idHeader.ParentTableID
}

type FilterPropUpdater interface {
	FilterIDInterface

	// Normally, UpdateProps would be named updateProps if all the property updaters were in the same
	// pacakge. However, in this case, the calculated field formula is updated in the CalcField package
	// so the function name needs to start with an upper case, so a FieldPropUpdater defined
	// in the CalcField package can be used.
	updateProps(appEngContext appengine.Context, filterForUpdate *RecordFilter) error
}

func updateFilterProps(appEngContext appengine.Context, propUpdater FilterPropUpdater) (*RecordFilter, error) {

	filterForUpdate, getErr := getFilter(appEngContext, propUpdater.GetParentTableID(), propUpdater.GetFilterID())
	if getErr != nil {
		return nil, getErr
	}

	// Do the actual property update through the FilterPropUpdater interface
	if propUpdateErr := propUpdater.updateProps(appEngContext, filterForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateFilterProps: Unable to update existing filter properties: %v", propUpdateErr)
	}

	updatedFilter, updateErr := updateExistingFilter(appEngContext, filterForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateFilterProps: error updating filter: %v", updateErr)
	}

	return updatedFilter, nil

}

type FilterRenameParams struct {
	FilterIDHeader
	Name string `json:"name"`
}

func (updateParams FilterRenameParams) updateProps(appEngContext appengine.Context, filterForUpdate *RecordFilter) error {

	sanitizedName, sanitizeErr := generic.SanitizeName(updateParams.Name)
	if sanitizeErr != nil {
		return fmt.Errorf("Update filter name: sanitize name: %v", sanitizeErr)
	}

	if validateErr := validateUnusedFilterName(appEngContext, filterForUpdate.ParentTableID, sanitizedName); validateErr != nil {
		return fmt.Errorf("Update filter name: validate unused name: error = %v", validateErr)
	}

	filterForUpdate.Name = sanitizedName

	return nil
}
