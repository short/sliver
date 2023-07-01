// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the Amazon Web Services Systems Manager document (SSM document) and all
// managed node associations to the document. Before you delete the document, we
// recommend that you use DeleteAssociation to disassociate all managed nodes that
// are associated with the document.
func (c *Client) DeleteDocument(ctx context.Context, params *DeleteDocumentInput, optFns ...func(*Options)) (*DeleteDocumentOutput, error) {
	if params == nil {
		params = &DeleteDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDocument", params, optFns, c.addOperationDeleteDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDocumentInput struct {

	// The name of the document.
	//
	// This member is required.
	Name *string

	// The version of the document that you want to delete. If not provided, all
	// versions of the document are deleted.
	DocumentVersion *string

	// Some SSM document types require that you specify a Force flag before you can
	// delete the document. For example, you must specify a Force flag to delete a
	// document of type ApplicationConfigurationSchema . You can restrict access to the
	// Force flag in an Identity and Access Management (IAM) policy.
	Force bool

	// The version name of the document that you want to delete. If not provided, all
	// versions of the document are deleted.
	VersionName *string

	noSmithyDocumentSerde
}

type DeleteDocumentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDocument{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteDocumentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDocument(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteDocument(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DeleteDocument",
	}
}
