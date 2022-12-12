// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The data source `aquasec.getRoles` provides a method to query all roles within the Aqua account managementrole database. The fields returned from this query are detailed in the Schema section below.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aquasec from "@pulumi/aquasec";
 *
 * const roles = aquasec.getRoles({});
 * export const firstUserName = roles.then(roles => roles.roles?[0]);
 * ```
 */
export function getRoles(opts?: pulumi.InvokeOptions): Promise<GetRolesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aquasec:index/getRoles:getRoles", {
    }, opts);
}

/**
 * A collection of values returned by getRoles.
 */
export interface GetRolesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly roles: outputs.GetRolesRole[];
}
