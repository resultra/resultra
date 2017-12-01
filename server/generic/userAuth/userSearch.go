package userAuth

import (
	"database/sql"
	"fmt"
	"strings"
)

type SearchUserMatch struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
}

type SearchUsersResults struct {
	MatchedUserInfo []SearchUserMatch `json:"matchedUserInfo"`
}

func searchUsers(trackerDBHandle *sql.DB, searchTerm string) (*SearchUsersResults, error) {

	sqlSearchTerm := strings.ToUpper(`%` + searchTerm + `%`)

	rows, queryErr := trackerDBHandle.Query(
		`SELECT user_id,user_name 
		FROM users 
		WHERE UPPER(user_name) LIKE $1`, sqlSearchTerm)
	if queryErr != nil {
		return nil, fmt.Errorf("searchUserInfo: Failure querying database: %v", queryErr)
	}
	defer rows.Close()

	matchingUsers := []SearchUserMatch{}
	for rows.Next() {
		var currMatch SearchUserMatch
		if scanErr := rows.Scan(&currMatch.UserID,
			&currMatch.UserName); scanErr != nil {
			return nil, fmt.Errorf("searchUserInfo: Failure querying database: %v", scanErr)

		}
		matchingUsers = append(matchingUsers, currMatch)
	}

	matchResults := SearchUsersResults{
		MatchedUserInfo: matchingUsers}

	return &matchResults, nil

}
