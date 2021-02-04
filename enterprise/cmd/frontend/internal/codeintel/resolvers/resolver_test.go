package resolvers

import (
	"testing"
)

func TestResolverTODO(t *testing.T) {
	// TODO
}

// func TestQueryResolver(t *testing.T) {
// 	mockDBStore := NewMockDBStore()
// 	mockLSIFStore := NewMockLSIFStore()
// 	gitServerClient := enqueuer.NewMockGitserverClient()

// 	resolver := NewResolver(mockDBStore, mockLSIFStore, gitServerClient, nil, nil, &observation.TestContext)
// 	queryResolver, err := resolver.QueryResolver(context.Background(), &gql.GitBlobLSIFDataArgs{
// 		Repo:      &types.Repo{ID: 50},
// 		Commit:    api.CommitID("deadbeef"),
// 		Path:      "/foo/bar.go",
// 		ExactPath: true,
// 		ToolName:  "lsif-go",
// 	})
// 	if err != nil {
// 		t.Fatalf("unexpected error: %s", err)
// 	}

// 	if queryResolver != nil {
// 		t.Errorf("expected nil-valued resolver")
// 	}
// }

// func TestFallbackIndexConfiguration(t *testing.T) {
// 	mockDBStore := NewMockDBStore()
// 	mockEnqueuerDBStore := enqueuer.NewMockDBStore()
// 	mockLSIFStore := NewMockLSIFStore()
// 	mockCodeIntelAPI := NewMockCodeIntelAPI() // returns no dumps
// 	gitServerClient := enqueuer.NewMockGitserverClient()
// 	indexEnqueuer := enqueuer.NewIndexEnqueuer(mockEnqueuerDBStore, gitServerClient, &observation.TestContext)

// 	resolver := NewResolver(mockDBStore, mockLSIFStore, mockCodeIntelAPI, indexEnqueuer, nil, &observation.TestContext)

// 	mockDBStore.GetIndexConfigurationByRepositoryIDFunc.SetDefaultReturn(dbstore.IndexConfiguration{}, false, nil)
// 	gitServerClient.HeadFunc.SetDefaultReturn("deadbeef", nil)
// 	gitServerClient.ListFilesFunc.SetDefaultReturn([]string{"go.mod"}, nil)

// 	json, err := resolver.IndexConfiguration(context.Background(), 0)

// 	if err != nil {
// 		t.Fatalf("unexpected error: %s", err)
// 	}

// 	diff := cmp.Diff(string(json), `{
// 	"shared_steps": [],
// 	"index_jobs": [
// 		{
// 			"steps": [
// 				{
// 					"root": "",
// 					"image": "sourcegraph/lsif-go:latest",
// 					"commands": [
// 						"go mod download"
// 					]
// 				}
// 			],
// 			"local_steps": [],
// 			"root": "",
// 			"indexer": "sourcegraph/lsif-go:latest",
// 			"indexer_args": [
// 				"lsif-go",
// 				"--no-animation"
// 			],
// 			"outfile": ""
// 		}
// 	]
// }`)

// 	if diff != "" {
// 		t.Fatalf("Unexpected fallback index configuration:\n%s\n", diff)
// 	}
// }
