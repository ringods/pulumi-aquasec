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
    public sealed class HostRuntimePolicyContainerExec
    {
        public readonly bool? BlockContainerExec;
        public readonly ImmutableArray<string> ContainerExecProcWhiteLists;
        public readonly bool? Enabled;
        public readonly ImmutableArray<string> ReverseShellIpWhiteLists;

        [OutputConstructor]
        private HostRuntimePolicyContainerExec(
            bool? blockContainerExec,

            ImmutableArray<string> containerExecProcWhiteLists,

            bool? enabled,

            ImmutableArray<string> reverseShellIpWhiteLists)
        {
            BlockContainerExec = blockContainerExec;
            ContainerExecProcWhiteLists = containerExecProcWhiteLists;
            Enabled = enabled;
            ReverseShellIpWhiteLists = reverseShellIpWhiteLists;
        }
    }
}
