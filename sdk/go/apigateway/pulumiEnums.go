// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type APIKeySource pulumi.String

const (
	APIKeySourceHEADER     = APIKeySource("HEADER")
	APIKeySourceAUTHORIZER = APIKeySource("AUTHORIZER")
)

func (APIKeySource) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e APIKeySource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e APIKeySource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e APIKeySource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e APIKeySource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IntegrationConnectionType pulumi.String

const (
	IntegrationConnectionTypeINTERNET  = IntegrationConnectionType("INTERNET")
	IntegrationConnectionType_VPC_LINK = IntegrationConnectionType("VPC_LINK")
)

func (IntegrationConnectionType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e IntegrationConnectionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationConnectionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationConnectionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IntegrationConnectionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IntegrationPassthroughBehavior pulumi.String

const (
	IntegrationPassthroughBehavior_When_no_match     = IntegrationPassthroughBehavior("when_no_match")
	IntegrationPassthroughBehavior_When_no_templates = IntegrationPassthroughBehavior("when_no_templates")
	IntegrationPassthroughBehaviorNever              = IntegrationPassthroughBehavior("never")
)

func (IntegrationPassthroughBehavior) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e IntegrationPassthroughBehavior) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationPassthroughBehavior) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationPassthroughBehavior) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IntegrationPassthroughBehavior) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IntegrationType pulumi.String

const (
	IntegrationTypeAws         = IntegrationType("aws")
	IntegrationType_Aws_proxy  = IntegrationType("aws_proxy")
	IntegrationTypeHttp        = IntegrationType("http")
	IntegrationType_Http_proxy = IntegrationType("http_proxy")
	IntegrationTypeMock        = IntegrationType("mock")
)

func (IntegrationType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e IntegrationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IntegrationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type Method pulumi.String

const (
	MethodANY     = Method("ANY")
	MethodGET     = Method("GET")
	MethodPUT     = Method("PUT")
	MethodPOST    = Method("POST")
	MethodDELETE  = Method("DELETE")
	MethodPATCH   = Method("PATCH")
	MethodOPTIONS = Method("OPTIONS")
	MethodHEAD    = Method("HEAD")
)

func (Method) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e Method) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Method) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Method) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Method) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RequestValidator pulumi.String

const (
	RequestValidatorALL          = RequestValidator("ALL")
	RequestValidator_PARAMS_ONLY = RequestValidator("PARAMS_ONLY")
	RequestValidator_BODY_ONLY   = RequestValidator("BODY_ONLY")
)

func (RequestValidator) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e RequestValidator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RequestValidator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RequestValidator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RequestValidator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
