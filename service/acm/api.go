// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package acm provides a client for AWS Certificate Manager.
package acm

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

const opAddTagsToCertificate = "AddTagsToCertificate"

// AddTagsToCertificateRequest generates a request for the AddTagsToCertificate operation.
func (c *ACM) AddTagsToCertificateRequest(input *AddTagsToCertificateInput) (req *request.Request, output *AddTagsToCertificateOutput) {
	op := &request.Operation{
		Name:       opAddTagsToCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddTagsToCertificateInput{}
	}

	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	output = &AddTagsToCertificateOutput{}
	req.Data = output
	return
}

// Adds one or more tags to an ACM Certificate. Tags are labels that you can
// use to identify and organize your AWS resources. Each tag consists of a key
// and an optional value. You specify the certificate on input by its Amazon
// Resource Name (ARN). You specify the tag by using a key-value pair.
//
//  You can apply a tag to just one certificate if you want to identify a specific
// characteristic of that certificate, or you can apply the same tag to multiple
// certificates if you want to filter for a common relationship among those
// certificates. Similarly, you can apply the same tag to multiple resources
// if you want to specify a relationship among those resources. For example,
// you can add the same tag to an ACM Certificate and an Elastic Load Balancing
// load balancer to indicate that they are both used by the same website. For
// more information, see Tagging ACM Certificates (http://docs.aws.amazon.com/acm/latest/userguide/tags.html).
//
// To remove one or more tags, use the RemoveTagsFromCertificate action. To
// view all of the tags that have been applied to the certificate, use the ListTagsForCertificate
// action.
func (c *ACM) AddTagsToCertificate(input *AddTagsToCertificateInput) (*AddTagsToCertificateOutput, error) {
	req, out := c.AddTagsToCertificateRequest(input)
	err := req.Send()
	return out, err
}

const opDeleteCertificate = "DeleteCertificate"

// DeleteCertificateRequest generates a request for the DeleteCertificate operation.
func (c *ACM) DeleteCertificateRequest(input *DeleteCertificateInput) (req *request.Request, output *DeleteCertificateOutput) {
	op := &request.Operation{
		Name:       opDeleteCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCertificateInput{}
	}

	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	output = &DeleteCertificateOutput{}
	req.Data = output
	return
}

