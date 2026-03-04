// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package withluminary_test

import (
	"context"
	"os"
	"testing"

	"github.com/withluminary/go-sdk"
	"github.com/withluminary/go-sdk/internal/testutil"
	"github.com/withluminary/go-sdk/option"
)

func TestAutoPagination(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := withluminary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	iter := client.Households.ListAutoPaging(context.TODO(), withluminary.HouseholdListParams{})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		household := iter.Current()
		t.Logf("%+v\n", household.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
