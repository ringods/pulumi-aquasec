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

    public sealed class GetApplicationScopeCategoryEntityScopeInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("variables")]
        private InputList<Inputs.GetApplicationScopeCategoryEntityScopeVariableInputArgs>? _variables;
        public InputList<Inputs.GetApplicationScopeCategoryEntityScopeVariableInputArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.GetApplicationScopeCategoryEntityScopeVariableInputArgs>());
            set => _variables = value;
        }

        public GetApplicationScopeCategoryEntityScopeInputArgs()
        {
        }
        public static new GetApplicationScopeCategoryEntityScopeInputArgs Empty => new GetApplicationScopeCategoryEntityScopeInputArgs();
    }
}