// Deletes an ACM Certificate and its associated private key. If this action
// succeeds, the certificate no longer appears in the list of ACM Certificates
// that can be displayed by calling the ListCertificates action or be retrieved
// by calling the GetCertificate action. The certificate will not be available
// for use by other AWS services.
//
// You cannot delete an ACM Certificate that is being used by another AWS service.
// To delete a certificate that is in use, the certificate association must
// first be removed.
func (c *ACM) DeleteCertificate(input *DeleteCertificateInput) (*DeleteCertificateOutput, error) {
	req, out := c.DeleteCertificateRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeCertificate = "DescribeCertificate"

// DescribeCertificateRequest generates a request for the DescribeCertificate operation.
func (c *ACM) DescribeCertificateRequest(input *DescribeCertificateInput) (req *request.Request, output *DescribeCertificateOutput) {
	op := &request.Operation{
		Name:       opDescribeCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCertificateInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeCertificateOutput{}
	req.Data = output
	return
}

// Returns a list of the fields contained in the specified ACM Certificate.
// For example, this action returns the certificate status, a flag that indicates
// whether the certificate is associated with any other AWS service, and the
// date at which the certificate request was created. You specify the ACM Certificate
// on input by its Amazon Resource Name (ARN).
func (c *ACM) DescribeCertificate(input *DescribeCertificateInput) (*DescribeCertificateOutput, error) {
	req, out := c.DescribeCertificateRequest(input)
	err := req.Send()
	return out, err
}

const opGetCertificate = "GetCertificate"

// GetCertificateRequest generates a request for the GetCertificate operation.
func (c *ACM) GetCertificateRequest(input *GetCertificateInput) (req *request.Request, output *GetCertificateOutput) {
	op := &request.Operation{
		Name:       opGetCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCertificateInput{}
	}

	req = c.newRequest(op, input, output)
	output = &GetCertificateOutput{}
	req.Data = output
	return
}

// Retrieves an ACM Certificate and certificate chain for the certificate specified
// by an ARN. The chain is an ordered list of certificates that contains the
// root certificate, intermediate certificates of subordinate CAs, and the ACM
// Certificate. The certificate and certificate chain are base64 encoded. If
// you want to decode the certificate chain to see the individual certificate
// fields, you can use OpenSSL.
//
//  Currently, ACM Certificates can be used only with Elastic Load Balancing
// and Amazon CloudFront.
func (c *ACM) GetCertificate(input *GetCertificateInput) (*GetCertificateOutput, error) {
	req, out := c.GetCertificateRequest(input)
	err := req.Send()
	return out, err
}

const opListCertificates = "ListCertificates"

// ListCertificatesRequest generates a request for the ListCertificates operation.
func (c *ACM) ListCertificatesRequest(input *ListCertificatesInput) (req *request.Request, output *ListCertificatesOutput) {
	op := &request.Operation{
		Name:       opListCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListCertificatesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListCertificatesOutput{}
	req.Data = output
	return
}

// Retrieves a list of the ACM Certificate ARNs, and the domain name for each
// ARN, owned by the calling account. You can filter the list based on the CertificateStatuses
// parameter, and you can display up to MaxItems certificates at one time. If
// you have more than MaxItems certificates, use the NextToken marker from the
// response object in your next call to the ListCertificates action to retrieve
// the next set of certificate ARNs.
func (c *ACM) ListCertificates(input *ListCertificatesInput) (*ListCertificatesOutput, error) {
	req, out := c.ListCertificatesRequest(input)
	err := req.Send()
	return out, err
}

func (c *ACM) ListCertificatesPages(input *ListCertificatesInput, fn func(p *ListCertificatesOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.ListCertificatesRequest(input)
	page.Handlers.Build.PushBack(request.MakeAddToUserAgentFreeFormHandler("Paginator"))
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*ListCertificatesOutput), lastPage)
	})
}

const opListTagsForCertificate = "ListTagsForCertificate"

