// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getIntegrationRegistry(args: GetIntegrationRegistryArgs, opts?: pulumi.InvokeOptions): Promise<GetIntegrationRegistryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aquasec:index/getIntegrationRegistry:getIntegrationRegistry", {
        "imageCreationDateCondition": args.imageCreationDateCondition,
        "name": args.name,
        "pullImageAge": args.pullImageAge,
        "pullImageCount": args.pullImageCount,
        "scannerNames": args.scannerNames,
        "scannerType": args.scannerType,
    }, opts);
}

/**
 * A collection of arguments for invoking getIntegrationRegistry.
 */
export interface GetIntegrationRegistryArgs {
    /**
     * Additional condition for pulling and rescanning images, Defaults to 'none'
     */
    imageCreationDateCondition?: string;
    /**
     * The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
     */
    name: string;
    /**
     * When auto pull image enabled, sets maximum age of auto pulled images
     */
    pullImageAge?: string;
    /**
     * When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
     */
    pullImageCount?: number;
    /**
     * List of scanner names
     */
    scannerNames?: string[];
    /**
     * Scanner type
     */
    scannerType?: string;
}

/**
 * A collection of values returned by getIntegrationRegistry.
 */
export interface GetIntegrationRegistryResult {
    /**
     * Automatically clean up images and repositories which are no longer present in the registry from Aqua console
     */
    readonly autoCleanup: boolean;
    /**
     * Whether to automatically pull images from the registry on creation and daily
     */
    readonly autoPull: boolean;
    /**
     * The interval in days to start pulling new images from the registry, Defaults to 1
     */
    readonly autoPullInterval: number;
    /**
     * Maximum number of repositories to pull every day, defaults to 100
     */
    readonly autoPullMax: number;
    /**
     * Whether to automatically pull and rescan images from the registry on creation and daily
     */
    readonly autoPullRescan: boolean;
    /**
     * The time of day to start pulling new images from the registry, in the format HH:MM (24-hour clock), defaults to 03:00
     */
    readonly autoPullTime: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Additional condition for pulling and rescanning images, Defaults to 'none'
     */
    readonly imageCreationDateCondition: string;
    /**
     * The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
     */
    readonly name: string;
    /**
     * The password for registry authentication
     */
    readonly password: string;
    /**
     * List of possible prefixes to image names pulled from the registry
     */
    readonly prefixes: string[];
    /**
     * When auto pull image enabled, sets maximum age of auto pulled images
     */
    readonly pullImageAge: string;
    /**
     * When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
     */
    readonly pullImageCount: number;
    /**
     * List of scanner names
     */
    readonly scannerNames: string[];
    /**
     * Scanner type
     */
    readonly scannerType: string;
    /**
     * Registry type (HUB / V1 / V2 / ENGINE / AWS / GCR).
     */
    readonly type: string;
    /**
     * The URL, address or region of the registry
     */
    readonly url: string;
    /**
     * The username for registry authentication.
     */
    readonly username: string;
}

export function getIntegrationRegistryOutput(args: GetIntegrationRegistryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIntegrationRegistryResult> {
    return pulumi.output(args).apply(a => getIntegrationRegistry(a, opts))
}

/**
 * A collection of arguments for invoking getIntegrationRegistry.
 */
export interface GetIntegrationRegistryOutputArgs {
    /**
     * Additional condition for pulling and rescanning images, Defaults to 'none'
     */
    imageCreationDateCondition?: pulumi.Input<string>;
    /**
     * The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces
     */
    name: pulumi.Input<string>;
    /**
     * When auto pull image enabled, sets maximum age of auto pulled images
     */
    pullImageAge?: pulumi.Input<string>;
    /**
     * When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
     */
    pullImageCount?: pulumi.Input<number>;
    /**
     * List of scanner names
     */
    scannerNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Scanner type
     */
    scannerType?: pulumi.Input<string>;
}
