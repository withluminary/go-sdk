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

// DocumentSummaryService contains methods and other services that help with
// interacting with the luminary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentSummaryService] method instead.
type DocumentSummaryService struct {
	Options []option.RequestOption
}

// NewDocumentSummaryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentSummaryService(opts ...option.RequestOption) (r DocumentSummaryService) {
	r = DocumentSummaryService{}
	r.Options = opts
	return
}

// Retrieve a specific document summary
func (r *DocumentSummaryService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DocumentSummary, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("document-summaries/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing document summary
func (r *DocumentSummaryService) Update(ctx context.Context, id string, body DocumentSummaryUpdateParams, opts ...option.RequestOption) (res *DocumentSummary, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("document-summaries/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of document summaries using cursor-based pagination
func (r *DocumentSummaryService) List(ctx context.Context, query DocumentSummaryListParams, opts ...option.RequestOption) (res *DocumentSummaryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "document-summaries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Download the document summary content in the specified format
func (r *DocumentSummaryService) Download(ctx context.Context, id string, query DocumentSummaryDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("document-summaries/%s/download", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DocumentSummary struct {
	// Unique identifier for the document summary
	ID string `json:"id,required"`
	// Timestamp when the summary was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Display name for the summary
	DisplayName string `json:"display_name,required"`
	// ID of the document this summary belongs to
	DocumentID string `json:"document_id,required"`
	// ID of the household this summary belongs to
	HouseholdID string `json:"household_id,required"`
	// The summary text content
	Summary string `json:"summary,required"`
	// Timestamp when the summary was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Indicates if the summary was AI-generated or user-entered
	//
	// Any of "AI_AUTO", "USER".
	EntryMode DocumentSummaryEntryMode `json:"entry_mode"`
	// Format of the summary content
	//
	// Any of "MARKDOWN", "PLAIN_TEXT".
	SummaryFormat DocumentSummaryFormat `json:"summary_format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		DisplayName   respjson.Field
		DocumentID    respjson.Field
		HouseholdID   respjson.Field
		Summary       respjson.Field
		UpdatedAt     respjson.Field
		EntryMode     respjson.Field
		SummaryFormat respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentSummary) RawJSON() string { return r.JSON.raw }
func (r *DocumentSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates if the summary was AI-generated or user-entered
type DocumentSummaryEntryMode string

const (
	DocumentSummaryEntryModeAIAuto DocumentSummaryEntryMode = "AI_AUTO"
	DocumentSummaryEntryModeUser   DocumentSummaryEntryMode = "USER"
)

// Format of the summary content
type DocumentSummaryFormat string

const (
	DocumentSummaryFormatMarkdown  DocumentSummaryFormat = "MARKDOWN"
	DocumentSummaryFormatPlainText DocumentSummaryFormat = "PLAIN_TEXT"
)

type PageInfo struct {
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
func (r PageInfo) RawJSON() string { return r.JSON.raw }
func (r *PageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentSummaryListResponse struct {
	Data     []DocumentSummary `json:"data,required"`
	PageInfo PageInfo          `json:"page_info,required"`
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
func (r DocumentSummaryListResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentSummaryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentSummaryUpdateParams struct {
	// Display name for the summary
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// The summary text content
	Summary param.Opt[string] `json:"summary,omitzero"`
	// Indicates if the summary was AI-generated or user-entered
	//
	// Any of "AI_AUTO", "USER".
	EntryMode DocumentSummaryEntryMode `json:"entry_mode,omitzero"`
	// Format of the summary content
	//
	// Any of "MARKDOWN", "PLAIN_TEXT".
	SummaryFormat DocumentSummaryFormat `json:"summary_format,omitzero"`
	paramObj
}

func (r DocumentSummaryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentSummaryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentSummaryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentSummaryListParams struct {
	// Cursor for forward pagination. Returns items after this cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returns items before this cursor.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter summaries by document ID
	DocumentID param.Opt[string] `query:"document_id,omitzero" json:"-"`
	// Filter summaries by household ID
	HouseholdID param.Opt[string] `query:"household_id,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentSummaryListParams]'s query parameters as
// `url.Values`.
func (r DocumentSummaryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentSummaryDownloadParams struct {
	// Output format for the download
	//
	// Any of "pdf".
	Format DocumentSummaryDownloadParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentSummaryDownloadParams]'s query parameters as
// `url.Values`.
func (r DocumentSummaryDownloadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Output format for the download
type DocumentSummaryDownloadParamsFormat string

const (
	DocumentSummaryDownloadParamsFormatPdf DocumentSummaryDownloadParamsFormat = "pdf"
)