// ListTagsForCertificateRequest generates a request for the ListTagsForCertificate operation.
func (c *ACM) ListTagsForCertificateRequest(input *ListTagsForCertificateInput) (req *request.Request, output *ListTagsForCertificateOutput) {
	op := &request.Operation{
		Name:       opListTagsForCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsForCertificateInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListTagsForCertificateOutput{}
	req.Data = output
	return
}

// Lists the tags that have been applied to the ACM Certificate. Use the certificate
// ARN to specify the certificate. To add a tag to an ACM Certificate, use the
// AddTagsToCertificate action. To delete a tag, use the RemoveTagsFromCertificate
// action.
func (c *ACM) ListTagsForCertificate(input *ListTagsForCertificateInput) (*ListTagsForCertificateOutput, error) {
	req, out := c.ListTagsForCertificateRequest(input)
	err := req.Send()
	return out, err
}

const opRemoveTagsFromCertificate = "RemoveTagsFromCertificate"

// RemoveTagsFromCertificateRequest generates a request for the RemoveTagsFromCertificate operation.
func (c *ACM) RemoveTagsFromCertificateRequest(input *RemoveTagsFromCertificateInput) (req *request.Request, output *RemoveTagsFromCertificateOutput) {
	op := &request.Operation{
		Name:       opRemoveTagsFromCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveTagsFromCertificateInput{}
	}

	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	output = &RemoveTagsFromCertificateOutput{}
	req.Data = output
	return
}

// Remove one or more tags from an ACM Certificate. A tag consists of a key-value
// pair. If you do not specify the value portion of the tag when calling this
// function, the tag will be removed regardless of value. If you specify a value,
// the tag is removed only if it is associated with the specified value.
//
// To add tags to a certificate, use the AddTagsToCertificate action. To view
// all of the tags that have been applied to a specific ACM Certificate, use
// the ListTagsForCertificate action.
func (c *ACM) RemoveTagsFromCertificate(input *RemoveTagsFromCertificateInput) (*RemoveTagsFromCertificateOutput, error) {
	req, out := c.RemoveTagsFromCertificateRequest(input)
	err := req.Send()
	return out, err
}

const opRequestCertificate = "RequestCertificate"

// RequestCertificateRequest generates a request for the RequestCertificate operation.
func (c *ACM) RequestCertificateRequest(input *RequestCertificateInput) (req *request.Request, output *RequestCertificateOutput) {
	op := &request.Operation{
		Name:       opRequestCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RequestCertificateInput{}
	}

	req = c.newRequest(op, input, output)
	output = &RequestCertificateOutput{}
	req.Data = output
	return
}

// Requests an ACM Certificate for use with other AWS services. To request an
// ACM Certificate, you must specify the fully qualified domain name (FQDN)
// for your site. You can also specify additional FQDNs if users can reach your
// site by using other names. For each domain name you specify, email is sent
// to the domain owner to request approval to issue the certificate. After receiving
// approval from the domain owner, the ACM Certificate is issued. For more information,
// see the  AWS Certificate Manager User Guide  (http://docs.aws.amazon.com/acm/latest/userguide/overview.html).
func (c *ACM) RequestCertificate(input *RequestCertificateInput) (*RequestCertificateOutput, error) {
	req, out := c.RequestCertificateRequest(input)
	err := req.Send()
	return out, err
}

const opResendValidationEmail = "ResendValidationEmail"

// ResendValidationEmailRequest generates a request for the ResendValidationEmail operation.
func (c *ACM) ResendValidationEmailRequest(input *ResendValidationEmailInput) (req *request.Request, output *ResendValidationEmailOutput) {
	op := &request.Operation{
		Name:       opResendValidationEmail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResendValidationEmailInput{}
	}

	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	output = &ResendValidationEmailOutput{}
	req.Data = output
	return
}

// Resends the email that requests domain ownership validation. The domain owner
// or an authorized representative must approve the ACM Certificate before it
// can be issued. The certificate can be approved by clicking a link in the
// mail to navigate to the Amazon certificate approval website and then clicking
// I Approve. However, the validation email can be blocked by spam filters.
// Therefore, if you do not receive the original mail, you can request that
// the mail be resent within 72 hours of requesting the ACM Certificate. If
// more than 72 hours have elapsed since your original request or since your
// last attempt to resend validation mail, you must request a new certificate.
func (c *ACM) ResendValidationEmail(input *ResendValidationEmailInput) (*ResendValidationEmailOutput, error) {
	req, out := c.ResendValidationEmailRequest(input)
	err := req.Send()
	return out, err
}

type AddTagsToCertificateInput struct {
	_ struct{} `type:"structure"`

	// String that contains the ARN of the ACM Certificate to which the tag is to
	// be applied. This must be of the form:
	//
	//  arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	//  For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	CertificateArn *string `min:"20" type:"string" required:"true"`

	// The key-value pair that defines the tag. The tag value is optional.
	Tags []*Tag `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s AddTagsToCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AddTagsToCertificateInput) GoString() string {
	return s.String()
}

type AddTagsToCertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddTagsToCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AddTagsToCertificateOutput) GoString() string {
	return s.String()
}

// This structure is returned in the response object of the DescribeCertificate
// action.
type CertificateDetail struct {
	_ struct{} `type:"structure"`

	// Amazon Resource Name (ARN) of the certificate. This is of the form:
	//
	//  arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	//  For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	CertificateArn *string `min:"20" type:"string"`

	// Time at which the certificate was requested.
	CreatedAt *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Fully qualified domain name (FQDN), such as www.example.com or example.com,
	// for the certificate.
	DomainName *string `min:"1" type:"string"`

	// References a DomainValidation structure that contains the domain name in
	// the certificate and the email address that can be used for validation.
	DomainValidationOptions []*DomainValidation `min:"1" type:"list"`

	// List that identifies ARNs that are using the certificate. A single ACM Certificate
	// can be used by multiple AWS resources.
	InUseBy []*string `type:"list"`

	// Time at which the certificate was issued.
	IssuedAt *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The X.500 distinguished name of the CA that issued and signed the certificate.
	Issuer *string `type:"string"`

	// Asymmetric algorithm used to generate the public and private key pair. Currently
	// the only supported value is RSA_2048.
	KeyAlgorithm *string `type:"string" enum:"KeyAlgorithm"`

	// Time after which the certificate is not valid.
	NotAfter *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Time before which the certificate is not valid.
	NotBefore *time.Time `type:"timestamp" timestampFormat:"unix"`

	// A RevocationReason enumeration value that indicates why the certificate was
	// revoked. This value exists only if the certificate has been revoked. This
	// can be one of the following vales:  UNSPECIFIED
	//
	// KEY_COMPROMISE
	//
	// CA_COMPROMISE
	//
	// AFFILIATION_CHANGED
	//
	// SUPERCEDED
	//
	// CESSATION_OF_OPERATION
	//
	// CERTIFICATE_HOLD
	//
	// REMOVE_FROM_CRL
	//
	// PRIVILEGE_WITHDRAWN
	//
	// A_A_COMPROMISE
	RevocationReason *string `type:"string" enum:"RevocationReason"`

	// The time, if any, at which the certificate was revoked. This value exists
	// only if the certificate has been revoked.
	RevokedAt *time.Time `type:"timestamp" timestampFormat:"unix"`

	// String that contains the serial number of the certificate.
	Serial *string `type:"string"`

	// Algorithm used to generate a signature. Currently the only supported value
	// is SHA256WITHRSA.
	SignatureAlgorithm *string `type:"string"`

	// A CertificateStatus enumeration value that can contain one of the following:
	//  PENDING_VALIDATION
	//
	// ISSUED
	//
	// INACTIVE
	//
	// EXPIRED
	//
	// REVOKED
	//
	// FAILED
	//
	// VALIDATION_TIMED_OUT
	Status *string `type:"string" enum:"CertificateStatus"`

	// The X.500 distinguished name of the entity associated with the public key
	// contained in the certificate.
	Subject *string `type:"string"`

	// One or more domain names (subject alternative names) included in the certificate
	// request. After the certificate is issued, this list includes the domain names
	// bound to the public key contained in the certificate. The subject alternative
	// names include the canonical domain name (CN) of the certificate and additional
	// domain names that can be used to connect to the website.
	SubjectAlternativeNames []*string `min:"1" type:"list"`
}

// String returns the string representation
func (s CertificateDetail) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CertificateDetail) GoString() string {
	return s.String()
}

// This structure is returned in the response object of ListCertificates action.
type CertificateSummary struct {
	_ struct{} `type:"structure"`

	// Amazon Resource Name (ARN) of the certificate. This is of the form:
	//
	//  arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	//  For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	CertificateArn *string `min:"20" type:"string"`

	// Fully qualified domain name (FQDN), such as www.example.com or example.com,
	// for the certificate.
	DomainName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CertificateSummary) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CertificateSummary) GoString() string {
	return s.String()
}

type DeleteCertificateInput struct {
	_ struct{} `type:"structure"`

	// String that contains the ARN of the ACM Certificate to be deleted. This must
	// be of the form:
	//
	//  arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	//  For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	CertificateArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCertificateInput) GoString() string {
	return s.String()
}

type DeleteCertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCertificateOutput) GoString() string {
	return s.String()
}

type DescribeCertificateInput struct {
	_ struct{} `type:"structure"`

	// String that contains an ACM Certificate ARN. The ARN must be of the form:
	//
	//  arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	//  For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	CertificateArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCertificateInput) GoString() string {
	return s.String()
}

type DescribeCertificateOutput struct {
	_ struct{} `type:"structure"`

	// Contains a CertificateDetail structure that lists the fields of an ACM Certificate.
	Certificate *CertificateDetail `type:"structure"`
}

// String returns the string representation
func (s DescribeCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCertificateOutput) GoString() string {
	return s.String()
}

// Structure that contains the domain name, the base validation domain to which
// validation email is sent, and the email addresses used to validate the domain
// identity.
type DomainValidation struct {
	_ struct{} `type:"structure"`

	// Fully Qualified Domain Name (FQDN) of the form www.example.com or example.com
	DomainName *string `min:"1" type:"string" required:"true"`

	// The base validation domain that acts as the suffix of the email addresses
	// that are used to send the emails.
	ValidationDomain *string `min:"1" type:"string"`

	// A list of contact address for the domain registrant.
	ValidationEmails []*string `type:"list"`
}

// String returns the string representation
func (s DomainValidation) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainValidation) GoString() string {
	return s.String()
}

// This structure is used in the request object of the RequestCertificate action.
type DomainValidationOption struct {
	_ struct{} `type:"structure"`

	// Fully Qualified Domain Name (FQDN) of the certificate being requested.
	DomainName *string `min:"1" type:"string" required:"true"`

	// The domain to which validation email is sent. This is the base validation
	// domain that will act as the suffix of the email addresses. This must be the
	// same as the DomainName value or a superdomain of the DomainName value. For
	// example, if you requested a certificate for site.subdomain.example.com and
	// specify a ValidationDomain of subdomain.example.com, ACM sends email to the
	// domain registrant, technical contact, and administrative contact in WHOIS
	// for the base domain and the following five addresses:  admin@subdomain.example.com
	//
	// administrator@subdomain.example.com
	//
	// hostmaster@subdomain.example.com
	//
	// postmaster@subdomain.example.com
	//
	// webmaster@subdomain.example.com
	ValidationDomain *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DomainValidationOption) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainValidationOption) GoString() string {
	return s.String()
}

type GetCertificateInput struct {
	_ struct{} `type:"structure"`

	// String that contains a certificate ARN in the following format:
	//
	//  arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	//  For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	CertificateArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetCertificateInput) GoString() string {
	return s.String()
}

type GetCertificateOutput struct {
	_ struct{} `type:"structure"`

	// String that contains the ACM Certificate represented by the ARN specified
	// at input.
	Certificate *string `min:"1" type:"string"`

	// The certificate chain that contains the root certificate issued by the certificate
	// authority (CA).
	CertificateChain *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetCertificateOutput) GoString() string {
	return s.String()
}

type ListCertificatesInput struct {
	_ struct{} `type:"structure"`

	// Identifies the statuses of the ACM Certificates for which you want to retrieve
	// the ARNs. This can be one or more of the following values:  PENDING_VALIDATION
	//
	// ISSUED
	//
	// INACTIVE
	//
	// EXPIRED
	//
	// VALIDATION_TIMED_OUT
	//
	// REVOKED
	//
	// FAILED
	CertificateStatuses []*string `type:"list"`

	// Specify this parameter when paginating results to indicate the maximum number
	// of ACM Certificates that you want to display for each response. If there
	// are additional certificates beyond the maximum you specify, use the NextToken
	// value in your next call to the ListCertificates action.
	MaxItems *int64 `min:"1" type:"integer"`

	// String that contains an opaque marker of the next ACM Certificate ARN to
	// be displayed. Use this parameter when paginating results, and only in a subsequent
	// request after you've received a response where the results have been truncated.
	// Set it to an empty string the first time you call this action, and set it
	// to the value of the NextToken element you receive in the response object
	// for subsequent calls.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListCertificatesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCertificatesInput) GoString() string {
	return s.String()
}

type ListCertificatesOutput struct {
	_ struct{} `type:"structure"`

	// A list of the certificate ARNs.
	CertificateSummaryList []*CertificateSummary `type:"list"`

	// If the list has been truncated, this value is present and should be used
	// for the NextToken input parameter on your next call to ListCertificates.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListCertificatesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCertificatesOutput) GoString() string {
	return s.String()
}

type ListTagsForCertificateInput struct {
	_ struct{} `type:"structure"`

	// String that contains the ARN of the ACM Certificate for which you want to
	// list the tags. This must be of the form:
	//
	//  arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	//  For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	CertificateArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTagsForCertificateInput) GoString() string {
	return s.String()
}

type ListTagsForCertificateOutput struct {
	_ struct{} `type:"structure"`

	// The key-value pairs that define the applied tags.
	Tags []*Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s ListTagsForCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTagsForCertificateOutput) GoString() string {
	return s.String()
}

type RemoveTagsFromCertificateInput struct {
	_ struct{} `type:"structure"`

	// String that contains the ARN of the ACM Certificate with one or more tags
	// that you want to remove. This must be of the form:
	//
	//  arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	//  For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	CertificateArn *string `min:"20" type:"string" required:"true"`

	// The key-value pair that defines the tag to remove.
	Tags []*Tag `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s RemoveTagsFromCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveTagsFromCertificateInput) GoString() string {
	return s.String()
}

type RemoveTagsFromCertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveTagsFromCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveTagsFromCertificateOutput) GoString() string {
	return s.String()
}

