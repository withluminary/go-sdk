// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package withluminary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/withluminary/go-sdk/internal/apijson"
	"github.com/withluminary/go-sdk/internal/requestconfig"
	"github.com/withluminary/go-sdk/option"
	"github.com/withluminary/go-sdk/packages/param"
	"github.com/withluminary/go-sdk/packages/respjson"
)

// EntityValuationService contains methods and other services that help with
// interacting with the luminary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityValuationService] method instead.
type EntityValuationService struct {
	Options []option.RequestOption
}

// NewEntityValuationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEntityValuationService(opts ...option.RequestOption) (r EntityValuationService) {
	r = EntityValuationService{}
	r.Options = opts
	return
}

// Add a new valuation to the entity's history
func (r *EntityValuationService) New(ctx context.Context, id string, body EntityValuationNewParams, opts ...option.RequestOption) (res *Valuation, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/valuation", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the most recent valuation with flattened asset values by type
func (r *EntityValuationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Valuation, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/valuation", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Valuation struct {
	// Unique identifier with valuationv2\_ prefix
	ID string `json:"id,required"`
	// Timestamp when the valuation was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Total value of all directly held assets in USD
	DirectlyHeldAssetValue float64 `json:"directly_held_asset_value,required"`
	// List of individual assets in this valuation
	DirectlyHeldAssets []ValuationDirectlyHeldAsset `json:"directly_held_assets,required"`
	// The date this valuation is effective
	EffectiveDate time.Time `json:"effective_date,required" format:"date"`
	// Entity ID this valuation belongs to
	EntityID string `json:"entity_id,required"`
	// Total value of all assets minus liabilities in USD
	TotalValue float64 `json:"total_value,required"`
	// Timestamp when the valuation was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Free-form notes about this valuation
	Description string `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CreatedAt              respjson.Field
		DirectlyHeldAssetValue respjson.Field
		DirectlyHeldAssets     respjson.Field
		EffectiveDate          respjson.Field
		EntityID               respjson.Field
		TotalValue             respjson.Field
		UpdatedAt              respjson.Field
		Description            respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Valuation) RawJSON() string { return r.JSON.raw }
func (r *Valuation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ValuationDirectlyHeldAsset struct {
	// Asset ID
	ID         string                               `json:"id,required"`
	AssetClass ValuationDirectlyHeldAssetAssetClass `json:"asset_class,required"`
	// Display name of the asset
	DisplayName string `json:"display_name,required"`
	// Value of this asset in USD
	Value float64 `json:"value,required"`
	// External ID from the static asset (if available)
	ExternalID string `json:"external_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AssetClass  respjson.Field
		DisplayName respjson.Field
		Value       respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ValuationDirectlyHeldAsset) RawJSON() string { return r.JSON.raw }
func (r *ValuationDirectlyHeldAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ValuationDirectlyHeldAssetAssetClass struct {
	// Asset class ID
	ID string `json:"id,required"`
	// Display name of the asset class
	DisplayName string `json:"display_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DisplayName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ValuationDirectlyHeldAssetAssetClass) RawJSON() string { return r.JSON.raw }
func (r *ValuationDirectlyHeldAssetAssetClass) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityValuationNewParams struct {
	// List of assets to include in this valuation
	DirectlyHeldAssets []EntityValuationNewParamsDirectlyHeldAsset `json:"directly_held_assets,omitzero,required"`
	// The date this valuation is effective
	EffectiveDate time.Time `json:"effective_date,required" format:"date"`
	// Free-form notes about this valuation
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r EntityValuationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EntityValuationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntityValuationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AssetClassID, DisplayName, Value are required.
type EntityValuationNewParamsDirectlyHeldAsset struct {
	// Asset class ID to associate with this asset
	AssetClassID string `json:"asset_class_id,required"`
	// Display name of the asset
	DisplayName string `json:"display_name,required"`
	// Value of this asset in USD
	Value float64 `json:"value,required"`
	// External ID for the asset
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	paramObj
}

func (r EntityValuationNewParamsDirectlyHeldAsset) MarshalJSON() (data []byte, err error) {
	type shadow EntityValuationNewParamsDirectlyHeldAsset
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntityValuationNewParamsDirectlyHeldAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
