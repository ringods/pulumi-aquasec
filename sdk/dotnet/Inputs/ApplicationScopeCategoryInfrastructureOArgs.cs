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

    public sealed class ApplicationScopeCategoryInfrastructureOArgs : global::Pulumi.ResourceArgs
    {
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        [Input("variables")]
        private InputList<Inputs.ApplicationScopeCategoryInfrastructureOVariableArgs>? _variables;
        public InputList<Inputs.ApplicationScopeCategoryInfrastructureOVariableArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.ApplicationScopeCategoryInfrastructureOVariableArgs>());
            set => _variables = value;
        }

        public ApplicationScopeCategoryInfrastructureOArgs()
        {
        }
        public static new ApplicationScopeCategoryInfrastructureOArgs Empty => new ApplicationScopeCategoryInfrastructureOArgs();
    }
}
