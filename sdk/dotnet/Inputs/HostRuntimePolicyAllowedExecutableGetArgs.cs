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

    public sealed class HostRuntimePolicyAllowedExecutableGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowExecutables")]
        private InputList<string>? _allowExecutables;

        /// <summary>
        /// List of allowed executables.
        /// </summary>
        public InputList<string> AllowExecutables
        {
            get => _allowExecutables ?? (_allowExecutables = new InputList<string>());
            set => _allowExecutables = value;
        }

        [Input("allowRootExecutables")]
        private InputList<string>? _allowRootExecutables;

        /// <summary>
        /// List of allowed root executables.
        /// </summary>
        public InputList<string> AllowRootExecutables
        {
            get => _allowRootExecutables ?? (_allowRootExecutables = new InputList<string>());
            set => _allowRootExecutables = value;
        }

        /// <summary>
        /// Whether allowed executables configuration is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Whether to treat executables separately.
        /// </summary>
        [Input("separateExecutables")]
        public Input<bool>? SeparateExecutables { get; set; }

        public HostRuntimePolicyAllowedExecutableGetArgs()
        {
        }
        public static new HostRuntimePolicyAllowedExecutableGetArgs Empty => new HostRuntimePolicyAllowedExecutableGetArgs();
    }
}
