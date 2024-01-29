// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Aquasec.Outputs
{

    [OutputType]
    public sealed class HostRuntimePolicyRestrictedVolume
    {
        /// <summary>
        /// Whether restricted volumes are enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// List of restricted volumes.
        /// </summary>
        public readonly ImmutableArray<string> Volumes;

        [OutputConstructor]
        private HostRuntimePolicyRestrictedVolume(
            bool? enabled,

            ImmutableArray<string> volumes)
        {
            Enabled = enabled;
            Volumes = volumes;
        }
    }
}
