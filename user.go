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
func (r *UserService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *User, err error) {
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
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[User], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "users"
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

// Retrieve a paginated list of users using cursor-based pagination
func (r *UserService) ListAutoPaging(ctx context.Context, query UserListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[User] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

type User struct {
	// Unique identifier with user\_ prefix
	ID string `json:"id" api:"required"`
	// Timestamp when the user was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Email address of the user
	Email string `json:"email" api:"required" format:"email"`
	// First name of the user
	FirstName string `json:"first_name" api:"required"`
	// Last name of the user
	LastName string `json:"last_name" api:"required"`
	// Timestamp when the user was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
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
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
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
