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

    public sealed class HostRuntimePolicyReadonlyRegistryGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("exceptionalReadonlyRegistryPaths")]
        private InputList<string>? _exceptionalReadonlyRegistryPaths;
        public InputList<string> ExceptionalReadonlyRegistryPaths
        {
            get => _exceptionalReadonlyRegistryPaths ?? (_exceptionalReadonlyRegistryPaths = new InputList<string>());
            set => _exceptionalReadonlyRegistryPaths = value;
        }

        [Input("exceptionalReadonlyRegistryProcesses")]
        private InputList<string>? _exceptionalReadonlyRegistryProcesses;
        public InputList<string> ExceptionalReadonlyRegistryProcesses
        {
            get => _exceptionalReadonlyRegistryProcesses ?? (_exceptionalReadonlyRegistryProcesses = new InputList<string>());
            set => _exceptionalReadonlyRegistryProcesses = value;
        }

        [Input("exceptionalReadonlyRegistryUsers")]
        private InputList<string>? _exceptionalReadonlyRegistryUsers;
        public InputList<string> ExceptionalReadonlyRegistryUsers
        {
            get => _exceptionalReadonlyRegistryUsers ?? (_exceptionalReadonlyRegistryUsers = new InputList<string>());
            set => _exceptionalReadonlyRegistryUsers = value;
        }

        [Input("readonlyRegistryPaths")]
        private InputList<string>? _readonlyRegistryPaths;
        public InputList<string> ReadonlyRegistryPaths
        {
            get => _readonlyRegistryPaths ?? (_readonlyRegistryPaths = new InputList<string>());
            set => _readonlyRegistryPaths = value;
        }

        [Input("readonlyRegistryProcesses")]
        private InputList<string>? _readonlyRegistryProcesses;
        public InputList<string> ReadonlyRegistryProcesses
        {
            get => _readonlyRegistryProcesses ?? (_readonlyRegistryProcesses = new InputList<string>());
            set => _readonlyRegistryProcesses = value;
        }

        [Input("readonlyRegistryUsers")]
        private InputList<string>? _readonlyRegistryUsers;
        public InputList<string> ReadonlyRegistryUsers
        {
            get => _readonlyRegistryUsers ?? (_readonlyRegistryUsers = new InputList<string>());
            set => _readonlyRegistryUsers = value;
        }

        public HostRuntimePolicyReadonlyRegistryGetArgs()
        {
        }
        public static new HostRuntimePolicyReadonlyRegistryGetArgs Empty => new HostRuntimePolicyReadonlyRegistryGetArgs();
    }
}
