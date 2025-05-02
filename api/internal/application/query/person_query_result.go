package query

import "api/internal/application/common"

type PersonQueryResult struct {
	Result *common.PersonResult
}

type PersonQueryListResult struct {
	Result []*common.PersonResult
}
