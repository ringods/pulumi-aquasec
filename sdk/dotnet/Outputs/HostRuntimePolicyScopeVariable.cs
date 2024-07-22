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
    public sealed class HostRuntimePolicyScopeVariable
    {
        /// <summary>
        /// Class of supported scope.
        /// </summary>
        public readonly string Attribute;
        /// <summary>
        /// Name assigned to the attribute.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Value assigned to the attribute.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private HostRuntimePolicyScopeVariable(
            string attribute,

            string? name,

            string value)
        {
            Attribute = attribute;
            Name = name;
            Value = value;
        }
    }
}
