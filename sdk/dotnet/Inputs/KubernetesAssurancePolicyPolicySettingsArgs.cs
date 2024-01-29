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

    public sealed class KubernetesAssurancePolicyPolicySettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("enforce")]
        public Input<bool>? Enforce { get; set; }

        [Input("isAuditChecked")]
        public Input<bool>? IsAuditChecked { get; set; }

        [Input("warn")]
        public Input<bool>? Warn { get; set; }

        [Input("warningMessage")]
        public Input<string>? WarningMessage { get; set; }

        public KubernetesAssurancePolicyPolicySettingsArgs()
        {
        }
        public static new KubernetesAssurancePolicyPolicySettingsArgs Empty => new KubernetesAssurancePolicyPolicySettingsArgs();
    }
}
