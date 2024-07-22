// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Aquasec.Inputs
{

    public sealed class FirewallPolicyOutboundNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the specified resources are allowed to receive data or requests.
        /// </summary>
        [Input("allow", required: true)]
        public Input<bool> Allow { get; set; } = null!;

        /// <summary>
        /// Range of ports affected by firewall.
        /// </summary>
        [Input("portRange", required: true)]
        public Input<string> PortRange { get; set; } = null!;

        /// <summary>
        /// Information of the resource.
        /// </summary>
        [Input("resource")]
        public Input<string>? Resource { get; set; }

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        public FirewallPolicyOutboundNetworkGetArgs()
        {
        }
        public static new FirewallPolicyOutboundNetworkGetArgs Empty => new FirewallPolicyOutboundNetworkGetArgs();
    }
}
