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

    public sealed class VmwareAssurancePolicyKubernetesControlGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("avdId")]
        public Input<string>? AvdId { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ootb")]
        public Input<bool>? Ootb { get; set; }

        [Input("scriptId")]
        public Input<int>? ScriptId { get; set; }

        [Input("severity")]
        public Input<string>? Severity { get; set; }

        public VmwareAssurancePolicyKubernetesControlGetArgs()
        {
        }
        public static new VmwareAssurancePolicyKubernetesControlGetArgs Empty => new VmwareAssurancePolicyKubernetesControlGetArgs();
    }
}
