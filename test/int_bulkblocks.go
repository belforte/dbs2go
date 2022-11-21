package main

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/dmwm/dbs2go/dbs"
	"github.com/dmwm/dbs2go/web"
)

// this file contains logic for bulkblocks API
// both sequential and concurrent bulkblocks are tested
// HTTP request body data is defined in test/data/integration/bulkblocks_data.json. This is generated using generateBulkBlocksData in test/integration_cases.go
// sequential bulkblocks data is under the seq_parent_bulk and seq_child_bulk fields in test/data/integration/bulkblocks_data.json
// concurrent bulkblocks data is under the con_parent_bulk and con_child_bulk fields in test/data/integration/bulkblocks_data.json
// bulkblocks_data.json is loaded into BulkBlocksData struct defined in test/integration_cases.go
// the HTTP request body is defined by dbs.BulkBlocks struct defined in dbs/bulkblocks.go
// sequential bulkblocks data is loaded into SequentialParentData and SequentialChildData in BulkBlocksData struct
// concurrent bulkblocks data is loaded into ConcurrentParentData and ConcurrentChildData in BulkBlocksData struct
// the HTTP handlers and endpoints are defined in the EndpointTestCase struct defined in test/integration_cases.go

// bulkblocks test table
func getBulkBlocksTestTable(t *testing.T) EndpointTestCase {
	return EndpointTestCase{
		description:     "Test bulkblocks",
		defaultHandler:  web.BulkBlocksHandler,
		defaultEndpoint: "/dbs/bulkblocks",
		testCases: []testCase{
			{
				description:          "Test POST sequential parent bulkblocks",
				serverType:           "DBSWriter",
				concurrentBulkBlocks: false,
				method:               "POST",
				input:                BulkBlocksData.SequentialParentData,
				params: url.Values{
					"block_name": []string{TestData.ParentStepchainBlock + "2"},
				},
				output:   []Response{},
				handler:  web.FilesHandler,
				respCode: http.StatusOK,
			},
			{
				description:          "Test POST sequential child bulkblocks",
				serverType:           "DBSWriter",
				concurrentBulkBlocks: false,
				method:               "POST",
				input:                BulkBlocksData.SequentialChildData,
				params: url.Values{
					"block_name": []string{TestData.StepchainBlock + "2"},
				},
				output:   []Response{},
				handler:  web.FilesHandler,
				respCode: http.StatusOK,
			},
			{
				description:          "Test POST concurrent parent bulkblocks",
				serverType:           "DBSWriter",
				concurrentBulkBlocks: true,
				method:               "POST",
				input:                BulkBlocksData.ConcurrentParentData,
				output:               []Response{},
				params: url.Values{
					"block_name": []string{TestData.ParentStepchainBlock},
				},
				handler:  web.FilesHandler,
				respCode: http.StatusOK,
			},
			{
				description:          "Test POST concurrent child bulkblocks",
				serverType:           "DBSWriter",
				concurrentBulkBlocks: true,
				method:               "POST",
				input:                BulkBlocksData.ConcurrentChildData,
				output:               []Response{},
				params: url.Values{
					"block_name": []string{TestData.StepchainBlock},
				},
				handler:  web.FilesHandler,
				respCode: http.StatusOK,
			},
		},
	}
}

// bulkblocks test table
func getBulkBlocksLargeFileLumiInsertTestTable(t *testing.T) EndpointTestCase {
	return EndpointTestCase{
		description:     "Test concurrent bulkblocks when fileLumiChunkSize less than number fileLumis inserted",
		defaultHandler:  web.BulkBlocksHandler,
		defaultEndpoint: "/dbs/bulkblocks",
		testCases: []testCase{
			{
				description:          "Test POST with fileLumiChunk size 20",
				serverType:           "DBSWriter",
				method:               "POST",
				fileLumiChunkSize:    20,
				concurrentBulkBlocks: true,
				input:                LargeBulkBlocksData,
				params: url.Values{
					"block_name": []string{TestData.StepchainBlock + "2"},
				},
				output:   []Response{},
				handler:  web.FilesHandler,
				respCode: http.StatusOK,
			},
		},
	}
}

// test that we will get DBSError when use the same block with bulkblocks API
func bulkblocksTheSameBlockInsertTestTable(t *testing.T) EndpointTestCase {
	// there are multiple blocks to insert, but everything is started from parent blocks
	bName := TestData.ParentStepchainBlock
	reason := fmt.Sprintf("Block %s already exists", bName)
	msg := "Data already exist in DBS"
	dbsError := dbs.DBSError{
		Function: "dbs.bulkblocks.checkBlockExist",
		Code:     dbs.BlockAlreadyExists,
		Reason:   reason,
		Message:  msg,
	}
	hrec := createHTTPError("POST", "/dbs/bulkblocks")
	errorResp := createServerErrorResponse(hrec, &dbsError)
	return EndpointTestCase{
		description:     "Test concurrent bulkblocks with the same block name twice",
		defaultHandler:  web.BulkBlocksHandler,
		defaultEndpoint: "/dbs/bulkblocks",
		testCases: []testCase{
			{
				description:          "Test POST with the same block",
				serverType:           "DBSWriter",
				method:               "POST",
				concurrentBulkBlocks: true,
				input:                BulkBlocksData.ConcurrentParentData,
				output: []Response{
					errorResp,
				},
				handler:  web.BulkBlocksHandler,
				respCode: http.StatusBadRequest,
			},
		},
	}
}
