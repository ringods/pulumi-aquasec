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
    public sealed class GetImageAssurancePolicyScopeVariableResult
    {
        public readonly string Attribute;
        public readonly string Name;
        public readonly string Value;

        [OutputConstructor]
        private GetImageAssurancePolicyScopeVariableResult(
            string attribute,

            string name,

            string value)
        {
            Attribute = attribute;
            Name = name;
            Value = value;
        }
    }
}
