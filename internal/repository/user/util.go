package repository_user

import (
	"api-example/internal/domain"
	"fmt"
	"strings"
)

func getQueryFilter(filter domain.User) string {
	queryFilter := []string{}

	if filter.Email != "" {
		queryFilter = append(queryFilter, fmt.Sprintf("email = '%s'", filter.Email))
	}

	if filter.FullName != "" {
		queryFilter = append(queryFilter, "fullname LIKE '%"+filter.FullName+"%'")
	}

	if filter.Status != "" {
		queryFilter = append(queryFilter, fmt.Sprintf("status = '%s'", filter.Status))
	}

	if len(queryFilter) == 0 {
		return "1=1"
	}

	return strings.Join(queryFilter, " AND ")
}
