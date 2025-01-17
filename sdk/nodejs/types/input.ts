// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";

import * as aws from "@pulumi/aws";

/**
 * A route that that APIGateway should accept and forward to some type of destination. All routes
 * have an incoming path that they match against.  However, destinations are determined by the kind
 * of the route.
 */
export interface RouteArgs {
    /**
     * The `content-type` to serve the file as.  Only valid when `localPath` points to a file.  If
     * `localPath` points to a directory, the content types for all files will be inferred.
     */
    contentType?: pulumi.Input<string>;
    /**
     * A raw Swagger object to include verbatim in the integration for this path.
     */
    data?: any;
    /**
     * A Lambda function which will handle the route for the given path and method.
     */
    eventHandler?: pulumi.Input<aws.lambda.Function>;
    /**
     * By default a `localPath` hosting static content will also serve 'index.html' in response to a request on a directory.
     * To disable this pass `false` or supply a new index document name.
     */
    index?: pulumi.Input<string | boolean>;
    /**
     * The local path on disk to create static S3 resources for.  Files will be uploaded into S3
     * objects, and directories will be recursively walked into.
     */
    localPath?: pulumi.Input<string>;
    /**
     * The REST method of the route to match.  Only valid with `eventHandler` or `data` routes.
     */
    method?: pulumi.Input<enums.Method>;
    /**
     * The path on the API that will serve this route.  If not prefixed with `/`,
     * then a `/` will be added automatically to the beginning.
     */
    path: pulumi.Input<string>;
    /**
     * The target for an integration route (see https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-integration-types.html).
     */
    target?: pulumi.Input<inputs.TargetArgs>;
}

export interface SwaggerGatewayResponse {
    responseParameters?: {[key: string]: string};
    responseTemplates?: {[key: string]: string};
    statusCode?: number;
}

export interface TargetArgs {
    /**
     * The (id) of the VpcLink used for the integration when connectionType=VPC_LINK and undefined,
     * otherwise.
     */
    connectionId?: pulumi.Input<string>;
    /**
     * The type of the network connection to the integration endpoint. The valid value is `INTERNET`
     * for connections through the public routable internet or `VPC_LINK` for private connections
     * between API Gateway and a network load balancer in a VPC. The default value is `INTERNET`.
     */
    connectionType?: pulumi.Input<enums.IntegrationConnectionType>;
    /**
     * Specifies the integration's HTTP method type.  Currently, the only supported type is 'ANY'.
     */
    httpMethod?: pulumi.Input<"ANY">;
    /**
     * Specifies how the method request body of an unmapped content type will be passed through the
     * integration request to the back end without transformation.
     *
     * The valid value is one of the following:
     *
     * * `WHEN_NO_MATCH`: passes the method request body through the integration request to the back end
     * without transformation when the method request content type does not match any content type
     * associated with the mapping templates defined in the integration request.
     *
     * * `WHEN_NO_TEMPLATES`: passes the method request body through the integration request to the back
     * end without transformation when no mapping template is defined in the integration request. If
     * a template is defined when this option is selected, the method request of an unmapped
     * content-type will be rejected with an HTTP 415 Unsupported Media Type response.
     *
     * * `NEVER`: rejects the method request with an HTTP 415 Unsupported Media Type response when
     * either the method request content type does not match any content type associated with the
     * mapping templates defined in the integration request or no mapping template is defined in the
     * integration request.
     *
     * Defaults to `WHEN_NO_MATCH` if unspecified.
     */
    passthroughBehaviour?: pulumi.Input<enums.IntegrationPassthroughBehavior>;
    /**
     * Specifies an API method integration type. The valid value is one of the following:
     *
     * * `aws`: for integrating the API method request with an AWS service action, including the Lambda
     * function-invoking action. With the Lambda function-invoking action, this is referred to as
     * the Lambda custom integration. With any other AWS service action, this is known as AWS
     * integration.
     *
     * * `aws_proxy`: for integrating the API method request with the Lambda function-invoking action
     * with the client request passed through as-is. This integration is also referred to as the
     * Lambda proxy integration.
     *
     * * `http`: for integrating the API method request with an HTTP endpoint, including a private HTTP
     * endpoint within a VPC. This integration is also referred to as the HTTP custom integration.
     *
     * * `http_proxy`: for integrating the API method request with an HTTP endpoint, including a private
     * HTTP endpoint within a VPC, with the client request passed through as-is. This is also
     * referred to as the HTTP proxy integration.
     *
     * * `mock`: for integrating the API method request with API Gateway as a "loop-back" endpoint
     * without invoking any backend.
     */
    type: pulumi.Input<enums.IntegrationType>;
    /**
     * Specifies Uniform Resource Identifier (URI) of the integration endpoint.
     *
     * For HTTP or HTTP_PROXY integrations, the URI must be a fully formed, encoded HTTP(S) URL
     * according to the RFC-3986 specification, for either standard integration, where
     * connectionType is not VPC_LINK, or private integration, where connectionType is VPC_LINK. For
     * a private HTTP integration, the URI is not used for routing.
     *
     * For AWS or AWS_PROXY integrations, the URI is of the form
     * arn:aws:apigateway:{region}:{subdomain.service|service}:path|action/{service_api}. Here,
     * {Region} is the API Gateway region (e.g., us-east-1); {service} is the name of the integrated
     * AWS service (e.g., s3); and {subdomain} is a designated subdomain supported by certain AWS
     * service for fast host-name lookup. action can be used for an AWS service action-based API,
     * using an Action={name}&{p1}={v1}&p2={v2}... query string. The ensuing {service_api} refers to
     * a supported action {name} plus any required input parameters. Alternatively, path can be used
     * for an AWS service path-based API. The ensuing service_api refers to the path to an AWS
     * service resource, including the region of the integrated AWS service, if applicable. For
     * example, for integration with the S3 API of GetObject, the uri can be either
     * arn:aws:apigateway:us-west-2:s3:action/GetObject&Bucket={bucket}&Key={key} or
     * arn:aws:apigateway:us-west-2:s3:path/{bucket}/{key}.
     */
    uri: pulumi.Input<string>;
}
