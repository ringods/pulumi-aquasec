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
    public sealed class HostRuntimePolicyRegistryAccessMonitoring
    {
        public readonly bool? Enabled;
        public readonly ImmutableArray<string> ExceptionalMonitoredRegistryPaths;
        public readonly ImmutableArray<string> ExceptionalMonitoredRegistryProcesses;
        public readonly ImmutableArray<string> ExceptionalMonitoredRegistryUsers;
        public readonly bool? MonitoredRegistryAttributes;
        public readonly bool? MonitoredRegistryCreate;
        public readonly bool? MonitoredRegistryDelete;
        public readonly bool? MonitoredRegistryModify;
        public readonly ImmutableArray<string> MonitoredRegistryPaths;
        public readonly ImmutableArray<string> MonitoredRegistryProcesses;
        public readonly bool? MonitoredRegistryRead;
        public readonly ImmutableArray<string> MonitoredRegistryUsers;

        [OutputConstructor]
        private HostRuntimePolicyRegistryAccessMonitoring(
            bool? enabled,

            ImmutableArray<string> exceptionalMonitoredRegistryPaths,

            ImmutableArray<string> exceptionalMonitoredRegistryProcesses,

            ImmutableArray<string> exceptionalMonitoredRegistryUsers,

            bool? monitoredRegistryAttributes,

            bool? monitoredRegistryCreate,

            bool? monitoredRegistryDelete,

            bool? monitoredRegistryModify,

            ImmutableArray<string> monitoredRegistryPaths,

            ImmutableArray<string> monitoredRegistryProcesses,

            bool? monitoredRegistryRead,

            ImmutableArray<string> monitoredRegistryUsers)
        {
            Enabled = enabled;
            ExceptionalMonitoredRegistryPaths = exceptionalMonitoredRegistryPaths;
            ExceptionalMonitoredRegistryProcesses = exceptionalMonitoredRegistryProcesses;
            ExceptionalMonitoredRegistryUsers = exceptionalMonitoredRegistryUsers;
            MonitoredRegistryAttributes = monitoredRegistryAttributes;
            MonitoredRegistryCreate = monitoredRegistryCreate;
            MonitoredRegistryDelete = monitoredRegistryDelete;
            MonitoredRegistryModify = monitoredRegistryModify;
            MonitoredRegistryPaths = monitoredRegistryPaths;
            MonitoredRegistryProcesses = monitoredRegistryProcesses;
            MonitoredRegistryRead = monitoredRegistryRead;
            MonitoredRegistryUsers = monitoredRegistryUsers;
        }
    }
}
