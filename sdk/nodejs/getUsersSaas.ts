// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The data source `aquasec.getUsersSaas` provides a method to query all saas users within the Aqua users management. The fields returned from this query are detailed in the Schema section below.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aquasec from "@pulumi/aquasec";
 *
 * const users = aquasec.getUsers({});
 * export const firstUserEmail = data.aquasec_users_saas.users.users[0].email;
 * ```
 */
export function getUsersSaas(opts?: pulumi.InvokeOptions): Promise<GetUsersSaasResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aquasec:index/getUsersSaas:getUsersSaas", {
    }, opts);
}

/**
 * A collection of values returned by getUsersSaas.
 */
export interface GetUsersSaasResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly users: outputs.GetUsersSaasUser[];
}
