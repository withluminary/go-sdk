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
		AddressLine1:  withluminary.String("address_line1"),
		AddressLine2:  withluminary.String("address_line2"),
		City:          withluminary.String("city"),
		Country:       withluminary.String("country"),
		DateOfBirth:   withluminary.Time(time.Now()),
		Email:         withluminary.String("dev@stainless.com"),
		IsBeneficiary: withluminary.Bool(true),
		IsDeceased:    withluminary.Bool(true),
		IsGrantor:     withluminary.Bool(true),
		IsPrimary:     withluminary.Bool(true),
		IsTrustee:     withluminary.Bool(true),
		MiddleName:    withluminary.String("middle_name"),
		Notes:         withluminary.String("notes"),
		PostalCode:    withluminary.String("postal_code"),
		State:         withluminary.String("state"),
		Suffix:        withluminary.String("suffix"),
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
			AddressLine1:  withluminary.String("address_line1"),
			AddressLine2:  withluminary.String("address_line2"),
			City:          withluminary.String("city"),
			Country:       withluminary.String("country"),
			DateOfBirth:   withluminary.Time(time.Now()),
			DateOfDeath:   withluminary.Time(time.Now()),
			Email:         withluminary.String("dev@stainless.com"),
			FirstName:     withluminary.String("first_name"),
			IsBeneficiary: withluminary.Bool(true),
			IsDeceased:    withluminary.Bool(true),
			IsGrantor:     withluminary.Bool(true),
			IsPrimary:     withluminary.Bool(true),
			IsTrustee:     withluminary.Bool(true),
			LastName:      withluminary.String("last_name"),
			MiddleName:    withluminary.String("middle_name"),
			Notes:         withluminary.String("notes"),
			PostalCode:    withluminary.String("postal_code"),
			State:         withluminary.String("state"),
			Suffix:        withluminary.String("suffix"),
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
		HouseholdID: withluminary.String("household_id"),
		IsPrimary:   withluminary.Bool(true),
		Limit:       withluminary.Int(1),
		Offset:      withluminary.Int(0),
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
