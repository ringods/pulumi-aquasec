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
    public sealed class GetContainerRuntimePolicyFileBlockResult
    {
        public readonly ImmutableArray<string> BlockFilesProcesses;
        public readonly ImmutableArray<string> BlockFilesUsers;
        public readonly bool? Enabled;
        public readonly ImmutableArray<string> ExceptionalBlockFiles;
        public readonly ImmutableArray<string> ExceptionalBlockFilesProcesses;
        public readonly ImmutableArray<string> ExceptionalBlockFilesUsers;
        public readonly ImmutableArray<string> FilenameBlockLists;

        [OutputConstructor]
        private GetContainerRuntimePolicyFileBlockResult(
            ImmutableArray<string> blockFilesProcesses,

            ImmutableArray<string> blockFilesUsers,

            bool? enabled,

            ImmutableArray<string> exceptionalBlockFiles,

            ImmutableArray<string> exceptionalBlockFilesProcesses,

            ImmutableArray<string> exceptionalBlockFilesUsers,

            ImmutableArray<string> filenameBlockLists)
        {
            BlockFilesProcesses = blockFilesProcesses;
            BlockFilesUsers = blockFilesUsers;
            Enabled = enabled;
            ExceptionalBlockFiles = exceptionalBlockFiles;
            ExceptionalBlockFilesProcesses = exceptionalBlockFilesProcesses;
            ExceptionalBlockFilesUsers = exceptionalBlockFilesUsers;
            FilenameBlockLists = filenameBlockLists;
        }
    }
}
