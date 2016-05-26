package recordFilter

import (
	"appengine"
	"fmt"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/datastoreWrapper"
	"resultra/datasheet/server/table"
)

type FilterIDInterface interface {
	GetFilterID() string
}

type FilterIDHeader struct {
	FilterID string `json:"filterID"`
}

func (idHeader FilterIDHeader) GetFilterID() string {
	return idHeader.FilterID
}

type FilterPropUpdater interface {
	FilterIDInterface

	// Normally, UpdateProps would be named updateProps if all the property updaters were in the same
	// pacakge. However, in this case, the calculated field formula is updated in the CalcField package
	// so the function name needs to start with an upper case, so a FieldPropUpdater defined
	// in the CalcField package can be used.
	updateProps(appEngContext appengine.Context, filterForUpdate *RecordFilter) error
}

func updateFilterProps(appEngContext appengine.Context, propUpdater FilterPropUpdater) (*RecordFilterRef, error) {

	filterForUpdate, getErr := getFilter(appEngContext, propUpdater.GetFilterID())
	if getErr != nil {
		return nil, getErr
	}

	// Do the actual property update through the FilterPropUpdater interface
	if propUpdateErr := propUpdater.updateProps(appEngContext, filterForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("UpdateFilterProps: Unable to update existing filter properties: %v", propUpdateErr)
	}

	updatedFilterRef, updateErr := updateExistingFilter(appEngContext, propUpdater.GetFilterID(), filterForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("UpdateFilterProps: error updating filter: %v", updateErr)
	}

	return updatedFilterRef, nil

}

type FilterRenameParams struct {
	FilterIDHeader
	Name string `json:"name"`
}

func (updateParams FilterRenameParams) updateProps(appEngContext appengine.Context, filterForUpdate *RecordFilter) error {

	parentTableID, parentErr := datastoreWrapper.GetParentID(updateParams.GetFilterID(), table.TableEntityKind)
	if parentErr != nil {
		return fmt.Errorf("Update filter name: can't get parent table ID: error = %v", parentErr)
	}

	sanitizedName, sanitizeErr := generic.SanitizeName(updateParams.Name)
	if sanitizeErr != nil {
		return fmt.Errorf("Update filter name: sanitize name: %v", sanitizeErr)
	}

	if validateErr := validateUnusedFilterName(appEngContext, parentTableID, sanitizedName); validateErr != nil {
		return fmt.Errorf("Update filter name: validate unused name: error = %v", validateErr)
	}

	filterForUpdate.Name = sanitizedName

	return nil
}
