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

    public sealed class GetContainerRuntimePolicyReadonlyFilesInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("exceptionalReadonlyFiles")]
        private InputList<string>? _exceptionalReadonlyFiles;
        public InputList<string> ExceptionalReadonlyFiles
        {
            get => _exceptionalReadonlyFiles ?? (_exceptionalReadonlyFiles = new InputList<string>());
            set => _exceptionalReadonlyFiles = value;
        }

        [Input("exceptionalReadonlyFilesProcesses")]
        private InputList<string>? _exceptionalReadonlyFilesProcesses;
        public InputList<string> ExceptionalReadonlyFilesProcesses
        {
            get => _exceptionalReadonlyFilesProcesses ?? (_exceptionalReadonlyFilesProcesses = new InputList<string>());
            set => _exceptionalReadonlyFilesProcesses = value;
        }

        [Input("exceptionalReadonlyFilesUsers")]
        private InputList<string>? _exceptionalReadonlyFilesUsers;
        public InputList<string> ExceptionalReadonlyFilesUsers
        {
            get => _exceptionalReadonlyFilesUsers ?? (_exceptionalReadonlyFilesUsers = new InputList<string>());
            set => _exceptionalReadonlyFilesUsers = value;
        }

        [Input("readonlyFiles")]
        private InputList<string>? _readonlyFiles;
        public InputList<string> ReadonlyFiles
        {
            get => _readonlyFiles ?? (_readonlyFiles = new InputList<string>());
            set => _readonlyFiles = value;
        }

        [Input("readonlyFilesProcesses")]
        private InputList<string>? _readonlyFilesProcesses;
        public InputList<string> ReadonlyFilesProcesses
        {
            get => _readonlyFilesProcesses ?? (_readonlyFilesProcesses = new InputList<string>());
            set => _readonlyFilesProcesses = value;
        }

        [Input("readonlyFilesUsers")]
        private InputList<string>? _readonlyFilesUsers;
        public InputList<string> ReadonlyFilesUsers
        {
            get => _readonlyFilesUsers ?? (_readonlyFilesUsers = new InputList<string>());
            set => _readonlyFilesUsers = value;
        }

        public GetContainerRuntimePolicyReadonlyFilesInputArgs()
        {
        }
        public static new GetContainerRuntimePolicyReadonlyFilesInputArgs Empty => new GetContainerRuntimePolicyReadonlyFilesInputArgs();
    }
}
