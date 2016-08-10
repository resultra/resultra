package userRole

import (
	"fmt"
	"resultra/datasheet/server/generic/databaseWrapper"
)

const FormRolePrivsNone string = "none"
const FormRolePrivsView string = "view"
const FormRolePrivsEdit string = "edit"

type FormPriv struct {
	FormID string `json:"formID"`
	Privs  string `json:"privs"`
}

func verifyFormRolePrivs(privs string) error {
	if (privs == FormRolePrivsNone) ||
		(privs == FormRolePrivsView) ||
		(privs == FormRolePrivsEdit) {
		return nil
	} else {
		return fmt.Errorf("verifyFormRolePrivs: Invalid privileges: %v", privs)
	}
}

func setFormRolePrivs(roleID string, formID string, privs string) error {

	if privsErr := verifyFormRolePrivs(privs); privsErr != nil {
		return fmt.Errorf("setFormRolePrivs: error = %v", privsErr)
	}

	if _, deleteErr := databaseWrapper.DBHandle().Exec(
		`DELETE FROM form_role_privs where role_id=$1 and form_id=$2`,
		roleID, formID); deleteErr != nil {
		return fmt.Errorf("setFormRolePrivs: Can't delete old privs: error = %v", deleteErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO form_role_privs (role_id,form_id,privs) VALUES ($1,$2,$3)`,
		roleID, formID, privs); insertErr != nil {
		return fmt.Errorf("setFormRolePrivs: Can't set form privileges: error = %v", insertErr)
	}

	return nil

}
