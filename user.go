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

// UserService contains methods and other services that help with interacting with
// the luminary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// Retrieve detailed information about a specific user
func (r *UserService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *UserGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a paginated list of users using cursor-based pagination
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *UserListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type UserGetResponse struct {
	// Unique identifier with user\_ prefix
	ID string `json:"id,required"`
	// Timestamp when the user was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Email address of the user
	Email string `json:"email,required" format:"email"`
	// First name of the user
	FirstName string `json:"first_name,required"`
	// Last name of the user
	LastName string `json:"last_name,required"`
	// Timestamp when the user was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListResponse struct {
	Data     []UserListResponseData   `json:"data,required"`
	PageInfo UserListResponsePageInfo `json:"page_info,required"`
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
func (r UserListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListResponseData struct {
	// Unique identifier with user\_ prefix
	ID string `json:"id,required"`
	// Timestamp when the user was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Email address of the user
	Email string `json:"email,required" format:"email"`
	// First name of the user
	FirstName string `json:"first_name,required"`
	// Last name of the user
	LastName string `json:"last_name,required"`
	// Timestamp when the user was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListResponseData) RawJSON() string { return r.JSON.raw }
func (r *UserListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListResponsePageInfo struct {
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
func (r UserListResponsePageInfo) RawJSON() string { return r.JSON.raw }
func (r *UserListResponsePageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListParams struct {
	// Cursor for forward pagination. Returns items after this cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returns items before this cursor.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
