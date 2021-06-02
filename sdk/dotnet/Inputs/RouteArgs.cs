// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Apigateway.Inputs
{

    /// <summary>
    /// A route that that APIGateway should accept and forward to some type of destination. All routes
    /// have an incoming path that they match against.  However, destinations are determined by the kind
    /// of the route.
    /// </summary>
    public sealed class RouteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The `content-type` to serve the file as.  Only valid when `localPath` points to a file.  If
        /// `localPath` points to a directory, the content types for all files will be inferred.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// A raw Swagger object to include verbatim in the integration for this path.
        /// </summary>
        [Input("data")]
        public Input<object>? Data { get; set; }

        /// <summary>
        /// A Lambda function which will handle the route for the given path and method.
        /// </summary>
        [Input("eventHandler")]
        public Input<Pulumi.Aws.Lambda.Function>? EventHandler { get; set; }

        /// <summary>
        /// By default a `localPath` hosting static content will also serve 'index.html' in response to a request on a directory.
        /// To disable this pass `false` or supply a new index document name.
        /// </summary>
        [Input("index")]
        public InputUnion<string, bool>? Index { get; set; }

        /// <summary>
        /// The local path on disk to create static S3 resources for.  Files will be uploaded into S3
        /// objects, and directories will be recursively walked into.
        /// </summary>
        [Input("localPath")]
        public Input<string>? LocalPath { get; set; }

        /// <summary>
        /// The REST method of the route to match.  Only valid with `eventHandler` or `data` routes.
        /// </summary>
        [Input("method")]
        public Input<Pulumi.Apigateway.Method>? Method { get; set; }

        /// <summary>
        /// The path on the API that will serve this route.  If not prefixed with `/`,
        /// then a `/` will be added automatically to the beginning.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The target for an integration route (see https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-integration-types.html).
        /// </summary>
        [Input("target")]
        public Input<Inputs.TargetArgs>? Target { get; set; }

        public RouteArgs()
        {
        }
    }
}
