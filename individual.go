// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package withluminary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/withluminary/go-sdk/internal/apijson"
	"github.com/withluminary/go-sdk/internal/apiquery"
	"github.com/withluminary/go-sdk/internal/requestconfig"
	"github.com/withluminary/go-sdk/option"
	"github.com/withluminary/go-sdk/packages/pagination"
	"github.com/withluminary/go-sdk/packages/param"
	"github.com/withluminary/go-sdk/packages/respjson"
)

// IndividualService contains methods and other services that help with interacting
// with the luminary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndividualService] method instead.
type IndividualService struct {
	Options []option.RequestOption
}

// NewIndividualService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIndividualService(opts ...option.RequestOption) (r IndividualService) {
	r = IndividualService{}
	r.Options = opts
	return
}

// Create a new client profile/individual with the provided data
func (r *IndividualService) New(ctx context.Context, body IndividualNewParams, opts ...option.RequestOption) (res *Individual, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "individuals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve detailed information about a specific client profile
func (r *IndividualService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Individual, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("individuals/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing client profile with new data
func (r *IndividualService) Update(ctx context.Context, id string, body IndividualUpdateParams, opts ...option.RequestOption) (res *Individual, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("individuals/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of client profiles/individuals using cursor-based
// pagination
func (r *IndividualService) List(ctx context.Context, query IndividualListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Individual], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "individuals"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a paginated list of client profiles/individuals using cursor-based
// pagination
func (r *IndividualService) ListAutoPaging(ctx context.Context, query IndividualListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Individual] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Soft delete a client profile (marks as deleted but preserves data)
func (r *IndividualService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("individuals/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type Individual struct {
	// Unique identifier with client*profile* prefix
	ID string `json:"id" api:"required"`
	// Timestamp when the individual was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// First name of the individual
	FirstName string `json:"first_name" api:"required"`
	// Household ID this individual belongs to
	HouseholdID string `json:"household_id" api:"required"`
	// Whether this client profile should be an eligible beneficiary for entities and
	// gifts
	IsBeneficiary bool `json:"is_beneficiary" api:"required"`
	// Whether this client profile is deceased
	IsDeceased bool `json:"is_deceased" api:"required"`
	// Whether this client profile should be an eligible grantor/owner/other principal
	// for entities
	IsGrantor bool `json:"is_grantor" api:"required"`
	// Whether this is one of the (at most) two primary clients on this household
	IsPrimary bool `json:"is_primary" api:"required"`
	// Whether this client profile should be an eligible trustee for entities
	IsTrustee bool `json:"is_trustee" api:"required"`
	// Last name of the individual
	LastName string `json:"last_name" api:"required"`
	// Timestamp when the individual was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Street address line 1 (from address edge)
	AddressLine1 string `json:"address_line1" api:"nullable"`
	// Street address line 2 (from address edge)
	AddressLine2 string `json:"address_line2" api:"nullable"`
	// City (from address edge)
	City string `json:"city" api:"nullable"`
	// Country (from address edge)
	Country string `json:"country" api:"nullable"`
	// Date of birth (encrypted field)
	DateOfBirth time.Time `json:"date_of_birth" api:"nullable" format:"date"`
	// Date of death if applicable (encrypted field)
	DateOfDeath time.Time `json:"date_of_death" api:"nullable" format:"date"`
	// Timestamp when the individual was soft deleted
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// Email address
	Email string `json:"email" api:"nullable" format:"email"`
	// Middle name of the individual
	MiddleName string `json:"middle_name" api:"nullable"`
	// Notes about the client profile
	Notes string `json:"notes" api:"nullable"`
	// ZIP or postal code (from address edge)
	PostalCode string `json:"postal_code" api:"nullable"`
	// State or province (from address edge)
	State string `json:"state" api:"nullable"`
	// Name suffix (Jr., Sr., III, etc.)
	Suffix string `json:"suffix" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		FirstName     respjson.Field
		HouseholdID   respjson.Field
		IsBeneficiary respjson.Field
		IsDeceased    respjson.Field
		IsGrantor     respjson.Field
		IsPrimary     respjson.Field
		IsTrustee     respjson.Field
		LastName      respjson.Field
		UpdatedAt     respjson.Field
		AddressLine1  respjson.Field
		AddressLine2  respjson.Field
		City          respjson.Field
		Country       respjson.Field
		DateOfBirth   respjson.Field
		DateOfDeath   respjson.Field
		DeletedAt     respjson.Field
		Email         respjson.Field
		MiddleName    respjson.Field
		Notes         respjson.Field
		PostalCode    respjson.Field
		State         respjson.Field
		Suffix        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Individual) RawJSON() string { return r.JSON.raw }
func (r *Individual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualNewParams struct {
	// First name of the individual
	FirstName string `json:"first_name" api:"required"`
	// Household ID this individual belongs to
	HouseholdID string `json:"household_id" api:"required"`
	// Last name of the individual
	LastName string `json:"last_name" api:"required"`
	// Street address line 1
	AddressLine1 param.Opt[string] `json:"address_line1,omitzero"`
	// Street address line 2
	AddressLine2 param.Opt[string] `json:"address_line2,omitzero"`
	// City
	City param.Opt[string] `json:"city,omitzero"`
	// Country
	Country param.Opt[string] `json:"country,omitzero"`
	// Date of birth
	DateOfBirth param.Opt[time.Time] `json:"date_of_birth,omitzero" format:"date"`
	// Email address
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Middle name of the individual
	MiddleName param.Opt[string] `json:"middle_name,omitzero"`
	// Notes about the client profile
	Notes param.Opt[string] `json:"notes,omitzero"`
	// ZIP or postal code
	PostalCode param.Opt[string] `json:"postal_code,omitzero"`
	// State or province code (2 letter code)
	State param.Opt[string] `json:"state,omitzero"`
	// Name suffix
	Suffix param.Opt[string] `json:"suffix,omitzero"`
	// Whether this client profile should be an eligible beneficiary for entities and
	// gifts
	IsBeneficiary param.Opt[bool] `json:"is_beneficiary,omitzero"`
	// Whether the individual is deceased
	IsDeceased param.Opt[bool] `json:"is_deceased,omitzero"`
	// Whether this client profile should be an eligible grantor/owner/other principal
	// for entities
	IsGrantor param.Opt[bool] `json:"is_grantor,omitzero"`
	// Whether this is a primary client of the household (at most 2 per household)
	IsPrimary param.Opt[bool] `json:"is_primary,omitzero"`
	// Whether this client profile should be an eligible trustee for entities
	IsTrustee param.Opt[bool] `json:"is_trustee,omitzero"`
	paramObj
}

func (r IndividualNewParams) MarshalJSON() (data []byte, err error) {
	type shadow IndividualNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IndividualNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualUpdateParams struct {
	// Street address line 1
	AddressLine1 param.Opt[string] `json:"address_line1,omitzero"`
	// Street address line 2
	AddressLine2 param.Opt[string] `json:"address_line2,omitzero"`
	// City
	City param.Opt[string] `json:"city,omitzero"`
	// Country
	Country param.Opt[string] `json:"country,omitzero"`
	// Date of birth
	DateOfBirth param.Opt[time.Time] `json:"date_of_birth,omitzero" format:"date"`
	// Date of death if applicable
	DateOfDeath param.Opt[time.Time] `json:"date_of_death,omitzero" format:"date"`
	// Email address
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Middle name of the individual
	MiddleName param.Opt[string] `json:"middle_name,omitzero"`
	// Notes about the client profile
	Notes param.Opt[string] `json:"notes,omitzero"`
	// ZIP or postal code
	PostalCode param.Opt[string] `json:"postal_code,omitzero"`
	// State or province code (2 letter code)
	State param.Opt[string] `json:"state,omitzero"`
	// Name suffix
	Suffix param.Opt[string] `json:"suffix,omitzero"`
	// First name of the individual
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	// Whether this client profile should be an eligible beneficiary for entities and
	// gifts
	IsBeneficiary param.Opt[bool] `json:"is_beneficiary,omitzero"`
	// Whether the individual is deceased
	IsDeceased param.Opt[bool] `json:"is_deceased,omitzero"`
	// Whether this client profile should be an eligible grantor/owner/other principal
	// for entities
	IsGrantor param.Opt[bool] `json:"is_grantor,omitzero"`
	// Whether this is a primary client of the household (at most 2 per household)
	IsPrimary param.Opt[bool] `json:"is_primary,omitzero"`
	// Whether this client profile should be an eligible trustee for entities
	IsTrustee param.Opt[bool] `json:"is_trustee,omitzero"`
	// Last name of the individual
	LastName param.Opt[string] `json:"last_name,omitzero"`
	paramObj
}

func (r IndividualUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow IndividualUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IndividualUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualListParams struct {
	// Cursor for forward pagination. Returns items after this cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returns items before this cursor.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter individuals by household ID
	HouseholdID param.Opt[string] `query:"household_id,omitzero" json:"-"`
	// Filter by primary client status
	IsPrimary param.Opt[bool] `query:"is_primary,omitzero" json:"-"`
	// Maximum number of items to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IndividualListParams]'s query parameters as `url.Values`.
func (r IndividualListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
