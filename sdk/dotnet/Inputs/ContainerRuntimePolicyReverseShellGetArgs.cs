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

    public sealed class ContainerRuntimePolicyReverseShellGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("blockReverseShell")]
        public Input<bool>? BlockReverseShell { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("reverseShellIpWhiteLists")]
        private InputList<string>? _reverseShellIpWhiteLists;
        public InputList<string> ReverseShellIpWhiteLists
        {
            get => _reverseShellIpWhiteLists ?? (_reverseShellIpWhiteLists = new InputList<string>());
            set => _reverseShellIpWhiteLists = value;
        }

        [Input("reverseShellProcWhiteLists")]
        private InputList<string>? _reverseShellProcWhiteLists;
        public InputList<string> ReverseShellProcWhiteLists
        {
            get => _reverseShellProcWhiteLists ?? (_reverseShellProcWhiteLists = new InputList<string>());
            set => _reverseShellProcWhiteLists = value;
        }

        public ContainerRuntimePolicyReverseShellGetArgs()
        {
        }
        public static new ContainerRuntimePolicyReverseShellGetArgs Empty => new ContainerRuntimePolicyReverseShellGetArgs();
    }
}
