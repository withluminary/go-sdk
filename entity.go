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

	"github.com/stainless-sdks/withluminary-go/internal/apijson"
	"github.com/stainless-sdks/withluminary-go/internal/apiquery"
	"github.com/stainless-sdks/withluminary-go/internal/requestconfig"
	"github.com/stainless-sdks/withluminary-go/option"
	"github.com/stainless-sdks/withluminary-go/packages/param"
	"github.com/stainless-sdks/withluminary-go/packages/respjson"
)

// EntityService contains methods and other services that help with interacting
// with the luminary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityService] method instead.
type EntityService struct {
	Options   []option.RequestOption
	Valuation EntityValuationService
}

// NewEntityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEntityService(opts ...option.RequestOption) (r EntityService) {
	r = EntityService{}
	r.Options = opts
	r.Valuation = NewEntityValuationService(opts...)
	return
}

// Retrieve detailed information about a specific entity
func (r *EntityService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a paginated list of entities (trusts, businesses, accounts, etc.) using
// cursor-based pagination
func (r *EntityService) List(ctx context.Context, query EntityListParams, opts ...option.RequestOption) (res *EntityList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "entities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an entity and all of it's related data
func (r *EntityService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Entity struct {
	// Unique identifier with entity\_ prefix
	ID string `json:"id,required"`
	// Timestamp when the entity was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Display name of the entity
	DisplayName string `json:"display_name,required"`
	// Household ID this entity belongs to
	HouseholdID string `json:"household_id,required"`
	// Type of entity - determines the specific subtype and applicable fields
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
	Kind EntityKind `json:"kind,required"`
	// Lifecycle stage of the entity
	//
	// Any of "PRE_CREATED", "AI_CREATING", "AI_CREATION_FAILED", "AI_NEEDS_REVIEW",
	// "DRAFT", "READY_FOR_PROPOSAL", "IMPLEMENTATION", "ACTIVE", "COMPLETED",
	// "ARCHIVED".
	Stage EntityStage `json:"stage,required"`
	// Timestamp when the entity was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DisplayName respjson.Field
		HouseholdID respjson.Field
		Kind        respjson.Field
		Stage       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Entity) RawJSON() string { return r.JSON.raw }
func (r *Entity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle stage of the entity
type EntityStage string

const (
	EntityStagePreCreated       EntityStage = "PRE_CREATED"
	EntityStageAICreating       EntityStage = "AI_CREATING"
	EntityStageAICreationFailed EntityStage = "AI_CREATION_FAILED"
	EntityStageAINeedsReview    EntityStage = "AI_NEEDS_REVIEW"
	EntityStageDraft            EntityStage = "DRAFT"
	EntityStageReadyForProposal EntityStage = "READY_FOR_PROPOSAL"
	EntityStageImplementation   EntityStage = "IMPLEMENTATION"
	EntityStageActive           EntityStage = "ACTIVE"
	EntityStageCompleted        EntityStage = "COMPLETED"
	EntityStageArchived         EntityStage = "ARCHIVED"
)

// Type of entity - determines the specific subtype and applicable fields
type EntityKind string

const (
	EntityKindRevocableTrust                   EntityKind = "REVOCABLE_TRUST"
	EntityKindIrrevocableTrust                 EntityKind = "IRREVOCABLE_TRUST"
	EntityKindSlatTrust                        EntityKind = "SLAT_TRUST"
	EntityKindIlitTrust                        EntityKind = "ILIT_TRUST"
	EntityKindQprtTrust                        EntityKind = "QPRT_TRUST"
	EntityKindGratTrust                        EntityKind = "GRAT_TRUST"
	EntityKindCrtTrust                         EntityKind = "CRT_TRUST"
	EntityKindCltTrust                         EntityKind = "CLT_TRUST"
	EntityKindIndividualPersonalAccount        EntityKind = "INDIVIDUAL_PERSONAL_ACCOUNT"
	EntityKindJointPersonalAccount             EntityKind = "JOINT_PERSONAL_ACCOUNT"
	EntityKindCustodialPersonalAccount         EntityKind = "CUSTODIAL_PERSONAL_ACCOUNT"
	EntityKindInsurancePersonalAccount         EntityKind = "INSURANCE_PERSONAL_ACCOUNT"
	EntityKindQualifiedTuitionPersonalAccount  EntityKind = "QUALIFIED_TUITION_PERSONAL_ACCOUNT"
	EntityKindRetirementPersonalAccount        EntityKind = "RETIREMENT_PERSONAL_ACCOUNT"
	EntityKindDonorAdvisedFund                 EntityKind = "DONOR_ADVISED_FUND"
	EntityKindPrivateFoundation                EntityKind = "PRIVATE_FOUNDATION"
	EntityKindLlcBusinessEntity                EntityKind = "LLC_BUSINESS_ENTITY"
	EntityKindLpBusinessEntity                 EntityKind = "LP_BUSINESS_ENTITY"
	EntityKindGpBusinessEntity                 EntityKind = "GP_BUSINESS_ENTITY"
	EntityKindSoleProprietorshipBusinessEntity EntityKind = "SOLE_PROPRIETORSHIP_BUSINESS_ENTITY"
	EntityKindScorpBusinessEntity              EntityKind = "SCORP_BUSINESS_ENTITY"
	EntityKindCcorpBusinessEntity              EntityKind = "CCORP_BUSINESS_ENTITY"
)

type EntityList struct {
	Data     []Entity           `json:"data,required"`
	PageInfo EntityListPageInfo `json:"page_info,required"`
	// Total number of items matching the query (across all pages)
	TotalCount int64 `json:"total_count,required"`
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
func (r EntityList) RawJSON() string { return r.JSON.raw }
func (r *EntityList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityListPageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"has_next_page,required"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"has_previous_page,required"`
	// Cursor pointing to the last item in the current page
	EndCursor string `json:"end_cursor,nullable"`
	// Cursor pointing to the first item in the current page
	StartCursor string `json:"start_cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage     respjson.Field
		HasPreviousPage respjson.Field
		EndCursor       respjson.Field
		StartCursor     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityListPageInfo) RawJSON() string { return r.JSON.raw }
func (r *EntityListPageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityListParams struct {
	// Cursor for forward pagination. Returns items after this cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returns items before this cursor.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter entities by household ID
	HouseholdID param.Opt[string] `query:"household_id,omitzero" json:"-"`
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

// URLQuery serializes [EntityListParams]'s query parameters as `url.Values`.
func (r EntityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
