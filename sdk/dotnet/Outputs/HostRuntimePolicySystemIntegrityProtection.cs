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
    public sealed class HostRuntimePolicySystemIntegrityProtection
    {
        public readonly bool? AuditSystemtimeChange;
        public readonly bool? Enabled;
        public readonly bool? MonitorAuditLogIntegrity;
        public readonly bool? WindowsServicesMonitoring;

        [OutputConstructor]
        private HostRuntimePolicySystemIntegrityProtection(
            bool? auditSystemtimeChange,

            bool? enabled,

            bool? monitorAuditLogIntegrity,

            bool? windowsServicesMonitoring)
        {
            AuditSystemtimeChange = auditSystemtimeChange;
            Enabled = enabled;
            MonitorAuditLogIntegrity = monitorAuditLogIntegrity;
            WindowsServicesMonitoring = windowsServicesMonitoring;
        }
    }
}
