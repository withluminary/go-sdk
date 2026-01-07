// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package withluminary_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/withluminary-go"
	"github.com/stainless-sdks/withluminary-go/internal/testutil"
	"github.com/stainless-sdks/withluminary-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	households, err := client.Households.List(context.TODO(), withluminary.HouseholdListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", households.Data)
}
