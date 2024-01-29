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
    public sealed class GetContainerRuntimePolicyLimitContainerPrivilegeResult
    {
        /// <summary>
        /// Whether to block adding capabilities.
        /// </summary>
        public readonly bool? BlockAddCapabilities;
        /// <summary>
        /// Whether container privilege limitations are enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Whether to limit IPC-related capabilities.
        /// </summary>
        public readonly bool? Ipcmode;
        /// <summary>
        /// Whether to limit network-related capabilities.
        /// </summary>
        public readonly bool? Netmode;
        /// <summary>
        /// Whether to limit process-related capabilities.
        /// </summary>
        public readonly bool? Pidmode;
        /// <summary>
        /// Whether to prevent low port binding.
        /// </summary>
        public readonly bool? PreventLowPortBinding;
        /// <summary>
        /// Whether to prevent the use of the root user.
        /// </summary>
        public readonly bool? PreventRootUser;
        /// <summary>
        /// Whether the container is run in privileged mode.
        /// </summary>
        public readonly bool? Privileged;
        /// <summary>
        /// Whether to use the host user.
        /// </summary>
        public readonly bool? UseHostUser;
        /// <summary>
        /// Whether to limit user-related capabilities.
        /// </summary>
        public readonly bool? Usermode;
        /// <summary>
        /// Whether to limit UTS-related capabilities.
        /// </summary>
        public readonly bool? Utsmode;

        [OutputConstructor]
        private GetContainerRuntimePolicyLimitContainerPrivilegeResult(
            bool? blockAddCapabilities,

            bool? enabled,

            bool? ipcmode,

            bool? netmode,

            bool? pidmode,

            bool? preventLowPortBinding,

            bool? preventRootUser,

            bool? privileged,

            bool? useHostUser,

            bool? usermode,

            bool? utsmode)
        {
            BlockAddCapabilities = blockAddCapabilities;
            Enabled = enabled;
            Ipcmode = ipcmode;
            Netmode = netmode;
            Pidmode = pidmode;
            PreventLowPortBinding = preventLowPortBinding;
            PreventRootUser = preventRootUser;
            Privileged = privileged;
            UseHostUser = useHostUser;
            Usermode = usermode;
            Utsmode = utsmode;
        }
    }
}
