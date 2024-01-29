// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getHostAssurancePolicy(args: GetHostAssurancePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetHostAssurancePolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aquasec:index/getHostAssurancePolicy:getHostAssurancePolicy", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getHostAssurancePolicy.
 */
export interface GetHostAssurancePolicyArgs {
    name: string;
}

/**
 * A collection of values returned by getHostAssurancePolicy.
 */
export interface GetHostAssurancePolicyResult {
    /**
     * List of explicitly allowed images.
     */
    readonly allowedImages: string[];
    readonly applicationScopes: string[];
    /**
     * Indicates if auditing for failures.
     */
    readonly auditOnFailure: boolean;
    /**
     * Name of user account that created the policy.
     */
    readonly author: string;
    readonly autoScanConfigured: boolean;
    readonly autoScanEnabled: boolean;
    readonly autoScanTimes: outputs.GetHostAssurancePolicyAutoScanTime[];
    /**
     * List of function's forbidden permissions.
     */
    readonly blacklistPermissions: string[];
    /**
     * Indicates if blacklist permissions is relevant.
     */
    readonly blacklistPermissionsEnabled: boolean;
    /**
     * List of blacklisted licenses.
     */
    readonly blacklistedLicenses: string[];
    /**
     * Indicates if license blacklist is relevant.
     */
    readonly blacklistedLicensesEnabled: boolean;
    /**
     * Indicates if failed images are blocked.
     */
    readonly blockFailed: boolean;
    readonly controlExcludeNoFix: boolean;
    /**
     * List of Custom user scripts for checks.
     */
    readonly customChecks: outputs.GetHostAssurancePolicyCustomCheck[];
    /**
     * Indicates if scanning should include custom checks.
     */
    readonly customChecksEnabled: boolean;
    readonly customSeverityEnabled: boolean;
    /**
     * Indicates if CVEs blacklist is relevant.
     */
    readonly cvesBlackListEnabled: boolean;
    /**
     * List of CVEs blacklisted items.
     */
    readonly cvesBlackLists: string[];
    /**
     * Indicates if CVEs whitelist is relevant.
     */
    readonly cvesWhiteListEnabled: boolean;
    /**
     * List of cves whitelisted licenses
     */
    readonly cvesWhiteLists: string[];
    /**
     * Identifier of the cvss severity.
     */
    readonly cvssSeverity: string;
    /**
     * Indicates if the cvss severity is scanned.
     */
    readonly cvssSeverityEnabled: boolean;
    /**
     * Indicates that policy should ignore cvss cases that do not have a known fix.
     */
    readonly cvssSeverityExcludeNoFix: boolean;
    readonly description: string;
    /**
     * Indicates if malware should block the image.
     */
    readonly disallowMalware: boolean;
    /**
     * Checks the host according to the Docker CIS benchmark, if Docker is found on the host.
     */
    readonly dockerCisEnabled: boolean;
    /**
     * Name of the container image.
     */
    readonly domain: string;
    readonly domainName: string;
    readonly dtaEnabled: boolean;
    readonly dtaSeverity: string;
    readonly enabled: boolean;
    readonly enforce: boolean;
    readonly enforceAfterDays: number;
    readonly enforceExcessivePermissions: boolean;
    readonly exceptionalMonitoredMalwarePaths: string[];
    /**
     * Indicates if cicd failures will fail the image.
     */
    readonly failCicd: boolean;
    readonly forbiddenLabels: outputs.GetHostAssurancePolicyForbiddenLabel[];
    readonly forbiddenLabelsEnabled: boolean;
    readonly forceMicroenforcer: boolean;
    readonly functionIntegrityEnabled: boolean;
    /**
     * The ID of this resource.
     */
    readonly id: string;
    readonly ignoreRecentlyPublishedVln: boolean;
    readonly ignoreRecentlyPublishedVlnPeriod: number;
    /**
     * Indicates if risk resources are ignored.
     */
    readonly ignoreRiskResourcesEnabled: boolean;
    /**
     * List of ignored risk resources.
     */
    readonly ignoredRiskResources: string[];
    /**
     * List of images.
     */
    readonly images: string[];
    /**
     * Performs a Kubernetes CIS benchmark check for the host.
     */
    readonly kubeCisEnabled: boolean;
    /**
     * List of labels.
     */
    readonly labels: string[];
    readonly malwareAction: string;
    /**
     * Value of allowed maximum score.
     */
    readonly maximumScore: number;
    /**
     * Indicates if exceeding the maximum score is scanned.
     */
    readonly maximumScoreEnabled: boolean;
    /**
     * Indicates that policy should ignore cases that do not have a known fix.
     */
    readonly maximumScoreExcludeNoFix: boolean;
    readonly monitoredMalwarePaths: string[];
    readonly name: string;
    /**
     * Indicates if raise a warning for images that should only be run as root.
     */
    readonly onlyNoneRootUsers: boolean;
    /**
     * Indicates if packages blacklist is relevant.
     */
    readonly packagesBlackListEnabled: boolean;
    /**
     * List of blacklisted images.
     */
    readonly packagesBlackLists: outputs.GetHostAssurancePolicyPackagesBlackList[];
    /**
     * Indicates if packages whitelist is relevant.
     */
    readonly packagesWhiteListEnabled: boolean;
    /**
     * List of whitelisted images.
     */
    readonly packagesWhiteLists: outputs.GetHostAssurancePolicyPackagesWhiteList[];
    readonly partialResultsImageFail: boolean;
    readonly readOnly: boolean;
    /**
     * List of registries.
     */
    readonly registries: string[];
    readonly registry: string;
    readonly requiredLabels: outputs.GetHostAssurancePolicyRequiredLabel[];
    readonly requiredLabelsEnabled: boolean;
    readonly scanNfsMounts: boolean;
    /**
     * Indicates if scan should include sensitive data in the image.
     */
    readonly scanSensitiveData: boolean;
    /**
     * Indicates if scanning should include scap.
     */
    readonly scapEnabled: boolean;
    /**
     * List of SCAP user scripts for checks.
     */
    readonly scapFiles: string[];
    readonly scopes: outputs.GetHostAssurancePolicyScope[];
    /**
     * List of trusted images.
     */
    readonly trustedBaseImages: outputs.GetHostAssurancePolicyTrustedBaseImage[];
    /**
     * Indicates if list of trusted base images is relevant.
     */
    readonly trustedBaseImagesEnabled: boolean;
    /**
     * List of whitelisted licenses.
     */
    readonly whitelistedLicenses: string[];
    /**
     * Indicates if license blacklist is relevant.
     */
    readonly whitelistedLicensesEnabled: boolean;
}
export function getHostAssurancePolicyOutput(args: GetHostAssurancePolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHostAssurancePolicyResult> {
    return pulumi.output(args).apply((a: any) => getHostAssurancePolicy(a, opts))
}

/**
 * A collection of arguments for invoking getHostAssurancePolicy.
 */
export interface GetHostAssurancePolicyOutputArgs {
    name: pulumi.Input<string>;
}
