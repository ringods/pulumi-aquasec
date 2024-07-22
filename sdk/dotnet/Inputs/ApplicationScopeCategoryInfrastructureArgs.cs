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

    public sealed class ApplicationScopeCategoryInfrastructureArgs : global::Pulumi.ResourceArgs
    {
        [Input("kubernetes")]
        private InputList<Inputs.ApplicationScopeCategoryInfrastructureKuberneteArgs>? _kubernetes;
        public InputList<Inputs.ApplicationScopeCategoryInfrastructureKuberneteArgs> Kubernetes
        {
            get => _kubernetes ?? (_kubernetes = new InputList<Inputs.ApplicationScopeCategoryInfrastructureKuberneteArgs>());
            set => _kubernetes = value;
        }

        [Input("os")]
        private InputList<Inputs.ApplicationScopeCategoryInfrastructureOArgs>? _os;
        public InputList<Inputs.ApplicationScopeCategoryInfrastructureOArgs> Os
        {
            get => _os ?? (_os = new InputList<Inputs.ApplicationScopeCategoryInfrastructureOArgs>());
            set => _os = value;
        }

        public ApplicationScopeCategoryInfrastructureArgs()
        {
        }
        public static new ApplicationScopeCategoryInfrastructureArgs Empty => new ApplicationScopeCategoryInfrastructureArgs();
    }
}
