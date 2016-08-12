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

type SetFormRolePrivsParams struct {
	FormID string `json:"formID"`
	RoleID string `json:"roleID"`
	Privs  string `json:"privs"`
}

func setFormRolePrivs(params SetFormRolePrivsParams) error {

	if privsErr := verifyFormRolePrivs(params.Privs); privsErr != nil {
		return fmt.Errorf("setFormRolePrivs: error = %v", privsErr)
	}

	if _, deleteErr := databaseWrapper.DBHandle().Exec(
		`DELETE FROM form_role_privs where role_id=$1 and form_id=$2`,
		params.RoleID, params.FormID); deleteErr != nil {
		return fmt.Errorf("setFormRolePrivs: Can't delete old privs: error = %v", deleteErr)
	}

	if _, insertErr := databaseWrapper.DBHandle().Exec(
		`INSERT INTO form_role_privs (role_id,form_id,privs) VALUES ($1,$2,$3)`,
		params.RoleID, params.FormID, params.Privs); insertErr != nil {
		return fmt.Errorf("setFormRolePrivs: Can't set form privileges: error = %v", insertErr)
	}

	return nil

}

type FormRolePriv struct {
	RoleID   string `json:"roleID"`
	RoleName string `json:"roleName"`
	Privs    string `json:"privs"`
}

func getFormRolePrivs(formID string) ([]FormRolePriv, error) {

	rows, queryErr := databaseWrapper.DBHandle().Query(
		`SELECT database_roles.role_id,database_roles.name,form_role_privs.privs
			FROM form_role_privs,database_roles
			WHERE form_role_privs.form_id=$1
				AND database_roles.role_id=form_role_privs.role_id`, formID)
	if queryErr != nil {
		return nil, fmt.Errorf("getFormRolePrivs: Failure querying database: %v", queryErr)
	}

	formRolePrivs := []FormRolePriv{}
	for rows.Next() {

		currPrivInfo := FormRolePriv{}

		if scanErr := rows.Scan(&currPrivInfo.RoleID, &currPrivInfo.RoleName, &currPrivInfo.Privs); scanErr != nil {
			return nil, fmt.Errorf("getFormRolePrivs: Failure querying database: %v", scanErr)
		}

		formRolePrivs = append(formRolePrivs, currPrivInfo)

	}

	return formRolePrivs, nil

}
