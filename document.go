// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package withluminary

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/withluminary/go-sdk/internal/apiform"
	"github.com/withluminary/go-sdk/internal/apijson"
	"github.com/withluminary/go-sdk/internal/apiquery"
	"github.com/withluminary/go-sdk/internal/requestconfig"
	"github.com/withluminary/go-sdk/option"
	"github.com/withluminary/go-sdk/packages/pagination"
	"github.com/withluminary/go-sdk/packages/param"
	"github.com/withluminary/go-sdk/packages/respjson"
)

// DocumentService contains methods and other services that help with interacting
// with the luminary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentService] method instead.
type DocumentService struct {
	Options []option.RequestOption
}

// NewDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentService(opts ...option.RequestOption) (r DocumentService) {
	r = DocumentService{}
	r.Options = opts
	return
}

// Create a new document with file content
func (r *DocumentService) New(ctx context.Context, body DocumentNewParams, opts ...option.RequestOption) (res *Document, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "documents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve document metadata
func (r *DocumentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Document, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("documents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update document metadata only
func (r *DocumentService) Update(ctx context.Context, id string, body DocumentUpdateParams, opts ...option.RequestOption) (res *Document, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("documents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of documents using cursor-based pagination
func (r *DocumentService) List(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Document], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "documents"
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

// Retrieve a paginated list of documents using cursor-based pagination
func (r *DocumentService) ListAutoPaging(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Document] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Soft delete a document (marks as deleted but preserves data)
func (r *DocumentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("documents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Download the binary content of the document file
func (r *DocumentService) Download(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("documents/%s/download", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve all summaries associated with a specific document
func (r *DocumentService) GetSummaries(ctx context.Context, id string, opts ...option.RequestOption) (res *DocumentGetSummariesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("documents/%s/document-summaries", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Document struct {
	// Unique identifier with document\_ prefix
	ID string `json:"id" api:"required"`
	// Timestamp when the document was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Household ID this document belongs to
	HouseholdID string `json:"household_id" api:"required"`
	// Display name of the document
	Name string `json:"name" api:"required"`
	// Type of document
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
	Type DocumentType `json:"type" api:"required"`
	// Timestamp when the document was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Whether this document should be used for AI suggestions
	EnableAISuggestions bool `json:"enable_ai_suggestions"`
	// Entity ID if this document is owned by an entity
	EntityID string `json:"entity_id" api:"nullable"`
	// Individual ID if this document is associated with an individual
	IndividualID string `json:"individual_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		HouseholdID         respjson.Field
		Name                respjson.Field
		Type                respjson.Field
		UpdatedAt           respjson.Field
		EnableAISuggestions respjson.Field
		EntityID            respjson.Field
		IndividualID        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Document) RawJSON() string { return r.JSON.raw }
func (r *Document) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentList struct {
	Data     []Document `json:"data" api:"required"`
	PageInfo PageInfo   `json:"page_info" api:"required"`
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
func (r DocumentList) RawJSON() string { return r.JSON.raw }
func (r *DocumentList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of document
type DocumentType string

const (
	DocumentTypeGratDesignSummary                    DocumentType = "GRAT_DESIGN_SUMMARY"
	DocumentTypeGeneratedPresentation                DocumentType = "GENERATED_PRESENTATION"
	DocumentTypeAssetValuation                       DocumentType = "ASSET_VALUATION"
	DocumentTypeSignedTrustDocument                  DocumentType = "SIGNED_TRUST_DOCUMENT"
	DocumentTypeTrustAmendment                       DocumentType = "TRUST_AMENDMENT"
	DocumentTypeTransferConfirmation                 DocumentType = "TRANSFER_CONFIRMATION"
	DocumentTypeExistingRemainderTrustDocument       DocumentType = "EXISTING_REMAINDER_TRUST_DOCUMENT"
	DocumentTypeBalanceSheet                         DocumentType = "BALANCE_SHEET"
	DocumentTypeWill                                 DocumentType = "WILL"
	DocumentTypeWillCodicil                          DocumentType = "WILL_CODICIL"
	DocumentTypePowerOfAttorney                      DocumentType = "POWER_OF_ATTORNEY"
	DocumentTypeAssignmentOfInterest                 DocumentType = "ASSIGNMENT_OF_INTEREST"
	DocumentTypeAssignmentOfTangibleProperty         DocumentType = "ASSIGNMENT_OF_TANGIBLE_PROPERTY"
	DocumentTypeLoanNoteAgreement                    DocumentType = "LOAN_NOTE_AGREEMENT"
	DocumentTypeArticlesOfIncorporation              DocumentType = "ARTICLES_OF_INCORPORATION"
	DocumentTypeOperatingAgreement                   DocumentType = "OPERATING_AGREEMENT"
	DocumentTypePartnershipAgreement                 DocumentType = "PARTNERSHIP_AGREEMENT"
	DocumentTypeAccountDocumentationStatement        DocumentType = "ACCOUNT_DOCUMENTATION_STATEMENT"
	DocumentTypeTaxIDConfirmation                    DocumentType = "TAX_ID_CONFIRMATION"
	DocumentTypeGiftTaxReturn                        DocumentType = "GIFT_TAX_RETURN"
	DocumentTypeIncomeTaxReturn                      DocumentType = "INCOME_TAX_RETURN"
	DocumentTypeTaxReceipt                           DocumentType = "TAX_RECEIPT"
	DocumentTypeTaxFiling                            DocumentType = "TAX_FILING"
	DocumentTypeCorporateBylaws                      DocumentType = "CORPORATE_BYLAWS"
	DocumentTypeLlcAgreement                         DocumentType = "LLC_AGREEMENT"
	DocumentTypeLlcAgreementAmendment                DocumentType = "LLC_AGREEMENT_AMENDMENT"
	DocumentTypeOperatingAgreementAmendment          DocumentType = "OPERATING_AGREEMENT_AMENDMENT"
	DocumentTypePartnershipAgreementAmendment        DocumentType = "PARTNERSHIP_AGREEMENT_AMENDMENT"
	DocumentTypeShareholdersAgreement                DocumentType = "SHAREHOLDERS_AGREEMENT"
	DocumentTypeStateBusinessFiling                  DocumentType = "STATE_BUSINESS_FILING"
	DocumentTypeLoggedContribution                   DocumentType = "LOGGED_CONTRIBUTION"
	DocumentTypeLoggedDistribution                   DocumentType = "LOGGED_DISTRIBUTION"
	DocumentTypeInsurancePolicy                      DocumentType = "INSURANCE_POLICY"
	DocumentTypeCrummeyLetter                        DocumentType = "CRUMMEY_LETTER"
	DocumentTypeInsurancePremiumPayment              DocumentType = "INSURANCE_PREMIUM_PAYMENT"
	DocumentTypeBeneficialOwnershipInformationReport DocumentType = "BENEFICIAL_OWNERSHIP_INFORMATION_REPORT"
	DocumentTypeFincenFiling                         DocumentType = "FINCEN_FILING"
	DocumentTypeHealthcareProxy                      DocumentType = "HEALTHCARE_PROXY"
	DocumentTypeLivingWill                           DocumentType = "LIVING_WILL"
	DocumentTypeDriversLicense                       DocumentType = "DRIVERS_LICENSE"
	DocumentTypePassport                             DocumentType = "PASSPORT"
	DocumentTypeDeed                                 DocumentType = "DEED"
	DocumentTypeOther                                DocumentType = "OTHER"
)

type DocumentGetSummariesResponse struct {
	Data []DocumentSummary `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentGetSummariesResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentGetSummariesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentNewParams struct {
	// The document file to upload
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	// Household ID this document belongs to
	HouseholdID string `json:"household_id" api:"required"`
	// Display name of the document
	Name string `json:"name" api:"required"`
	// Type of document
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
	Type DocumentType `json:"type,omitzero" api:"required"`
	// Whether this document should be used for AI suggestions
	EnableAISuggestions param.Opt[bool] `json:"enable_ai_suggestions,omitzero"`
	// Entity ID if this document is owned by an entity
	EntityID param.Opt[string] `json:"entity_id,omitzero"`
	// Individual ID if associated with an individual
	IndividualID param.Opt[string] `json:"individual_id,omitzero"`
	paramObj
}

func (r DocumentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type DocumentUpdateParams struct {
	// Entity ID if this document is owned by an entity
	EntityID param.Opt[string] `json:"entity_id,omitzero"`
	// Individual ID if associated with an individual
	IndividualID param.Opt[string] `json:"individual_id,omitzero"`
	// Whether this document should be used for AI suggestions
	EnableAISuggestions param.Opt[bool] `json:"enable_ai_suggestions,omitzero"`
	// Display name of the document
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of document
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
	Type DocumentType `json:"type,omitzero"`
	paramObj
}

func (r DocumentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentListParams struct {
	// Cursor for forward pagination. Returns items after this cursor.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returns items before this cursor.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter documents by household ID
	HouseholdID param.Opt[string] `query:"household_id,omitzero" json:"-"`
	// Maximum number of items to return per page.
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

// URLQuery serializes [DocumentListParams]'s query parameters as `url.Values`.
func (r DocumentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
