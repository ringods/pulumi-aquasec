// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ImageAssurancePolicy extends pulumi.CustomResource {
    /**
     * Get an existing ImageAssurancePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageAssurancePolicyState, opts?: pulumi.CustomResourceOptions): ImageAssurancePolicy {
        return new ImageAssurancePolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aquasec:index/imageAssurancePolicy:ImageAssurancePolicy';

    /**
     * Returns true if the given object is an instance of ImageAssurancePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ImageAssurancePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ImageAssurancePolicy.__pulumiType;
    }

    /**
     * List of explicitly allowed images.
     */
    public readonly allowedImages!: pulumi.Output<string[] | undefined>;
    public readonly applicationScopes!: pulumi.Output<string[]>;
    /**
     * Indicates if auditing for failures.
     */
    public readonly auditOnFailure!: pulumi.Output<boolean | undefined>;
    /**
     * Name of user account that created the policy.
     */
    public /*out*/ readonly author!: pulumi.Output<string>;
    public readonly autoScanConfigured!: pulumi.Output<boolean | undefined>;
    public readonly autoScanEnabled!: pulumi.Output<boolean | undefined>;
    public readonly autoScanTimes!: pulumi.Output<outputs.ImageAssurancePolicyAutoScanTime[]>;
    /**
     * List of function's forbidden permissions.
     */
    public readonly blacklistPermissions!: pulumi.Output<string[] | undefined>;
    /**
     * Indicates if blacklist permissions is relevant.
     */
    public readonly blacklistPermissionsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of blacklisted licenses.
     */
    public readonly blacklistedLicenses!: pulumi.Output<string[] | undefined>;
    /**
     * Lndicates if license blacklist is relevant.
     */
    public readonly blacklistedLicensesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if failed images are blocked.
     */
    public readonly blockFailed!: pulumi.Output<boolean | undefined>;
    public readonly controlExcludeNoFix!: pulumi.Output<boolean | undefined>;
    /**
     * List of Custom user scripts for checks.
     */
    public readonly customChecks!: pulumi.Output<outputs.ImageAssurancePolicyCustomCheck[] | undefined>;
    /**
     * Indicates if scanning should include custom checks.
     */
    public readonly customChecksEnabled!: pulumi.Output<boolean | undefined>;
    public readonly customSeverityEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if cves blacklist is relevant.
     */
    public readonly cvesBlackListEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of cves blacklisted items.
     */
    public readonly cvesBlackLists!: pulumi.Output<string[] | undefined>;
    /**
     * Indicates if cves whitelist is relevant.
     */
    public readonly cvesWhiteListEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of cves whitelisted licenses
     */
    public readonly cvesWhiteLists!: pulumi.Output<string[] | undefined>;
    /**
     * Identifier of the cvss severity.
     */
    public readonly cvssSeverity!: pulumi.Output<string | undefined>;
    /**
     * Indicates if the cvss severity is scanned.
     */
    public readonly cvssSeverityEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates that policy should ignore cvss cases that do not have a known fix.
     */
    public readonly cvssSeverityExcludeNoFix!: pulumi.Output<boolean | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Indicates if malware should block the image.
     */
    public readonly disallowMalware!: pulumi.Output<boolean | undefined>;
    public readonly dockerCisEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the container image.
     */
    public readonly domain!: pulumi.Output<string | undefined>;
    public readonly domainName!: pulumi.Output<string | undefined>;
    public readonly dtaEnabled!: pulumi.Output<boolean | undefined>;
    public readonly dtaSeverity!: pulumi.Output<string | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly enforce!: pulumi.Output<boolean | undefined>;
    public readonly enforceAfterDays!: pulumi.Output<number | undefined>;
    public readonly enforceExcessivePermissions!: pulumi.Output<boolean | undefined>;
    public readonly exceptionalMonitoredMalwarePaths!: pulumi.Output<string[] | undefined>;
    /**
     * Indicates if cicd failures will fail the image.
     */
    public readonly failCicd!: pulumi.Output<boolean | undefined>;
    public readonly forbiddenLabels!: pulumi.Output<outputs.ImageAssurancePolicyForbiddenLabel[] | undefined>;
    public readonly forbiddenLabelsEnabled!: pulumi.Output<boolean | undefined>;
    public readonly forceMicroenforcer!: pulumi.Output<boolean | undefined>;
    public readonly functionIntegrityEnabled!: pulumi.Output<boolean | undefined>;
    public readonly ignoreRecentlyPublishedVln!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly ignoreRecentlyPublishedVlnPeriod!: pulumi.Output<number>;
    /**
     * Indicates if risk resources are ignored.
     */
    public readonly ignoreRiskResourcesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of ignored risk resources.
     */
    public readonly ignoredRiskResources!: pulumi.Output<string[] | undefined>;
    /**
     * List of images.
     */
    public readonly images!: pulumi.Output<string[] | undefined>;
    public readonly kubeCisEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of labels.
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    public readonly malwareAction!: pulumi.Output<string | undefined>;
    /**
     * Value of allowed maximum score.
     */
    public readonly maximumScore!: pulumi.Output<number | undefined>;
    /**
     * Indicates if exceeding the maximum score is scanned.
     */
    public readonly maximumScoreEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates that policy should ignore cases that do not have a known fix.
     */
    public readonly maximumScoreExcludeNoFix!: pulumi.Output<boolean | undefined>;
    public readonly monitoredMalwarePaths!: pulumi.Output<string[] | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * Indicates if raise a warning for images that should only be run as root.
     */
    public readonly onlyNoneRootUsers!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if packages blacklist is relevant.
     */
    public readonly packagesBlackListEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of backlisted images.
     */
    public readonly packagesBlackLists!: pulumi.Output<outputs.ImageAssurancePolicyPackagesBlackList[] | undefined>;
    /**
     * Indicates if packages whitelist is relevant.
     */
    public readonly packagesWhiteListEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of whitelisted images.
     */
    public readonly packagesWhiteLists!: pulumi.Output<outputs.ImageAssurancePolicyPackagesWhiteList[] | undefined>;
    public readonly partialResultsImageFail!: pulumi.Output<boolean | undefined>;
    public readonly readOnly!: pulumi.Output<boolean | undefined>;
    /**
     * List of registries.
     */
    public readonly registries!: pulumi.Output<string[] | undefined>;
    public readonly registry!: pulumi.Output<string | undefined>;
    public readonly requiredLabels!: pulumi.Output<outputs.ImageAssurancePolicyRequiredLabel[] | undefined>;
    public readonly requiredLabelsEnabled!: pulumi.Output<boolean | undefined>;
    public readonly scanNfsMounts!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if scan should include sensitive data in the image.
     */
    public readonly scanSensitiveData!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if scanning should include scap.
     */
    public readonly scapEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of SCAP user scripts for checks.
     */
    public readonly scapFiles!: pulumi.Output<string[] | undefined>;
    public readonly scopes!: pulumi.Output<outputs.ImageAssurancePolicyScope[]>;
    /**
     * List of trusted images.
     */
    public readonly trustedBaseImages!: pulumi.Output<outputs.ImageAssurancePolicyTrustedBaseImage[] | undefined>;
    /**
     * Indicates if list of trusted base images is relevant.
     */
    public readonly trustedBaseImagesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of whitelisted licenses.
     */
    public readonly whitelistedLicenses!: pulumi.Output<string[] | undefined>;
    /**
     * Indicates if license blacklist is relevant.
     */
    public readonly whitelistedLicensesEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ImageAssurancePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageAssurancePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageAssurancePolicyArgs | ImageAssurancePolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImageAssurancePolicyState | undefined;
            resourceInputs["allowedImages"] = state ? state.allowedImages : undefined;
            resourceInputs["applicationScopes"] = state ? state.applicationScopes : undefined;
            resourceInputs["auditOnFailure"] = state ? state.auditOnFailure : undefined;
            resourceInputs["author"] = state ? state.author : undefined;
            resourceInputs["autoScanConfigured"] = state ? state.autoScanConfigured : undefined;
            resourceInputs["autoScanEnabled"] = state ? state.autoScanEnabled : undefined;
            resourceInputs["autoScanTimes"] = state ? state.autoScanTimes : undefined;
            resourceInputs["blacklistPermissions"] = state ? state.blacklistPermissions : undefined;
            resourceInputs["blacklistPermissionsEnabled"] = state ? state.blacklistPermissionsEnabled : undefined;
            resourceInputs["blacklistedLicenses"] = state ? state.blacklistedLicenses : undefined;
            resourceInputs["blacklistedLicensesEnabled"] = state ? state.blacklistedLicensesEnabled : undefined;
            resourceInputs["blockFailed"] = state ? state.blockFailed : undefined;
            resourceInputs["controlExcludeNoFix"] = state ? state.controlExcludeNoFix : undefined;
            resourceInputs["customChecks"] = state ? state.customChecks : undefined;
            resourceInputs["customChecksEnabled"] = state ? state.customChecksEnabled : undefined;
            resourceInputs["customSeverityEnabled"] = state ? state.customSeverityEnabled : undefined;
            resourceInputs["cvesBlackListEnabled"] = state ? state.cvesBlackListEnabled : undefined;
            resourceInputs["cvesBlackLists"] = state ? state.cvesBlackLists : undefined;
            resourceInputs["cvesWhiteListEnabled"] = state ? state.cvesWhiteListEnabled : undefined;
            resourceInputs["cvesWhiteLists"] = state ? state.cvesWhiteLists : undefined;
            resourceInputs["cvssSeverity"] = state ? state.cvssSeverity : undefined;
            resourceInputs["cvssSeverityEnabled"] = state ? state.cvssSeverityEnabled : undefined;
            resourceInputs["cvssSeverityExcludeNoFix"] = state ? state.cvssSeverityExcludeNoFix : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disallowMalware"] = state ? state.disallowMalware : undefined;
            resourceInputs["dockerCisEnabled"] = state ? state.dockerCisEnabled : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["dtaEnabled"] = state ? state.dtaEnabled : undefined;
            resourceInputs["dtaSeverity"] = state ? state.dtaSeverity : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["enforce"] = state ? state.enforce : undefined;
            resourceInputs["enforceAfterDays"] = state ? state.enforceAfterDays : undefined;
            resourceInputs["enforceExcessivePermissions"] = state ? state.enforceExcessivePermissions : undefined;
            resourceInputs["exceptionalMonitoredMalwarePaths"] = state ? state.exceptionalMonitoredMalwarePaths : undefined;
            resourceInputs["failCicd"] = state ? state.failCicd : undefined;
            resourceInputs["forbiddenLabels"] = state ? state.forbiddenLabels : undefined;
            resourceInputs["forbiddenLabelsEnabled"] = state ? state.forbiddenLabelsEnabled : undefined;
            resourceInputs["forceMicroenforcer"] = state ? state.forceMicroenforcer : undefined;
            resourceInputs["functionIntegrityEnabled"] = state ? state.functionIntegrityEnabled : undefined;
            resourceInputs["ignoreRecentlyPublishedVln"] = state ? state.ignoreRecentlyPublishedVln : undefined;
            resourceInputs["ignoreRecentlyPublishedVlnPeriod"] = state ? state.ignoreRecentlyPublishedVlnPeriod : undefined;
            resourceInputs["ignoreRiskResourcesEnabled"] = state ? state.ignoreRiskResourcesEnabled : undefined;
            resourceInputs["ignoredRiskResources"] = state ? state.ignoredRiskResources : undefined;
            resourceInputs["images"] = state ? state.images : undefined;
            resourceInputs["kubeCisEnabled"] = state ? state.kubeCisEnabled : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["malwareAction"] = state ? state.malwareAction : undefined;
            resourceInputs["maximumScore"] = state ? state.maximumScore : undefined;
            resourceInputs["maximumScoreEnabled"] = state ? state.maximumScoreEnabled : undefined;
            resourceInputs["maximumScoreExcludeNoFix"] = state ? state.maximumScoreExcludeNoFix : undefined;
            resourceInputs["monitoredMalwarePaths"] = state ? state.monitoredMalwarePaths : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["onlyNoneRootUsers"] = state ? state.onlyNoneRootUsers : undefined;
            resourceInputs["packagesBlackListEnabled"] = state ? state.packagesBlackListEnabled : undefined;
            resourceInputs["packagesBlackLists"] = state ? state.packagesBlackLists : undefined;
            resourceInputs["packagesWhiteListEnabled"] = state ? state.packagesWhiteListEnabled : undefined;
            resourceInputs["packagesWhiteLists"] = state ? state.packagesWhiteLists : undefined;
            resourceInputs["partialResultsImageFail"] = state ? state.partialResultsImageFail : undefined;
            resourceInputs["readOnly"] = state ? state.readOnly : undefined;
            resourceInputs["registries"] = state ? state.registries : undefined;
            resourceInputs["registry"] = state ? state.registry : undefined;
            resourceInputs["requiredLabels"] = state ? state.requiredLabels : undefined;
            resourceInputs["requiredLabelsEnabled"] = state ? state.requiredLabelsEnabled : undefined;
            resourceInputs["scanNfsMounts"] = state ? state.scanNfsMounts : undefined;
            resourceInputs["scanSensitiveData"] = state ? state.scanSensitiveData : undefined;
            resourceInputs["scapEnabled"] = state ? state.scapEnabled : undefined;
            resourceInputs["scapFiles"] = state ? state.scapFiles : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
            resourceInputs["trustedBaseImages"] = state ? state.trustedBaseImages : undefined;
            resourceInputs["trustedBaseImagesEnabled"] = state ? state.trustedBaseImagesEnabled : undefined;
            resourceInputs["whitelistedLicenses"] = state ? state.whitelistedLicenses : undefined;
            resourceInputs["whitelistedLicensesEnabled"] = state ? state.whitelistedLicensesEnabled : undefined;
        } else {
            const args = argsOrState as ImageAssurancePolicyArgs | undefined;
            if ((!args || args.applicationScopes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationScopes'");
            }
            resourceInputs["allowedImages"] = args ? args.allowedImages : undefined;
            resourceInputs["applicationScopes"] = args ? args.applicationScopes : undefined;
            resourceInputs["auditOnFailure"] = args ? args.auditOnFailure : undefined;
            resourceInputs["autoScanConfigured"] = args ? args.autoScanConfigured : undefined;
            resourceInputs["autoScanEnabled"] = args ? args.autoScanEnabled : undefined;
            resourceInputs["autoScanTimes"] = args ? args.autoScanTimes : undefined;
            resourceInputs["blacklistPermissions"] = args ? args.blacklistPermissions : undefined;
            resourceInputs["blacklistPermissionsEnabled"] = args ? args.blacklistPermissionsEnabled : undefined;
            resourceInputs["blacklistedLicenses"] = args ? args.blacklistedLicenses : undefined;
            resourceInputs["blacklistedLicensesEnabled"] = args ? args.blacklistedLicensesEnabled : undefined;
            resourceInputs["blockFailed"] = args ? args.blockFailed : undefined;
            resourceInputs["controlExcludeNoFix"] = args ? args.controlExcludeNoFix : undefined;
            resourceInputs["customChecks"] = args ? args.customChecks : undefined;
            resourceInputs["customChecksEnabled"] = args ? args.customChecksEnabled : undefined;
            resourceInputs["customSeverityEnabled"] = args ? args.customSeverityEnabled : undefined;
            resourceInputs["cvesBlackListEnabled"] = args ? args.cvesBlackListEnabled : undefined;
            resourceInputs["cvesBlackLists"] = args ? args.cvesBlackLists : undefined;
            resourceInputs["cvesWhiteListEnabled"] = args ? args.cvesWhiteListEnabled : undefined;
            resourceInputs["cvesWhiteLists"] = args ? args.cvesWhiteLists : undefined;
            resourceInputs["cvssSeverity"] = args ? args.cvssSeverity : undefined;
            resourceInputs["cvssSeverityEnabled"] = args ? args.cvssSeverityEnabled : undefined;
            resourceInputs["cvssSeverityExcludeNoFix"] = args ? args.cvssSeverityExcludeNoFix : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disallowMalware"] = args ? args.disallowMalware : undefined;
            resourceInputs["dockerCisEnabled"] = args ? args.dockerCisEnabled : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["dtaEnabled"] = args ? args.dtaEnabled : undefined;
            resourceInputs["dtaSeverity"] = args ? args.dtaSeverity : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["enforce"] = args ? args.enforce : undefined;
            resourceInputs["enforceAfterDays"] = args ? args.enforceAfterDays : undefined;
            resourceInputs["enforceExcessivePermissions"] = args ? args.enforceExcessivePermissions : undefined;
            resourceInputs["exceptionalMonitoredMalwarePaths"] = args ? args.exceptionalMonitoredMalwarePaths : undefined;
            resourceInputs["failCicd"] = args ? args.failCicd : undefined;
            resourceInputs["forbiddenLabels"] = args ? args.forbiddenLabels : undefined;
            resourceInputs["forbiddenLabelsEnabled"] = args ? args.forbiddenLabelsEnabled : undefined;
            resourceInputs["forceMicroenforcer"] = args ? args.forceMicroenforcer : undefined;
            resourceInputs["functionIntegrityEnabled"] = args ? args.functionIntegrityEnabled : undefined;
            resourceInputs["ignoreRecentlyPublishedVln"] = args ? args.ignoreRecentlyPublishedVln : undefined;
            resourceInputs["ignoreRiskResourcesEnabled"] = args ? args.ignoreRiskResourcesEnabled : undefined;
            resourceInputs["ignoredRiskResources"] = args ? args.ignoredRiskResources : undefined;
            resourceInputs["images"] = args ? args.images : undefined;
            resourceInputs["kubeCisEnabled"] = args ? args.kubeCisEnabled : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["malwareAction"] = args ? args.malwareAction : undefined;
            resourceInputs["maximumScore"] = args ? args.maximumScore : undefined;
            resourceInputs["maximumScoreEnabled"] = args ? args.maximumScoreEnabled : undefined;
            resourceInputs["maximumScoreExcludeNoFix"] = args ? args.maximumScoreExcludeNoFix : undefined;
            resourceInputs["monitoredMalwarePaths"] = args ? args.monitoredMalwarePaths : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["onlyNoneRootUsers"] = args ? args.onlyNoneRootUsers : undefined;
            resourceInputs["packagesBlackListEnabled"] = args ? args.packagesBlackListEnabled : undefined;
            resourceInputs["packagesBlackLists"] = args ? args.packagesBlackLists : undefined;
            resourceInputs["packagesWhiteListEnabled"] = args ? args.packagesWhiteListEnabled : undefined;
            resourceInputs["packagesWhiteLists"] = args ? args.packagesWhiteLists : undefined;
            resourceInputs["partialResultsImageFail"] = args ? args.partialResultsImageFail : undefined;
            resourceInputs["readOnly"] = args ? args.readOnly : undefined;
            resourceInputs["registries"] = args ? args.registries : undefined;
            resourceInputs["registry"] = args ? args.registry : undefined;
            resourceInputs["requiredLabels"] = args ? args.requiredLabels : undefined;
            resourceInputs["requiredLabelsEnabled"] = args ? args.requiredLabelsEnabled : undefined;
            resourceInputs["scanNfsMounts"] = args ? args.scanNfsMounts : undefined;
            resourceInputs["scanSensitiveData"] = args ? args.scanSensitiveData : undefined;
            resourceInputs["scapEnabled"] = args ? args.scapEnabled : undefined;
            resourceInputs["scapFiles"] = args ? args.scapFiles : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["trustedBaseImages"] = args ? args.trustedBaseImages : undefined;
            resourceInputs["trustedBaseImagesEnabled"] = args ? args.trustedBaseImagesEnabled : undefined;
            resourceInputs["whitelistedLicenses"] = args ? args.whitelistedLicenses : undefined;
            resourceInputs["whitelistedLicensesEnabled"] = args ? args.whitelistedLicensesEnabled : undefined;
            resourceInputs["author"] = undefined /*out*/;
            resourceInputs["ignoreRecentlyPublishedVlnPeriod"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ImageAssurancePolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ImageAssurancePolicy resources.
 */
export interface ImageAssurancePolicyState {
    /**
     * List of explicitly allowed images.
     */
    allowedImages?: pulumi.Input<pulumi.Input<string>[]>;
    applicationScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates if auditing for failures.
     */
    auditOnFailure?: pulumi.Input<boolean>;
    /**
     * Name of user account that created the policy.
     */
    author?: pulumi.Input<string>;
    autoScanConfigured?: pulumi.Input<boolean>;
    autoScanEnabled?: pulumi.Input<boolean>;
    autoScanTimes?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyAutoScanTime>[]>;
    /**
     * List of function's forbidden permissions.
     */
    blacklistPermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates if blacklist permissions is relevant.
     */
    blacklistPermissionsEnabled?: pulumi.Input<boolean>;
    /**
     * List of blacklisted licenses.
     */
    blacklistedLicenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Lndicates if license blacklist is relevant.
     */
    blacklistedLicensesEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates if failed images are blocked.
     */
    blockFailed?: pulumi.Input<boolean>;
    controlExcludeNoFix?: pulumi.Input<boolean>;
    /**
     * List of Custom user scripts for checks.
     */
    customChecks?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyCustomCheck>[]>;
    /**
     * Indicates if scanning should include custom checks.
     */
    customChecksEnabled?: pulumi.Input<boolean>;
    customSeverityEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates if cves blacklist is relevant.
     */
    cvesBlackListEnabled?: pulumi.Input<boolean>;
    /**
     * List of cves blacklisted items.
     */
    cvesBlackLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates if cves whitelist is relevant.
     */
    cvesWhiteListEnabled?: pulumi.Input<boolean>;
    /**
     * List of cves whitelisted licenses
     */
    cvesWhiteLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Identifier of the cvss severity.
     */
    cvssSeverity?: pulumi.Input<string>;
    /**
     * Indicates if the cvss severity is scanned.
     */
    cvssSeverityEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates that policy should ignore cvss cases that do not have a known fix.
     */
    cvssSeverityExcludeNoFix?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    /**
     * Indicates if malware should block the image.
     */
    disallowMalware?: pulumi.Input<boolean>;
    dockerCisEnabled?: pulumi.Input<boolean>;
    /**
     * Name of the container image.
     */
    domain?: pulumi.Input<string>;
    domainName?: pulumi.Input<string>;
    dtaEnabled?: pulumi.Input<boolean>;
    dtaSeverity?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    enforce?: pulumi.Input<boolean>;
    enforceAfterDays?: pulumi.Input<number>;
    enforceExcessivePermissions?: pulumi.Input<boolean>;
    exceptionalMonitoredMalwarePaths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates if cicd failures will fail the image.
     */
    failCicd?: pulumi.Input<boolean>;
    forbiddenLabels?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyForbiddenLabel>[]>;
    forbiddenLabelsEnabled?: pulumi.Input<boolean>;
    forceMicroenforcer?: pulumi.Input<boolean>;
    functionIntegrityEnabled?: pulumi.Input<boolean>;
    ignoreRecentlyPublishedVln?: pulumi.Input<boolean>;
    ignoreRecentlyPublishedVlnPeriod?: pulumi.Input<number>;
    /**
     * Indicates if risk resources are ignored.
     */
    ignoreRiskResourcesEnabled?: pulumi.Input<boolean>;
    /**
     * List of ignored risk resources.
     */
    ignoredRiskResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of images.
     */
    images?: pulumi.Input<pulumi.Input<string>[]>;
    kubeCisEnabled?: pulumi.Input<boolean>;
    /**
     * List of labels.
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    malwareAction?: pulumi.Input<string>;
    /**
     * Value of allowed maximum score.
     */
    maximumScore?: pulumi.Input<number>;
    /**
     * Indicates if exceeding the maximum score is scanned.
     */
    maximumScoreEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates that policy should ignore cases that do not have a known fix.
     */
    maximumScoreExcludeNoFix?: pulumi.Input<boolean>;
    monitoredMalwarePaths?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    /**
     * Indicates if raise a warning for images that should only be run as root.
     */
    onlyNoneRootUsers?: pulumi.Input<boolean>;
    /**
     * Indicates if packages blacklist is relevant.
     */
    packagesBlackListEnabled?: pulumi.Input<boolean>;
    /**
     * List of backlisted images.
     */
    packagesBlackLists?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyPackagesBlackList>[]>;
    /**
     * Indicates if packages whitelist is relevant.
     */
    packagesWhiteListEnabled?: pulumi.Input<boolean>;
    /**
     * List of whitelisted images.
     */
    packagesWhiteLists?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyPackagesWhiteList>[]>;
    partialResultsImageFail?: pulumi.Input<boolean>;
    readOnly?: pulumi.Input<boolean>;
    /**
     * List of registries.
     */
    registries?: pulumi.Input<pulumi.Input<string>[]>;
    registry?: pulumi.Input<string>;
    requiredLabels?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyRequiredLabel>[]>;
    requiredLabelsEnabled?: pulumi.Input<boolean>;
    scanNfsMounts?: pulumi.Input<boolean>;
    /**
     * Indicates if scan should include sensitive data in the image.
     */
    scanSensitiveData?: pulumi.Input<boolean>;
    /**
     * Indicates if scanning should include scap.
     */
    scapEnabled?: pulumi.Input<boolean>;
    /**
     * List of SCAP user scripts for checks.
     */
    scapFiles?: pulumi.Input<pulumi.Input<string>[]>;
    scopes?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyScope>[]>;
    /**
     * List of trusted images.
     */
    trustedBaseImages?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyTrustedBaseImage>[]>;
    /**
     * Indicates if list of trusted base images is relevant.
     */
    trustedBaseImagesEnabled?: pulumi.Input<boolean>;
    /**
     * List of whitelisted licenses.
     */
    whitelistedLicenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates if license blacklist is relevant.
     */
    whitelistedLicensesEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ImageAssurancePolicy resource.
 */
export interface ImageAssurancePolicyArgs {
    /**
     * List of explicitly allowed images.
     */
    allowedImages?: pulumi.Input<pulumi.Input<string>[]>;
    applicationScopes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates if auditing for failures.
     */
    auditOnFailure?: pulumi.Input<boolean>;
    autoScanConfigured?: pulumi.Input<boolean>;
    autoScanEnabled?: pulumi.Input<boolean>;
    autoScanTimes?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyAutoScanTime>[]>;
    /**
     * List of function's forbidden permissions.
     */
    blacklistPermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates if blacklist permissions is relevant.
     */
    blacklistPermissionsEnabled?: pulumi.Input<boolean>;
    /**
     * List of blacklisted licenses.
     */
    blacklistedLicenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Lndicates if license blacklist is relevant.
     */
    blacklistedLicensesEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates if failed images are blocked.
     */
    blockFailed?: pulumi.Input<boolean>;
    controlExcludeNoFix?: pulumi.Input<boolean>;
    /**
     * List of Custom user scripts for checks.
     */
    customChecks?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyCustomCheck>[]>;
    /**
     * Indicates if scanning should include custom checks.
     */
    customChecksEnabled?: pulumi.Input<boolean>;
    customSeverityEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates if cves blacklist is relevant.
     */
    cvesBlackListEnabled?: pulumi.Input<boolean>;
    /**
     * List of cves blacklisted items.
     */
    cvesBlackLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates if cves whitelist is relevant.
     */
    cvesWhiteListEnabled?: pulumi.Input<boolean>;
    /**
     * List of cves whitelisted licenses
     */
    cvesWhiteLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Identifier of the cvss severity.
     */
    cvssSeverity?: pulumi.Input<string>;
    /**
     * Indicates if the cvss severity is scanned.
     */
    cvssSeverityEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates that policy should ignore cvss cases that do not have a known fix.
     */
    cvssSeverityExcludeNoFix?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    /**
     * Indicates if malware should block the image.
     */
    disallowMalware?: pulumi.Input<boolean>;
    dockerCisEnabled?: pulumi.Input<boolean>;
    /**
     * Name of the container image.
     */
    domain?: pulumi.Input<string>;
    domainName?: pulumi.Input<string>;
    dtaEnabled?: pulumi.Input<boolean>;
    dtaSeverity?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    enforce?: pulumi.Input<boolean>;
    enforceAfterDays?: pulumi.Input<number>;
    enforceExcessivePermissions?: pulumi.Input<boolean>;
    exceptionalMonitoredMalwarePaths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates if cicd failures will fail the image.
     */
    failCicd?: pulumi.Input<boolean>;
    forbiddenLabels?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyForbiddenLabel>[]>;
    forbiddenLabelsEnabled?: pulumi.Input<boolean>;
    forceMicroenforcer?: pulumi.Input<boolean>;
    functionIntegrityEnabled?: pulumi.Input<boolean>;
    ignoreRecentlyPublishedVln?: pulumi.Input<boolean>;
    /**
     * Indicates if risk resources are ignored.
     */
    ignoreRiskResourcesEnabled?: pulumi.Input<boolean>;
    /**
     * List of ignored risk resources.
     */
    ignoredRiskResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of images.
     */
    images?: pulumi.Input<pulumi.Input<string>[]>;
    kubeCisEnabled?: pulumi.Input<boolean>;
    /**
     * List of labels.
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    malwareAction?: pulumi.Input<string>;
    /**
     * Value of allowed maximum score.
     */
    maximumScore?: pulumi.Input<number>;
    /**
     * Indicates if exceeding the maximum score is scanned.
     */
    maximumScoreEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates that policy should ignore cases that do not have a known fix.
     */
    maximumScoreExcludeNoFix?: pulumi.Input<boolean>;
    monitoredMalwarePaths?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    /**
     * Indicates if raise a warning for images that should only be run as root.
     */
    onlyNoneRootUsers?: pulumi.Input<boolean>;
    /**
     * Indicates if packages blacklist is relevant.
     */
    packagesBlackListEnabled?: pulumi.Input<boolean>;
    /**
     * List of backlisted images.
     */
    packagesBlackLists?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyPackagesBlackList>[]>;
    /**
     * Indicates if packages whitelist is relevant.
     */
    packagesWhiteListEnabled?: pulumi.Input<boolean>;
    /**
     * List of whitelisted images.
     */
    packagesWhiteLists?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyPackagesWhiteList>[]>;
    partialResultsImageFail?: pulumi.Input<boolean>;
    readOnly?: pulumi.Input<boolean>;
    /**
     * List of registries.
     */
    registries?: pulumi.Input<pulumi.Input<string>[]>;
    registry?: pulumi.Input<string>;
    requiredLabels?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyRequiredLabel>[]>;
    requiredLabelsEnabled?: pulumi.Input<boolean>;
    scanNfsMounts?: pulumi.Input<boolean>;
    /**
     * Indicates if scan should include sensitive data in the image.
     */
    scanSensitiveData?: pulumi.Input<boolean>;
    /**
     * Indicates if scanning should include scap.
     */
    scapEnabled?: pulumi.Input<boolean>;
    /**
     * List of SCAP user scripts for checks.
     */
    scapFiles?: pulumi.Input<pulumi.Input<string>[]>;
    scopes?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyScope>[]>;
    /**
     * List of trusted images.
     */
    trustedBaseImages?: pulumi.Input<pulumi.Input<inputs.ImageAssurancePolicyTrustedBaseImage>[]>;
    /**
     * Indicates if list of trusted base images is relevant.
     */
    trustedBaseImagesEnabled?: pulumi.Input<boolean>;
    /**
     * List of whitelisted licenses.
     */
    whitelistedLicenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates if license blacklist is relevant.
     */
    whitelistedLicensesEnabled?: pulumi.Input<boolean>;
}
