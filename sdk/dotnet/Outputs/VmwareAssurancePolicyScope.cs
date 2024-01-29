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
    public sealed class VmwareAssurancePolicyScope
    {
        public readonly string? Expression;
        public readonly ImmutableArray<Outputs.VmwareAssurancePolicyScopeVariable> Variables;

        [OutputConstructor]
        private VmwareAssurancePolicyScope(
            string? expression,

            ImmutableArray<Outputs.VmwareAssurancePolicyScopeVariable> variables)
        {
            Expression = expression;
            Variables = variables;
        }
    }
}
