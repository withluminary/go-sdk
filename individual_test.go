// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package withluminary_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/withluminary/go-sdk"
	"github.com/withluminary/go-sdk/internal/testutil"
	"github.com/withluminary/go-sdk/option"
)

func TestIndividualNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Individuals.New(context.TODO(), withluminary.IndividualNewParams{
		FirstName:     "John",
		HouseholdID:   "household_01ARZ3NDEKTSV4RRFFQ69G5FAV",
		LastName:      "Smith",
		AddressLine1:  withluminary.String("x"),
		AddressLine2:  withluminary.String("x"),
		City:          withluminary.String("x"),
		Country:       withluminary.String("x"),
		DateOfBirth:   withluminary.Time(time.Now()),
		Email:         withluminary.String("dev@stainless.com"),
		IsBeneficiary: withluminary.Bool(true),
		IsDeceased:    withluminary.Bool(true),
		IsGrantor:     withluminary.Bool(true),
		IsPrimary:     withluminary.Bool(true),
		IsTrustee:     withluminary.Bool(true),
		MiddleName:    withluminary.String("x"),
		Notes:         withluminary.String("notes"),
		PostalCode:    withluminary.String("x"),
		State:         withluminary.String("xx"),
		Suffix:        withluminary.String("x"),
	})
	if err != nil {
		var apierr *withluminary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIndividualGet(t *testing.T) {
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
	_, err := client.Individuals.Get(context.TODO(), "id")
	if err != nil {
		var apierr *withluminary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIndividualUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Individuals.Update(
		context.TODO(),
		"id",
		withluminary.IndividualUpdateParams{
			AddressLine1:  withluminary.String("x"),
			AddressLine2:  withluminary.String("x"),
			City:          withluminary.String("x"),
			Country:       withluminary.String("x"),
			DateOfBirth:   withluminary.Time(time.Now()),
			DateOfDeath:   withluminary.Time(time.Now()),
			Email:         withluminary.String("dev@stainless.com"),
			FirstName:     withluminary.String("x"),
			IsBeneficiary: withluminary.Bool(true),
			IsDeceased:    withluminary.Bool(true),
			IsGrantor:     withluminary.Bool(true),
			IsPrimary:     withluminary.Bool(true),
			IsTrustee:     withluminary.Bool(true),
			LastName:      withluminary.String("x"),
			MiddleName:    withluminary.String("x"),
			Notes:         withluminary.String("notes"),
			PostalCode:    withluminary.String("x"),
			State:         withluminary.String("xx"),
			Suffix:        withluminary.String("x"),
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

func TestIndividualListWithOptionalParams(t *testing.T) {
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
	_, err := client.Individuals.List(context.TODO(), withluminary.IndividualListParams{
		After:       withluminary.String("eyJpZCI6ImhvdXNlaG9sZF8wMUFSWjNOREVLVFNWNFJSRkZRNjlHNUZBViJ9"),
		Before:      withluminary.String("eyJpZCI6ImhvdXNlaG9sZF8wMUFSWjNOREVLVFNWNFJSRkZRNjlHNUZBViJ9"),
		HouseholdID: withluminary.String("household_id"),
		IsPrimary:   withluminary.Bool(true),
		Limit:       withluminary.Int(1),
	})
	if err != nil {
		var apierr *withluminary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIndividualDelete(t *testing.T) {
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
	err := client.Individuals.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *withluminary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
