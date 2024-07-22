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

    public sealed class FunctionRuntimePolicyFileBlockArgs : global::Pulumi.ResourceArgs
    {
        [Input("blockFilesProcesses")]
        private InputList<string>? _blockFilesProcesses;
        public InputList<string> BlockFilesProcesses
        {
            get => _blockFilesProcesses ?? (_blockFilesProcesses = new InputList<string>());
            set => _blockFilesProcesses = value;
        }

        [Input("blockFilesUsers")]
        private InputList<string>? _blockFilesUsers;
        public InputList<string> BlockFilesUsers
        {
            get => _blockFilesUsers ?? (_blockFilesUsers = new InputList<string>());
            set => _blockFilesUsers = value;
        }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("exceptionalBlockFiles")]
        private InputList<string>? _exceptionalBlockFiles;
        public InputList<string> ExceptionalBlockFiles
        {
            get => _exceptionalBlockFiles ?? (_exceptionalBlockFiles = new InputList<string>());
            set => _exceptionalBlockFiles = value;
        }

        [Input("exceptionalBlockFilesProcesses")]
        private InputList<string>? _exceptionalBlockFilesProcesses;
        public InputList<string> ExceptionalBlockFilesProcesses
        {
            get => _exceptionalBlockFilesProcesses ?? (_exceptionalBlockFilesProcesses = new InputList<string>());
            set => _exceptionalBlockFilesProcesses = value;
        }

        [Input("exceptionalBlockFilesUsers")]
        private InputList<string>? _exceptionalBlockFilesUsers;
        public InputList<string> ExceptionalBlockFilesUsers
        {
            get => _exceptionalBlockFilesUsers ?? (_exceptionalBlockFilesUsers = new InputList<string>());
            set => _exceptionalBlockFilesUsers = value;
        }

        [Input("filenameBlockLists")]
        private InputList<string>? _filenameBlockLists;
        public InputList<string> FilenameBlockLists
        {
            get => _filenameBlockLists ?? (_filenameBlockLists = new InputList<string>());
            set => _filenameBlockLists = value;
        }

        public FunctionRuntimePolicyFileBlockArgs()
        {
        }
        public static new FunctionRuntimePolicyFileBlockArgs Empty => new FunctionRuntimePolicyFileBlockArgs();
    }
}
