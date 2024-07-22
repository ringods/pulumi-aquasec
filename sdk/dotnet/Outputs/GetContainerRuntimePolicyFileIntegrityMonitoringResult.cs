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
    public sealed class GetContainerRuntimePolicyFileIntegrityMonitoringResult
    {
        /// <summary>
        /// If true, file integrity monitoring is enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// List of paths to be excluded from monitoring.
        /// </summary>
        public readonly ImmutableArray<string> ExceptionalMonitoredFiles;
        /// <summary>
        /// List of processes to be excluded from monitoring.
        /// </summary>
        public readonly ImmutableArray<string> ExceptionalMonitoredFilesProcesses;
        /// <summary>
        /// List of users to be excluded from monitoring.
        /// </summary>
        public readonly ImmutableArray<string> ExceptionalMonitoredFilesUsers;
        /// <summary>
        /// List of paths to be monitored.
        /// </summary>
        public readonly ImmutableArray<string> MonitoredFiles;
        /// <summary>
        /// Whether to monitor file attribute operations.
        /// </summary>
        public readonly bool? MonitoredFilesAttributes;
        /// <summary>
        /// Whether to monitor file create operations.
        /// </summary>
        public readonly bool? MonitoredFilesCreate;
        /// <summary>
        /// Whether to monitor file delete operations.
        /// </summary>
        public readonly bool? MonitoredFilesDelete;
        /// <summary>
        /// Whether to monitor file modify operations.
        /// </summary>
        public readonly bool? MonitoredFilesModify;
        /// <summary>
        /// List of processes associated with monitored files.
        /// </summary>
        public readonly ImmutableArray<string> MonitoredFilesProcesses;
        /// <summary>
        /// Whether to monitor file read operations.
        /// </summary>
        public readonly bool? MonitoredFilesRead;
        /// <summary>
        /// List of users associated with monitored files.
        /// </summary>
        public readonly ImmutableArray<string> MonitoredFilesUsers;

        [OutputConstructor]
        private GetContainerRuntimePolicyFileIntegrityMonitoringResult(
            bool? enabled,

            ImmutableArray<string> exceptionalMonitoredFiles,

            ImmutableArray<string> exceptionalMonitoredFilesProcesses,

            ImmutableArray<string> exceptionalMonitoredFilesUsers,

            ImmutableArray<string> monitoredFiles,

            bool? monitoredFilesAttributes,

            bool? monitoredFilesCreate,

            bool? monitoredFilesDelete,

            bool? monitoredFilesModify,

            ImmutableArray<string> monitoredFilesProcesses,

            bool? monitoredFilesRead,

            ImmutableArray<string> monitoredFilesUsers)
        {
            Enabled = enabled;
            ExceptionalMonitoredFiles = exceptionalMonitoredFiles;
            ExceptionalMonitoredFilesProcesses = exceptionalMonitoredFilesProcesses;
            ExceptionalMonitoredFilesUsers = exceptionalMonitoredFilesUsers;
            MonitoredFiles = monitoredFiles;
            MonitoredFilesAttributes = monitoredFilesAttributes;
            MonitoredFilesCreate = monitoredFilesCreate;
            MonitoredFilesDelete = monitoredFilesDelete;
            MonitoredFilesModify = monitoredFilesModify;
            MonitoredFilesProcesses = monitoredFilesProcesses;
            MonitoredFilesRead = monitoredFilesRead;
            MonitoredFilesUsers = monitoredFilesUsers;
        }
    }
}
