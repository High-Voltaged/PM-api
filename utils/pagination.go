package utils

import (
	"api/consts"
	"net/url"
	"strconv"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func GetPaginationFromQuery(query url.Values) Pagination {
	pagination := Pagination{
		Limit: consts.DEFAULT_LIMIT,
		Page:  consts.DEFAULT_PAGE,
		Sort:  consts.DEFAULT_SORT,
	}

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			pagination.Limit, _ = strconv.Atoi(queryValue)
		case "page":
			pagination.Page, _ = strconv.Atoi(queryValue)
		case "sort":
			pagination.Sort = queryValue
		}
	}

	return pagination
}

func CalcPageOffset(opts Pagination) int {
	return (opts.Page - 1) * opts.Limit
}
