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

    public sealed class HostRuntimePolicySystemIntegrityProtectionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("auditSystemtimeChange")]
        public Input<bool>? AuditSystemtimeChange { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("monitorAuditLogIntegrity")]
        public Input<bool>? MonitorAuditLogIntegrity { get; set; }

        [Input("windowsServicesMonitoring")]
        public Input<bool>? WindowsServicesMonitoring { get; set; }

        public HostRuntimePolicySystemIntegrityProtectionGetArgs()
        {
        }
        public static new HostRuntimePolicySystemIntegrityProtectionGetArgs Empty => new HostRuntimePolicySystemIntegrityProtectionGetArgs();
    }
}
