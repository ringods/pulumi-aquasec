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
    public sealed class GetAquaLabelsAquaLabelResult
    {
        public readonly string Author;
        public readonly string Created;
        public readonly string Description;
        public readonly string Name;

        [OutputConstructor]
        private GetAquaLabelsAquaLabelResult(
            string author,

            string created,

            string description,

            string name)
        {
            Author = author;
            Created = created;
            Description = description;
            Name = name;
        }
    }
}
