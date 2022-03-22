package main

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/vkuznet/dbs2go/dbs"
	"github.com/vkuznet/dbs2go/web"
)

// datatiers endpoint tests
func getDatatiersTestTable(t *testing.T) EndpointTestCase {
	tiersReq := dbs.DataTiers{
		DATA_TIER_NAME: TestData.Tier,
		CREATE_BY:      "tester",
	}
	tiersResp := dbs.DataTiers{
		DATA_TIER_ID:   1,
		DATA_TIER_NAME: TestData.Tier,
		CREATE_BY:      "tester",
		CREATION_DATE:  0,
	}
	badReq := BadRequest{
		BAD_FIELD: "BAD",
	}
	return EndpointTestCase{
		description:     "Test datatiers",
		defaultHandler:  web.DatatiersHandler,
		defaultEndpoint: "/dbs/datatiers",
		testCases: []testCase{
			{
				description: "Test GET with no data",
				method:      "GET",
				serverType:  "DBSReader",
				input:       nil,
				params:      nil,
				output:      []Response{},
				respCode:    http.StatusOK,
			},
			{
				description: "Test bad POST",
				method:      "POST",
				serverType:  "DBSWriter",
				input:       badReq,
				params:      nil,
				respCode:    http.StatusBadRequest,
			},
			{
				description: "Test POST",
				method:      "POST",
				serverType:  "DBSWriter",
				input:       tiersReq,
				output: []Response{
					tiersResp,
				},
				params:   nil,
				respCode: http.StatusOK,
			},
			{
				description: "Test GET after POST",
				method:      "GET",
				serverType:  "DBSReader",
				output: []Response{
					tiersResp,
				},
				params:   nil,
				respCode: http.StatusOK,
			},
			{
				description: "Test GET with parameters",
				method:      "GET",
				serverType:  "DBSReader",
				params: url.Values{
					"data_tier_name": []string{
						TestData.Tier,
					},
				},
				output: []Response{
					tiersResp,
				},
				respCode: http.StatusOK,
			},
			{
				description: "Test GET with regex parameter",
				method:      "GET",
				serverType:  "DBSReader",
				params: url.Values{
					"data_tier_name": []string{"G*"},
				},
				output: []Response{
					tiersResp,
				},
				respCode: http.StatusOK,
			},
			{
				description: "Test GET with non-existing parameter value",
				method:      "GET",
				serverType:  "DBSReader",
				output:      []Response{},
				params: url.Values{
					"data_tier_name": []string{"A*"},
				},
				respCode: http.StatusOK,
			},
		},
	}

}