type RequestCertificateInput struct {
	_ struct{} `type:"structure"`

	// Fully qualified domain name (FQDN), such as www.example.com, of the site
	// you want to secure with an ACM Certificate. Use an asterisk (*) to create
	// a wildcard certificate that protects several sites in the same domain. For
	// example, *.example.com protects www.example.com, site.example.com, and images.example.com.
	DomainName *string `min:"1" type:"string" required:"true"`

	// The base validation domain that will act as the suffix of the email addresses
	// that are used to send the emails. This must be the same as the Domain value
	// or a superdomain of the Domain value. For example, if you requested a certificate
	// for test.example.com and specify DomainValidationOptions of example.com,
	// ACM sends email to the domain registrant, technical contact, and administrative
	// contact in WHOIS and the following five addresses:  admin@example.com
	//
	// administrator@example.com
	//
	// hostmaster@example.com
	//
	// postmaster@example.com
	//
	// webmaster@example.com
	DomainValidationOptions []*DomainValidationOption `min:"1" type:"list"`

	// Customer chosen string that can be used to distinguish between calls to RequestCertificate.
	// Idempotency tokens time out after one hour. Therefore, if you call RequestCertificate
	// multiple times with the same idempotency token within one hour, ACM recognizes
	// that you are requesting only one certificate and will issue only one. If
	// you change the idempotency token for each call, ACM recognizes that you are
	// requesting multiple certificates.
	IdempotencyToken *string `min:"1" type:"string"`

	// Additional FQDNs to be included in the Subject Alternative Name extension
	// of the ACM Certificate. For example, add the name www.example.net to a certificate
	// for which the DomainName field is www.example.com if users can reach your
	// site by using either name.
	SubjectAlternativeNames []*string `min:"1" type:"list"`
}

