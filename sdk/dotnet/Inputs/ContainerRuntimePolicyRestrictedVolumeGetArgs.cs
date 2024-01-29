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

    public sealed class ContainerRuntimePolicyRestrictedVolumeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether restricted volumes are enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("volumes")]
        private InputList<string>? _volumes;

        /// <summary>
        /// List of restricted volumes.
        /// </summary>
        public InputList<string> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<string>());
            set => _volumes = value;
        }

        public ContainerRuntimePolicyRestrictedVolumeGetArgs()
        {
        }
        public static new ContainerRuntimePolicyRestrictedVolumeGetArgs Empty => new ContainerRuntimePolicyRestrictedVolumeGetArgs();
    }
}
