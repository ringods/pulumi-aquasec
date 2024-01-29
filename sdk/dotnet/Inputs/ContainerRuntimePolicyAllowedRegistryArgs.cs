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

    public sealed class ContainerRuntimePolicyAllowedRegistryArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedRegistries")]
        private InputList<string>? _allowedRegistries;

        /// <summary>
        /// List of allowed registries.
        /// </summary>
        public InputList<string> AllowedRegistries
        {
            get => _allowedRegistries ?? (_allowedRegistries = new InputList<string>());
            set => _allowedRegistries = value;
        }

        /// <summary>
        /// Whether allowed registries are enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ContainerRuntimePolicyAllowedRegistryArgs()
        {
        }
        public static new ContainerRuntimePolicyAllowedRegistryArgs Empty => new ContainerRuntimePolicyAllowedRegistryArgs();
    }
}
