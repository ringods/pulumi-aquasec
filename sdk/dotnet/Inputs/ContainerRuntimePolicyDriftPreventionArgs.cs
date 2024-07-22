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

    public sealed class ContainerRuntimePolicyDriftPreventionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether drift prevention is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Whether to lockdown execution drift.
        /// </summary>
        [Input("execLockdown")]
        public Input<bool>? ExecLockdown { get; set; }

        [Input("execLockdownWhiteLists")]
        private InputList<string>? _execLockdownWhiteLists;

        /// <summary>
        /// List of items in the execution lockdown white list.
        /// </summary>
        public InputList<string> ExecLockdownWhiteLists
        {
            get => _execLockdownWhiteLists ?? (_execLockdownWhiteLists = new InputList<string>());
            set => _execLockdownWhiteLists = value;
        }

        /// <summary>
        /// Whether to lockdown image drift.
        /// </summary>
        [Input("imageLockdown")]
        public Input<bool>? ImageLockdown { get; set; }

        public ContainerRuntimePolicyDriftPreventionArgs()
        {
        }
        public static new ContainerRuntimePolicyDriftPreventionArgs Empty => new ContainerRuntimePolicyDriftPreventionArgs();
    }
}
