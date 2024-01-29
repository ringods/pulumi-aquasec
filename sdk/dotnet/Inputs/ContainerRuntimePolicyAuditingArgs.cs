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

    public sealed class ContainerRuntimePolicyAuditingArgs : global::Pulumi.ResourceArgs
    {
        [Input("auditAllNetwork")]
        public Input<bool>? AuditAllNetwork { get; set; }

        [Input("auditAllProcesses")]
        public Input<bool>? AuditAllProcesses { get; set; }

        [Input("auditFailedLogin")]
        public Input<bool>? AuditFailedLogin { get; set; }

        [Input("auditOsUserActivity")]
        public Input<bool>? AuditOsUserActivity { get; set; }

        [Input("auditProcessCmdline")]
        public Input<bool>? AuditProcessCmdline { get; set; }

        [Input("auditSuccessLogin")]
        public Input<bool>? AuditSuccessLogin { get; set; }

        [Input("auditUserAccountManagement")]
        public Input<bool>? AuditUserAccountManagement { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ContainerRuntimePolicyAuditingArgs()
        {
        }
        public static new ContainerRuntimePolicyAuditingArgs Empty => new ContainerRuntimePolicyAuditingArgs();
    }
}
