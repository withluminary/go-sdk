// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package withluminary_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/withluminary-go"
	"github.com/stainless-sdks/withluminary-go/internal/testutil"
	"github.com/stainless-sdks/withluminary-go/option"
)

func TestEntityValuationNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Entities.Valuation.New(
		context.TODO(),
		"id",
		withluminary.EntityValuationNewParams{
			DirectlyHeldAssets: []withluminary.EntityValuationNewParamsDirectlyHeldAsset{{
				AssetClassID: "asset_class_01ARZ3NDEKTSV4RRFFQ69G5FAV",
				DisplayName:  "Apple Inc. Stock",
				Value:        50000,
				ExternalID:   withluminary.String("AAPL-12345"),
			}},
			EffectiveDate: time.Now(),
			Description:   withluminary.String("description"),
		},
	)
	if err != nil {
		var apierr *withluminary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityValuationGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Entities.Valuation.Get(context.TODO(), "id")
	if err != nil {
		var apierr *withluminary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
