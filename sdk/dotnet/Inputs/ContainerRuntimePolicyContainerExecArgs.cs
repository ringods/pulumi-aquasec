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

    public sealed class ContainerRuntimePolicyContainerExecArgs : global::Pulumi.ResourceArgs
    {
        [Input("blockContainerExec")]
        public Input<bool>? BlockContainerExec { get; set; }

        [Input("containerExecProcWhiteLists")]
        private InputList<string>? _containerExecProcWhiteLists;
        public InputList<string> ContainerExecProcWhiteLists
        {
            get => _containerExecProcWhiteLists ?? (_containerExecProcWhiteLists = new InputList<string>());
            set => _containerExecProcWhiteLists = value;
        }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("reverseShellIpWhiteLists")]
        private InputList<string>? _reverseShellIpWhiteLists;
        public InputList<string> ReverseShellIpWhiteLists
        {
            get => _reverseShellIpWhiteLists ?? (_reverseShellIpWhiteLists = new InputList<string>());
            set => _reverseShellIpWhiteLists = value;
        }

        public ContainerRuntimePolicyContainerExecArgs()
        {
        }
        public static new ContainerRuntimePolicyContainerExecArgs Empty => new ContainerRuntimePolicyContainerExecArgs();
    }
}
