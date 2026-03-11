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

// HouseholdService contains methods and other services that help with interacting
// with the luminary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHouseholdService] method instead.
type HouseholdService struct {
	Options []option.RequestOption
}

// NewHouseholdService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHouseholdService(opts ...option.RequestOption) (r HouseholdService) {
	r = HouseholdService{}
	r.Options = opts
	return
}

// Create a new household with the provided data
func (r *HouseholdService) New(ctx context.Context, body HouseholdNewParams, opts ...option.RequestOption) (res *Household, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "households"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve detailed information about a specific household
func (r *HouseholdService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Household, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("households/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing household with new data
func (r *HouseholdService) Update(ctx context.Context, id string, body HouseholdUpdateParams, opts ...option.RequestOption) (res *Household, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("households/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of households using cursor-based pagination
func (r *HouseholdService) List(ctx context.Context, query HouseholdListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Household], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "households"
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

// Retrieve a paginated list of households using cursor-based pagination
func (r *HouseholdService) ListAutoPaging(ctx context.Context, query HouseholdListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Household] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Soft delete a household (marks as deleted but preserves data)
func (r *HouseholdService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("households/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a paginated list of documents belonging to a specific household
func (r *HouseholdService) ListDocuments(ctx context.Context, id string, query HouseholdListDocumentsParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Document], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("households/%s/documents", id)
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

// Retrieve a paginated list of documents belonging to a specific household
func (r *HouseholdService) ListDocumentsAutoPaging(ctx context.Context, id string, query HouseholdListDocumentsParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Document] {
	return pagination.NewCursorPaginationAutoPager(r.ListDocuments(ctx, id, query, opts...))
}

// Retrieve a paginated list of entities belonging to a specific household
func (r *HouseholdService) ListEntities(ctx context.Context, id string, query HouseholdListEntitiesParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Entity], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("households/%s/entities", id)
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

// Retrieve a paginated list of entities belonging to a specific household
func (r *HouseholdService) ListEntitiesAutoPaging(ctx context.Context, id string, query HouseholdListEntitiesParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Entity] {
	return pagination.NewCursorPaginationAutoPager(r.ListEntities(ctx, id, query, opts...))
}

// Retrieve a paginated list of client profiles/individuals belonging to a specific
// household
func (r *HouseholdService) ListIndividuals(ctx context.Context, id string, query HouseholdListIndividualsParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Individual], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("households/%s/individuals", id)
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

// Retrieve a paginated list of client profiles/individuals belonging to a specific
// household
func (r *HouseholdService) ListIndividualsAutoPaging(ctx context.Context, id string, query HouseholdListIndividualsParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Individual] {
	return pagination.NewCursorPaginationAutoPager(r.ListIndividuals(ctx, id, query, opts...))
}

type Household struct {
	// Unique identifier with household\_ prefix
	ID string `json:"id" api:"required"`
	// Timestamp when the household was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// User ID of the primary relationship owner
	PrimaryRelationshipOwnerID string `json:"primary_relationship_owner_id" api:"required"`
	// Timestamp when the household was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Display name for the household
	Name string `json:"name"`
	// Notes about the household
	Notes string `json:"notes" api:"nullable"`
	// Primary client profiles for this household (at most 2)
	PrimaryIndividuals []Individual `json:"primary_individuals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		CreatedAt                  respjson.Field
		PrimaryRelationshipOwnerID respjson.Field
		UpdatedAt                  respjson.Field
		Name                       respjson.Field
		Notes                      respjson.Field
		PrimaryIndividuals         respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Household) RawJSON() string { return r.JSON.raw }
func (r *Household) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualList struct {
	Data     []Individual `json:"data" api:"required"`
	PageInfo PageInfo     `json:"page_info" api:"required"`
	// Total number of items matching the query (across all pages)
	TotalCount int64 `json:"total_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		PageInfo    respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndividualList) RawJSON() string { return r.JSON.raw }
func (r *IndividualList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HouseholdNewParams struct {
	// User ID of the primary relationship owner
	PrimaryRelationshipOwnerID string `json:"primary_relationship_owner_id" api:"required"`
	// Optional notes about the household
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Primary client profiles to create for this household (at most 2)
	PrimaryIndividuals []HouseholdNewParamsPrimaryIndividual `json:"primary_individuals,omitzero"`
	paramObj
}

func (r HouseholdNewParams) MarshalJSON() (data []byte, err error) {
	type shadow HouseholdNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HouseholdNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FirstName, LastName, State are required.
type HouseholdNewParamsPrimaryIndividual struct {
	// First name of the individual
	FirstName string `json:"first_name" api:"required"`
	// Last name of the individual
	LastName string `json:"last_name" api:"required"`
	// State or province code (2 letter code)
	State string `json:"state" api:"required"`
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
	// Name suffix
	Suffix param.Opt[string] `json:"suffix,omitzero"`
	// Whether this client profile should be an eligible beneficiary for entities and
	// gifts
	IsBeneficiary param.Opt[bool] `json:"is_beneficiary,omitzero"`
	// Whether the individual is deceased
	IsDeceased param.Opt[bool] `json:"is_deceased,omitzero"`
	// Whether this client profile should be an eligible trustee for entities
	IsTrustee param.Opt[bool] `json:"is_trustee,omitzero"`
	paramObj
}

func (r HouseholdNewParamsPrimaryIndividual) MarshalJSON() (data []byte, err error) {
	type shadow HouseholdNewParamsPrimaryIndividual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HouseholdNewParamsPrimaryIndividual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HouseholdUpdateParams struct {
	// Notes about the household
	Notes param.Opt[string] `json:"notes,omitzero"`
	// User ID of the primary relationship owner
	PrimaryRelationshipOwnerID param.Opt[string] `json:"primary_relationship_owner_id,omitzero"`
	paramObj
}

func (r HouseholdUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow HouseholdUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HouseholdUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HouseholdListParams struct {
	// Cursor for forward pagination. Returns items after this cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returns items before this cursor.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HouseholdListParams]'s query parameters as `url.Values`.
func (r HouseholdListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HouseholdListDocumentsParams struct {
	// Cursor for forward pagination. Returns items after this cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returns items before this cursor.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by document type
	//
	// Any of "GRAT_DESIGN_SUMMARY", "GENERATED_PRESENTATION", "ASSET_VALUATION",
	// "SIGNED_TRUST_DOCUMENT", "TRUST_AMENDMENT", "TRANSFER_CONFIRMATION",
	// "EXISTING_REMAINDER_TRUST_DOCUMENT", "BALANCE_SHEET", "WILL", "WILL_CODICIL",
	// "POWER_OF_ATTORNEY", "ASSIGNMENT_OF_INTEREST",
	// "ASSIGNMENT_OF_TANGIBLE_PROPERTY", "LOAN_NOTE_AGREEMENT",
	// "ARTICLES_OF_INCORPORATION", "OPERATING_AGREEMENT", "PARTNERSHIP_AGREEMENT",
	// "ACCOUNT_DOCUMENTATION_STATEMENT", "TAX_ID_CONFIRMATION", "GIFT_TAX_RETURN",
	// "INCOME_TAX_RETURN", "TAX_RECEIPT", "TAX_FILING", "CORPORATE_BYLAWS",
	// "LLC_AGREEMENT", "LLC_AGREEMENT_AMENDMENT", "OPERATING_AGREEMENT_AMENDMENT",
	// "PARTNERSHIP_AGREEMENT_AMENDMENT", "SHAREHOLDERS_AGREEMENT",
	// "STATE_BUSINESS_FILING", "LOGGED_CONTRIBUTION", "LOGGED_DISTRIBUTION",
	// "INSURANCE_POLICY", "CRUMMEY_LETTER", "INSURANCE_PREMIUM_PAYMENT",
	// "BENEFICIAL_OWNERSHIP_INFORMATION_REPORT", "FINCEN_FILING", "HEALTHCARE_PROXY",
	// "LIVING_WILL", "DRIVERS_LICENSE", "PASSPORT", "DEED", "OTHER".
	Type DocumentType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HouseholdListDocumentsParams]'s query parameters as
// `url.Values`.
func (r HouseholdListDocumentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HouseholdListEntitiesParams struct {
	// Cursor for forward pagination. Returns items after this cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returns items before this cursor.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by entity kind/type
	//
	// Any of "REVOCABLE_TRUST", "IRREVOCABLE_TRUST", "SLAT_TRUST", "ILIT_TRUST",
	// "QPRT_TRUST", "GRAT_TRUST", "CRT_TRUST", "CLT_TRUST",
	// "INDIVIDUAL_PERSONAL_ACCOUNT", "JOINT_PERSONAL_ACCOUNT",
	// "CUSTODIAL_PERSONAL_ACCOUNT", "INSURANCE_PERSONAL_ACCOUNT",
	// "QUALIFIED_TUITION_PERSONAL_ACCOUNT", "RETIREMENT_PERSONAL_ACCOUNT",
	// "DONOR_ADVISED_FUND", "PRIVATE_FOUNDATION", "LLC_BUSINESS_ENTITY",
	// "LP_BUSINESS_ENTITY", "GP_BUSINESS_ENTITY",
	// "SOLE_PROPRIETORSHIP_BUSINESS_ENTITY", "SCORP_BUSINESS_ENTITY",
	// "CCORP_BUSINESS_ENTITY".
	Kind EntityKind `query:"kind,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HouseholdListEntitiesParams]'s query parameters as
// `url.Values`.
func (r HouseholdListEntitiesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HouseholdListIndividualsParams struct {
	// Cursor for forward pagination. Returns items after this cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returns items before this cursor.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter by primary client status
	IsPrimary param.Opt[bool] `query:"is_primary,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HouseholdListIndividualsParams]'s query parameters as
// `url.Values`.
func (r HouseholdListIndividualsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