// String returns the string representation
func (s RequestCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RequestCertificateInput) GoString() string {
	return s.String()
}

type RequestCertificateOutput struct {
	_ struct{} `type:"structure"`

	// String that contains the ARN of the issued certificate. This must be of the
	// form:
	//
	//  arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012
	CertificateArn *string `min:"20" type:"string"`
}

// String returns the string representation
func (s RequestCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RequestCertificateOutput) GoString() string {
	return s.String()
}

type ResendValidationEmailInput struct {
	_ struct{} `type:"structure"`

	// String that contains the ARN of the requested certificate. The certificate
	// ARN is generated and returned by the RequestCertificate action as soon as
	// the request is made. By default, using this parameter causes email to be
	// sent to all top-level domains you specified in the certificate request.
	//
	//  The ARN must be of the form:
	//
	//  arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012
	CertificateArn *string `min:"20" type:"string" required:"true"`

	// The Fully Qualified Domain Name (FQDN) of the certificate that needs to be
	// validated.
	Domain *string `min:"1" type:"string" required:"true"`

	// The base validation domain that will act as the suffix of the email addresses
	// that are used to send the emails. This must be the same as the Domain value
	// or a superdomain of the Domain value. For example, if you requested a certificate
	// for site.subdomain.example.com and specify a ValidationDomain of subdomain.example.com,
	// ACM sends email to the domain registrant, technical contact, and administrative
	// contact in WHOIS and the following five addresses:  admin@subdomain.example.com
	//
	// administrator@subdomain.example.com
	//
	// hostmaster@subdomain.example.com
	//
	// postmaster@subdomain.example.com
	//
	// webmaster@subdomain.example.com
	ValidationDomain *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ResendValidationEmailInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResendValidationEmailInput) GoString() string {
	return s.String()
}

type ResendValidationEmailOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ResendValidationEmailOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResendValidationEmailOutput) GoString() string {
	return s.String()
}

// A key-value pair that identifies or specifies metadata about an ACM resource.
type Tag struct {
	_ struct{} `type:"structure"`

	// The key of the tag.
	Key *string `min:"1" type:"string" required:"true"`

	// The value of the tag.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Tag) GoString() string {
	return s.String()
}

const (
	// @enum CertificateStatus
	CertificateStatusPendingValidation = "PENDING_VALIDATION"
	// @enum CertificateStatus
	CertificateStatusIssued = "ISSUED"
	// @enum CertificateStatus
	CertificateStatusInactive = "INACTIVE"
	// @enum CertificateStatus
	CertificateStatusExpired = "EXPIRED"
	// @enum CertificateStatus
	CertificateStatusValidationTimedOut = "VALIDATION_TIMED_OUT"
	// @enum CertificateStatus
	CertificateStatusRevoked = "REVOKED"
	// @enum CertificateStatus
	CertificateStatusFailed = "FAILED"
)

const (
	// @enum KeyAlgorithm
	KeyAlgorithmRsa2048 = "RSA_2048"
	// @enum KeyAlgorithm
	KeyAlgorithmEcPrime256v1 = "EC_prime256v1"
)

const (
	// @enum RevocationReason
	RevocationReasonUnspecified = "UNSPECIFIED"
	// @enum RevocationReason
	RevocationReasonKeyCompromise = "KEY_COMPROMISE"
	// @enum RevocationReason
	RevocationReasonCaCompromise = "CA_COMPROMISE"
	// @enum RevocationReason
	RevocationReasonAffiliationChanged = "AFFILIATION_CHANGED"
	// @enum RevocationReason
	RevocationReasonSuperceded = "SUPERCEDED"
	// @enum RevocationReason
	RevocationReasonCessationOfOperation = "CESSATION_OF_OPERATION"
	// @enum RevocationReason
	RevocationReasonCertificateHold = "CERTIFICATE_HOLD"
	// @enum RevocationReason
	RevocationReasonRemoveFromCrl = "REMOVE_FROM_CRL"
	// @enum RevocationReason
	RevocationReasonPrivilegeWithdrawn = "PRIVILEGE_WITHDRAWN"
	// @enum RevocationReason
	RevocationReasonAACompromise = "A_A_COMPROMISE"
)
