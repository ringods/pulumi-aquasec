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
    public sealed class FunctionRuntimePolicyBypassScope
    {
        /// <summary>
        /// Whether bypassing the scope is enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Scope configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.FunctionRuntimePolicyBypassScopeScope> Scopes;

        [OutputConstructor]
        private FunctionRuntimePolicyBypassScope(
            bool? enabled,

            ImmutableArray<Outputs.FunctionRuntimePolicyBypassScopeScope> scopes)
        {
            Enabled = enabled;
            Scopes = scopes;
        }
    }
}
