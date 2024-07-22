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

    public sealed class FunctionRuntimePolicyScopeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Scope expression.
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("variables", required: true)]
        private InputList<Inputs.FunctionRuntimePolicyScopeVariableArgs>? _variables;

        /// <summary>
        /// List of variables in the scope.
        /// </summary>
        public InputList<Inputs.FunctionRuntimePolicyScopeVariableArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.FunctionRuntimePolicyScopeVariableArgs>());
            set => _variables = value;
        }

        public FunctionRuntimePolicyScopeArgs()
        {
        }
        public static new FunctionRuntimePolicyScopeArgs Empty => new FunctionRuntimePolicyScopeArgs();
    }
}
