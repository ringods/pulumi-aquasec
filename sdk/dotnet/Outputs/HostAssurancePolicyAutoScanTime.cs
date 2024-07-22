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
    public sealed class HostAssurancePolicyAutoScanTime
    {
        public readonly int? Iteration;
        public readonly string? IterationType;
        public readonly string? Time;
        public readonly ImmutableArray<string> WeekDays;

        [OutputConstructor]
        private HostAssurancePolicyAutoScanTime(
            int? iteration,

            string? iterationType,

            string? time,

            ImmutableArray<string> weekDays)
        {
            Iteration = iteration;
            IterationType = iterationType;
            Time = time;
            WeekDays = weekDays;
        }
    }
}
