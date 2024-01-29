// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aquasec from "@pulumi/aquasec";
 *
 * const functionRuntimePolicy = aquasec.getFunctionRuntimePolicy({
 *     name: "FunctionRuntimePolicyName",
 * });
 * export const functionRuntimePolicyDetails = functionRuntimePolicy;
 * ```
 */
export function getFunctionRuntimePolicy(args: GetFunctionRuntimePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionRuntimePolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aquasec:index/getFunctionRuntimePolicy:getFunctionRuntimePolicy", {
        "driftPreventions": args.driftPreventions,
        "executableBlacklists": args.executableBlacklists,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getFunctionRuntimePolicy.
 */
export interface GetFunctionRuntimePolicyArgs {
    /**
     * Drift prevention configuration.
     */
    driftPreventions?: inputs.GetFunctionRuntimePolicyDriftPrevention[];
    /**
     * Executable blacklist configuration.
     */
    executableBlacklists?: inputs.GetFunctionRuntimePolicyExecutableBlacklist[];
    name: string;
}

/**
 * A collection of values returned by getFunctionRuntimePolicy.
 */
export interface GetFunctionRuntimePolicyResult {
    /**
     * Indicates the application scope of the service.
     */
    readonly applicationScopes: string[];
    /**
     * Username of the account that created the service.
     */
    readonly author: string;
    /**
     * If true, prevent creation of malicious executables in functions during their runtime post invocation.
     */
    readonly blockMaliciousExecutables: boolean;
    /**
     * List of processes that will be allowed
     */
    readonly blockMaliciousExecutablesAllowedProcesses: string[];
    /**
     * If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.
     */
    readonly blockRunningExecutablesInTmpFolder: boolean;
    /**
     * List of executables that are prevented from running in containers.
     */
    readonly blockedExecutables: string[];
    /**
     * The description of the function runtime policy
     */
    readonly description: string;
    /**
     * Drift prevention configuration.
     */
    readonly driftPreventions?: outputs.GetFunctionRuntimePolicyDriftPrevention[];
    /**
     * Indicates if the runtime policy is enabled or not.
     */
    readonly enabled: boolean;
    /**
     * Indicates that policy should effect container execution (not just for audit).
     */
    readonly enforce: boolean;
    /**
     * Executable blacklist configuration.
     */
    readonly executableBlacklists?: outputs.GetFunctionRuntimePolicyExecutableBlacklist[];
    /**
     * Honeypot User ID (Access Key)
     */
    readonly honeypotAccessKey: string;
    /**
     * List of options to apply the honeypot on (Environment Vairable, Layer, File)
     */
    readonly honeypotApplyOns: string[];
    /**
     * Honeypot User Password (Secret Key)
     */
    readonly honeypotSecretKey: string;
    /**
     * Serverless application name
     */
    readonly honeypotServerlessAppName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the function runtime policy
     */
    readonly name: string;
    /**
     * Logical expression of how to compute the dependency of the scope variables.
     */
    readonly scopeExpression: string;
    /**
     * List of scope attributes.
     */
    readonly scopeVariables: outputs.GetFunctionRuntimePolicyScopeVariable[];
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aquasec from "@pulumi/aquasec";
 *
 * const functionRuntimePolicy = aquasec.getFunctionRuntimePolicy({
 *     name: "FunctionRuntimePolicyName",
 * });
 * export const functionRuntimePolicyDetails = functionRuntimePolicy;
 * ```
 */
export function getFunctionRuntimePolicyOutput(args: GetFunctionRuntimePolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFunctionRuntimePolicyResult> {
    return pulumi.output(args).apply((a: any) => getFunctionRuntimePolicy(a, opts))
}

/**
 * A collection of arguments for invoking getFunctionRuntimePolicy.
 */
export interface GetFunctionRuntimePolicyOutputArgs {
    /**
     * Drift prevention configuration.
     */
    driftPreventions?: pulumi.Input<pulumi.Input<inputs.GetFunctionRuntimePolicyDriftPreventionArgs>[]>;
    /**
     * Executable blacklist configuration.
     */
    executableBlacklists?: pulumi.Input<pulumi.Input<inputs.GetFunctionRuntimePolicyExecutableBlacklistArgs>[]>;
    name: pulumi.Input<string>;
}
