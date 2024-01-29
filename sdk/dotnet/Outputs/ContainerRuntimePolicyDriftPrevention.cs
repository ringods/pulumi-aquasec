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
    public sealed class ContainerRuntimePolicyDriftPrevention
    {
        /// <summary>
        /// Whether drift prevention is enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Whether to lockdown execution drift.
        /// </summary>
        public readonly bool? ExecLockdown;
        /// <summary>
        /// List of items in the execution lockdown white list.
        /// </summary>
        public readonly ImmutableArray<string> ExecLockdownWhiteLists;
        /// <summary>
        /// Whether to lockdown image drift.
        /// </summary>
        public readonly bool? ImageLockdown;

        [OutputConstructor]
        private ContainerRuntimePolicyDriftPrevention(
            bool? enabled,

            bool? execLockdown,

            ImmutableArray<string> execLockdownWhiteLists,

            bool? imageLockdown)
        {
            Enabled = enabled;
            ExecLockdown = execLockdown;
            ExecLockdownWhiteLists = execLockdownWhiteLists;
            ImageLockdown = imageLockdown;
        }
    }
}
