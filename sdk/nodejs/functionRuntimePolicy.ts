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
 * const functionRuntimePolicy = new aquasec.FunctionRuntimePolicy("function_runtime_policy", {
 *     applicationScopes: ["Global"],
 *     blockMaliciousExecutables: true,
 *     blockMaliciousExecutablesAllowedProcesses: [
 *         "proc1",
 *         "proc2",
 *     ],
 *     blockRunningExecutablesInTmpFolder: true,
 *     blockedExecutables: [
 *         "exe1",
 *         "exe2",
 *     ],
 *     description: "function_runtime_policy",
 *     enabled: true,
 *     enforce: false,
 *     scopeVariables: [
 *         {
 *             attribute: "kubernetes.cluster",
 *             value: "default",
 *         },
 *         {
 *             attribute: "kubernetes.label",
 *             name: "app",
 *             value: "aqua",
 *         },
 *     ],
 * });
 * ```
 */
export class FunctionRuntimePolicy extends pulumi.CustomResource {
    /**
     * Get an existing FunctionRuntimePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionRuntimePolicyState, opts?: pulumi.CustomResourceOptions): FunctionRuntimePolicy {
        return new FunctionRuntimePolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aquasec:index/functionRuntimePolicy:FunctionRuntimePolicy';

    /**
     * Returns true if the given object is an instance of FunctionRuntimePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FunctionRuntimePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FunctionRuntimePolicy.__pulumiType;
    }

    /**
     * Indicates the application scope of the service.
     */
    public readonly applicationScopes!: pulumi.Output<string[]>;
    /**
     * Username of the account that created the service.
     */
    public /*out*/ readonly author!: pulumi.Output<string>;
    /**
     * If true, prevent creation of malicious executables in functions during their runtime post invocation.
     */
    public readonly blockMaliciousExecutables!: pulumi.Output<boolean | undefined>;
    /**
     * List of processes that will be allowed
     */
    public readonly blockMaliciousExecutablesAllowedProcesses!: pulumi.Output<string[] | undefined>;
    /**
     * If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.
     */
    public readonly blockRunningExecutablesInTmpFolder!: pulumi.Output<boolean | undefined>;
    /**
     * List of executables that are prevented from running in containers.
     */
    public readonly blockedExecutables!: pulumi.Output<string[] | undefined>;
    /**
     * The description of the function runtime policy
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Indicates if the runtime policy is enabled or not.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates that policy should effect container execution (not just for audit).
     */
    public readonly enforce!: pulumi.Output<boolean | undefined>;
    /**
     * Honeypot User ID (Access Key)
     */
    public readonly honeypotAccessKey!: pulumi.Output<string | undefined>;
    /**
     * List of options to apply the honeypot on (Environment Vairable, Layer, File)
     */
    public readonly honeypotApplyOns!: pulumi.Output<string[] | undefined>;
    /**
     * Honeypot User Password (Secret Key)
     */
    public readonly honeypotSecretKey!: pulumi.Output<string | undefined>;
    /**
     * Serverless application name
     */
    public readonly honeypotServerlessAppName!: pulumi.Output<string | undefined>;
    /**
     * Name of the function runtime policy
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Logical expression of how to compute the dependency of the scope variables.
     */
    public readonly scopeExpression!: pulumi.Output<string>;
    /**
     * List of scope attributes.
     */
    public readonly scopeVariables!: pulumi.Output<outputs.FunctionRuntimePolicyScopeVariable[]>;

    /**
     * Create a FunctionRuntimePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FunctionRuntimePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionRuntimePolicyArgs | FunctionRuntimePolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionRuntimePolicyState | undefined;
            resourceInputs["applicationScopes"] = state ? state.applicationScopes : undefined;
            resourceInputs["author"] = state ? state.author : undefined;
            resourceInputs["blockMaliciousExecutables"] = state ? state.blockMaliciousExecutables : undefined;
            resourceInputs["blockMaliciousExecutablesAllowedProcesses"] = state ? state.blockMaliciousExecutablesAllowedProcesses : undefined;
            resourceInputs["blockRunningExecutablesInTmpFolder"] = state ? state.blockRunningExecutablesInTmpFolder : undefined;
            resourceInputs["blockedExecutables"] = state ? state.blockedExecutables : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["enforce"] = state ? state.enforce : undefined;
            resourceInputs["honeypotAccessKey"] = state ? state.honeypotAccessKey : undefined;
            resourceInputs["honeypotApplyOns"] = state ? state.honeypotApplyOns : undefined;
            resourceInputs["honeypotSecretKey"] = state ? state.honeypotSecretKey : undefined;
            resourceInputs["honeypotServerlessAppName"] = state ? state.honeypotServerlessAppName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["scopeExpression"] = state ? state.scopeExpression : undefined;
            resourceInputs["scopeVariables"] = state ? state.scopeVariables : undefined;
        } else {
            const args = argsOrState as FunctionRuntimePolicyArgs | undefined;
            resourceInputs["applicationScopes"] = args ? args.applicationScopes : undefined;
            resourceInputs["blockMaliciousExecutables"] = args ? args.blockMaliciousExecutables : undefined;
            resourceInputs["blockMaliciousExecutablesAllowedProcesses"] = args ? args.blockMaliciousExecutablesAllowedProcesses : undefined;
            resourceInputs["blockRunningExecutablesInTmpFolder"] = args ? args.blockRunningExecutablesInTmpFolder : undefined;
            resourceInputs["blockedExecutables"] = args ? args.blockedExecutables : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["enforce"] = args ? args.enforce : undefined;
            resourceInputs["honeypotAccessKey"] = args ? args.honeypotAccessKey : undefined;
            resourceInputs["honeypotApplyOns"] = args ? args.honeypotApplyOns : undefined;
            resourceInputs["honeypotSecretKey"] = args?.honeypotSecretKey ? pulumi.secret(args.honeypotSecretKey) : undefined;
            resourceInputs["honeypotServerlessAppName"] = args ? args.honeypotServerlessAppName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopeExpression"] = args ? args.scopeExpression : undefined;
            resourceInputs["scopeVariables"] = args ? args.scopeVariables : undefined;
            resourceInputs["author"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["honeypotSecretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(FunctionRuntimePolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FunctionRuntimePolicy resources.
 */
export interface FunctionRuntimePolicyState {
    /**
     * Indicates the application scope of the service.
     */
    applicationScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Username of the account that created the service.
     */
    author?: pulumi.Input<string>;
    /**
     * If true, prevent creation of malicious executables in functions during their runtime post invocation.
     */
    blockMaliciousExecutables?: pulumi.Input<boolean>;
    /**
     * List of processes that will be allowed
     */
    blockMaliciousExecutablesAllowedProcesses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.
     */
    blockRunningExecutablesInTmpFolder?: pulumi.Input<boolean>;
    /**
     * List of executables that are prevented from running in containers.
     */
    blockedExecutables?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the function runtime policy
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates if the runtime policy is enabled or not.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Indicates that policy should effect container execution (not just for audit).
     */
    enforce?: pulumi.Input<boolean>;
    /**
     * Honeypot User ID (Access Key)
     */
    honeypotAccessKey?: pulumi.Input<string>;
    /**
     * List of options to apply the honeypot on (Environment Vairable, Layer, File)
     */
    honeypotApplyOns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Honeypot User Password (Secret Key)
     */
    honeypotSecretKey?: pulumi.Input<string>;
    /**
     * Serverless application name
     */
    honeypotServerlessAppName?: pulumi.Input<string>;
    /**
     * Name of the function runtime policy
     */
    name?: pulumi.Input<string>;
    /**
     * Logical expression of how to compute the dependency of the scope variables.
     */
    scopeExpression?: pulumi.Input<string>;
    /**
     * List of scope attributes.
     */
    scopeVariables?: pulumi.Input<pulumi.Input<inputs.FunctionRuntimePolicyScopeVariable>[]>;
}

/**
 * The set of arguments for constructing a FunctionRuntimePolicy resource.
 */
export interface FunctionRuntimePolicyArgs {
    /**
     * Indicates the application scope of the service.
     */
    applicationScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If true, prevent creation of malicious executables in functions during their runtime post invocation.
     */
    blockMaliciousExecutables?: pulumi.Input<boolean>;
    /**
     * List of processes that will be allowed
     */
    blockMaliciousExecutablesAllowedProcesses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If true, prevent running of executables in functions locate in /tmp folder during their runtime post invocation.
     */
    blockRunningExecutablesInTmpFolder?: pulumi.Input<boolean>;
    /**
     * List of executables that are prevented from running in containers.
     */
    blockedExecutables?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the function runtime policy
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates if the runtime policy is enabled or not.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Indicates that policy should effect container execution (not just for audit).
     */
    enforce?: pulumi.Input<boolean>;
    /**
     * Honeypot User ID (Access Key)
     */
    honeypotAccessKey?: pulumi.Input<string>;
    /**
     * List of options to apply the honeypot on (Environment Vairable, Layer, File)
     */
    honeypotApplyOns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Honeypot User Password (Secret Key)
     */
    honeypotSecretKey?: pulumi.Input<string>;
    /**
     * Serverless application name
     */
    honeypotServerlessAppName?: pulumi.Input<string>;
    /**
     * Name of the function runtime policy
     */
    name?: pulumi.Input<string>;
    /**
     * Logical expression of how to compute the dependency of the scope variables.
     */
    scopeExpression?: pulumi.Input<string>;
    /**
     * List of scope attributes.
     */
    scopeVariables?: pulumi.Input<pulumi.Input<inputs.FunctionRuntimePolicyScopeVariable>[]>;
}
