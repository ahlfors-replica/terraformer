// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteCertificateInput struct {
	_ struct{} `type:"structure"`

	// String that contains the ARN of the ACM certificate to be deleted. This must
	// be of the form:
	//
	// arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	//
	// CertificateArn is a required field
	CertificateArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCertificateInput"}

	if s.CertificateArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateArn"))
	}
	if s.CertificateArn != nil && len(*s.CertificateArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteCertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCertificate = "DeleteCertificate"

// DeleteCertificateRequest returns a request value for making API operation for
// AWS Certificate Manager.
//
// Deletes a certificate and its associated private key. If this action succeeds,
// the certificate no longer appears in the list that can be displayed by calling
// the ListCertificates action or be retrieved by calling the GetCertificate
// action. The certificate will not be available for use by AWS services integrated
// with ACM.
//
// You cannot delete an ACM certificate that is being used by another AWS service.
// To delete a certificate that is in use, the certificate association must
// first be removed.
//
//    // Example sending a request using DeleteCertificateRequest.
//    req := client.DeleteCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/DeleteCertificate
func (c *Client) DeleteCertificateRequest(input *DeleteCertificateInput) DeleteCertificateRequest {
	op := &aws.Operation{
		Name:       opDeleteCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCertificateInput{}
	}

	req := c.newRequest(op, input, &DeleteCertificateOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteCertificateRequest{Request: req, Input: input, Copy: c.DeleteCertificateRequest}
}

// DeleteCertificateRequest is the request type for the
// DeleteCertificate API operation.
type DeleteCertificateRequest struct {
	*aws.Request
	Input *DeleteCertificateInput
	Copy  func(*DeleteCertificateInput) DeleteCertificateRequest
}

// Send marshals and sends the DeleteCertificate API request.
func (r DeleteCertificateRequest) Send(ctx context.Context) (*DeleteCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCertificateResponse{
		DeleteCertificateOutput: r.Request.Data.(*DeleteCertificateOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCertificateResponse is the response type for the
// DeleteCertificate API operation.
type DeleteCertificateResponse struct {
	*DeleteCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCertificate request.
func (r *DeleteCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}