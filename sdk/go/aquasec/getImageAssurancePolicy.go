// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-aquasec/sdk/go/aquasec/internal"
)

func LookupImageAssurancePolicy(ctx *pulumi.Context, args *LookupImageAssurancePolicyArgs, opts ...pulumi.InvokeOption) (*LookupImageAssurancePolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupImageAssurancePolicyResult
	err := ctx.Invoke("aquasec:index/getImageAssurancePolicy:getImageAssurancePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImageAssurancePolicy.
type LookupImageAssurancePolicyArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getImageAssurancePolicy.
type LookupImageAssurancePolicyResult struct {
	// List of explicitly allowed images.
	AllowedImages     []string `pulumi:"allowedImages"`
	ApplicationScopes []string `pulumi:"applicationScopes"`
	// Indicates if auditing for failures.
	AuditOnFailure bool `pulumi:"auditOnFailure"`
	// Name of user account that created the policy.
	Author             string                                `pulumi:"author"`
	AutoScanConfigured bool                                  `pulumi:"autoScanConfigured"`
	AutoScanEnabled    bool                                  `pulumi:"autoScanEnabled"`
	AutoScanTimes      []GetImageAssurancePolicyAutoScanTime `pulumi:"autoScanTimes"`
	// List of function's forbidden permissions.
	BlacklistPermissions []string `pulumi:"blacklistPermissions"`
	// Indicates if blacklist permissions is relevant.
	BlacklistPermissionsEnabled bool `pulumi:"blacklistPermissionsEnabled"`
	// List of blacklisted licenses.
	BlacklistedLicenses []string `pulumi:"blacklistedLicenses"`
	// Indicates if license blacklist is relevant.
	BlacklistedLicensesEnabled bool `pulumi:"blacklistedLicensesEnabled"`
	// Indicates if failed images are blocked.
	BlockFailed         bool `pulumi:"blockFailed"`
	ControlExcludeNoFix bool `pulumi:"controlExcludeNoFix"`
	// List of Custom user scripts for checks.
	CustomChecks []GetImageAssurancePolicyCustomCheck `pulumi:"customChecks"`
	// Indicates if scanning should include custom checks.
	CustomChecksEnabled   bool `pulumi:"customChecksEnabled"`
	CustomSeverityEnabled bool `pulumi:"customSeverityEnabled"`
	// Indicates if CVEs blacklist is relevant.
	CvesBlackListEnabled bool `pulumi:"cvesBlackListEnabled"`
	// List of CVEs blacklisted items.
	CvesBlackLists []string `pulumi:"cvesBlackLists"`
	// Indicates if CVEs whitelist is relevant.
	CvesWhiteListEnabled bool `pulumi:"cvesWhiteListEnabled"`
	// List of cves whitelisted licenses
	CvesWhiteLists []string `pulumi:"cvesWhiteLists"`
	// Identifier of the cvss severity.
	CvssSeverity string `pulumi:"cvssSeverity"`
	// Indicates if the cvss severity is scanned.
	CvssSeverityEnabled bool `pulumi:"cvssSeverityEnabled"`
	// Indicates that policy should ignore cvss cases that do not have a known fix.
	CvssSeverityExcludeNoFix bool   `pulumi:"cvssSeverityExcludeNoFix"`
	Description              string `pulumi:"description"`
	// Indicates if malware should block the image.
	DisallowMalware bool `pulumi:"disallowMalware"`
	// Checks the host according to the Docker CIS benchmark, if Docker is found on the host.
	DockerCisEnabled bool `pulumi:"dockerCisEnabled"`
	// Name of the container image.
	Domain                           string   `pulumi:"domain"`
	DomainName                       string   `pulumi:"domainName"`
	DtaEnabled                       bool     `pulumi:"dtaEnabled"`
	DtaSeverity                      string   `pulumi:"dtaSeverity"`
	Enabled                          bool     `pulumi:"enabled"`
	Enforce                          bool     `pulumi:"enforce"`
	EnforceAfterDays                 int      `pulumi:"enforceAfterDays"`
	EnforceExcessivePermissions      bool     `pulumi:"enforceExcessivePermissions"`
	ExceptionalMonitoredMalwarePaths []string `pulumi:"exceptionalMonitoredMalwarePaths"`
	// Indicates if cicd failures will fail the image.
	FailCicd                 bool                                    `pulumi:"failCicd"`
	ForbiddenLabels          []GetImageAssurancePolicyForbiddenLabel `pulumi:"forbiddenLabels"`
	ForbiddenLabelsEnabled   bool                                    `pulumi:"forbiddenLabelsEnabled"`
	ForceMicroenforcer       bool                                    `pulumi:"forceMicroenforcer"`
	FunctionIntegrityEnabled bool                                    `pulumi:"functionIntegrityEnabled"`
	// The ID of this resource.
	Id                               string `pulumi:"id"`
	IgnoreRecentlyPublishedVln       bool   `pulumi:"ignoreRecentlyPublishedVln"`
	IgnoreRecentlyPublishedVlnPeriod int    `pulumi:"ignoreRecentlyPublishedVlnPeriod"`
	// Indicates if risk resources are ignored.
	IgnoreRiskResourcesEnabled bool `pulumi:"ignoreRiskResourcesEnabled"`
	// List of ignored risk resources.
	IgnoredRiskResources []string `pulumi:"ignoredRiskResources"`
	// List of images.
	Images []string `pulumi:"images"`
	// Performs a Kubernetes CIS benchmark check for the host.
	KubeCisEnabled bool `pulumi:"kubeCisEnabled"`
	// List of labels.
	Labels        []string `pulumi:"labels"`
	MalwareAction string   `pulumi:"malwareAction"`
	// Value of allowed maximum score.
	MaximumScore float64 `pulumi:"maximumScore"`
	// Indicates if exceeding the maximum score is scanned.
	MaximumScoreEnabled bool `pulumi:"maximumScoreEnabled"`
	// Indicates that policy should ignore cases that do not have a known fix.
	MaximumScoreExcludeNoFix bool     `pulumi:"maximumScoreExcludeNoFix"`
	MonitoredMalwarePaths    []string `pulumi:"monitoredMalwarePaths"`
	Name                     string   `pulumi:"name"`
	// Indicates if raise a warning for images that should only be run as root.
	OnlyNoneRootUsers bool `pulumi:"onlyNoneRootUsers"`
	// Indicates if packages blacklist is relevant.
	PackagesBlackListEnabled bool `pulumi:"packagesBlackListEnabled"`
	// List of blacklisted images.
	PackagesBlackLists []GetImageAssurancePolicyPackagesBlackList `pulumi:"packagesBlackLists"`
	// Indicates if packages whitelist is relevant.
	PackagesWhiteListEnabled bool `pulumi:"packagesWhiteListEnabled"`
	// List of whitelisted images.
	PackagesWhiteLists      []GetImageAssurancePolicyPackagesWhiteList `pulumi:"packagesWhiteLists"`
	PartialResultsImageFail bool                                       `pulumi:"partialResultsImageFail"`
	ReadOnly                bool                                       `pulumi:"readOnly"`
	// List of registries.
	Registries            []string                               `pulumi:"registries"`
	Registry              string                                 `pulumi:"registry"`
	RequiredLabels        []GetImageAssurancePolicyRequiredLabel `pulumi:"requiredLabels"`
	RequiredLabelsEnabled bool                                   `pulumi:"requiredLabelsEnabled"`
	ScanNfsMounts         bool                                   `pulumi:"scanNfsMounts"`
	// Indicates if scan should include sensitive data in the image.
	ScanSensitiveData bool `pulumi:"scanSensitiveData"`
	// Indicates if scanning should include scap.
	ScapEnabled bool `pulumi:"scapEnabled"`
	// List of SCAP user scripts for checks.
	ScapFiles []string                       `pulumi:"scapFiles"`
	Scopes    []GetImageAssurancePolicyScope `pulumi:"scopes"`
	// List of trusted images.
	TrustedBaseImages []GetImageAssurancePolicyTrustedBaseImage `pulumi:"trustedBaseImages"`
	// Indicates if list of trusted base images is relevant.
	TrustedBaseImagesEnabled bool `pulumi:"trustedBaseImagesEnabled"`
	// List of whitelisted licenses.
	WhitelistedLicenses []string `pulumi:"whitelistedLicenses"`
	// Indicates if license blacklist is relevant.
	WhitelistedLicensesEnabled bool `pulumi:"whitelistedLicensesEnabled"`
}

func LookupImageAssurancePolicyOutput(ctx *pulumi.Context, args LookupImageAssurancePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupImageAssurancePolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImageAssurancePolicyResult, error) {
			args := v.(LookupImageAssurancePolicyArgs)
			r, err := LookupImageAssurancePolicy(ctx, &args, opts...)
			var s LookupImageAssurancePolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImageAssurancePolicyResultOutput)
}

// A collection of arguments for invoking getImageAssurancePolicy.
type LookupImageAssurancePolicyOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupImageAssurancePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageAssurancePolicyArgs)(nil)).Elem()
}

// A collection of values returned by getImageAssurancePolicy.
type LookupImageAssurancePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupImageAssurancePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageAssurancePolicyResult)(nil)).Elem()
}

func (o LookupImageAssurancePolicyResultOutput) ToLookupImageAssurancePolicyResultOutput() LookupImageAssurancePolicyResultOutput {
	return o
}

func (o LookupImageAssurancePolicyResultOutput) ToLookupImageAssurancePolicyResultOutputWithContext(ctx context.Context) LookupImageAssurancePolicyResultOutput {
	return o
}

// List of explicitly allowed images.
func (o LookupImageAssurancePolicyResultOutput) AllowedImages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.AllowedImages }).(pulumi.StringArrayOutput)
}

func (o LookupImageAssurancePolicyResultOutput) ApplicationScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.ApplicationScopes }).(pulumi.StringArrayOutput)
}

// Indicates if auditing for failures.
func (o LookupImageAssurancePolicyResultOutput) AuditOnFailure() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.AuditOnFailure }).(pulumi.BoolOutput)
}

// Name of user account that created the policy.
func (o LookupImageAssurancePolicyResultOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) string { return v.Author }).(pulumi.StringOutput)
}

func (o LookupImageAssurancePolicyResultOutput) AutoScanConfigured() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.AutoScanConfigured }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) AutoScanEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.AutoScanEnabled }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) AutoScanTimes() GetImageAssurancePolicyAutoScanTimeArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []GetImageAssurancePolicyAutoScanTime { return v.AutoScanTimes }).(GetImageAssurancePolicyAutoScanTimeArrayOutput)
}

// List of function's forbidden permissions.
func (o LookupImageAssurancePolicyResultOutput) BlacklistPermissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.BlacklistPermissions }).(pulumi.StringArrayOutput)
}

// Indicates if blacklist permissions is relevant.
func (o LookupImageAssurancePolicyResultOutput) BlacklistPermissionsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.BlacklistPermissionsEnabled }).(pulumi.BoolOutput)
}

// List of blacklisted licenses.
func (o LookupImageAssurancePolicyResultOutput) BlacklistedLicenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.BlacklistedLicenses }).(pulumi.StringArrayOutput)
}

// Indicates if license blacklist is relevant.
func (o LookupImageAssurancePolicyResultOutput) BlacklistedLicensesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.BlacklistedLicensesEnabled }).(pulumi.BoolOutput)
}

// Indicates if failed images are blocked.
func (o LookupImageAssurancePolicyResultOutput) BlockFailed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.BlockFailed }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) ControlExcludeNoFix() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.ControlExcludeNoFix }).(pulumi.BoolOutput)
}

// List of Custom user scripts for checks.
func (o LookupImageAssurancePolicyResultOutput) CustomChecks() GetImageAssurancePolicyCustomCheckArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []GetImageAssurancePolicyCustomCheck { return v.CustomChecks }).(GetImageAssurancePolicyCustomCheckArrayOutput)
}

// Indicates if scanning should include custom checks.
func (o LookupImageAssurancePolicyResultOutput) CustomChecksEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.CustomChecksEnabled }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) CustomSeverityEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.CustomSeverityEnabled }).(pulumi.BoolOutput)
}

// Indicates if CVEs blacklist is relevant.
func (o LookupImageAssurancePolicyResultOutput) CvesBlackListEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.CvesBlackListEnabled }).(pulumi.BoolOutput)
}

// List of CVEs blacklisted items.
func (o LookupImageAssurancePolicyResultOutput) CvesBlackLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.CvesBlackLists }).(pulumi.StringArrayOutput)
}

// Indicates if CVEs whitelist is relevant.
func (o LookupImageAssurancePolicyResultOutput) CvesWhiteListEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.CvesWhiteListEnabled }).(pulumi.BoolOutput)
}

// List of cves whitelisted licenses
func (o LookupImageAssurancePolicyResultOutput) CvesWhiteLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.CvesWhiteLists }).(pulumi.StringArrayOutput)
}

// Identifier of the cvss severity.
func (o LookupImageAssurancePolicyResultOutput) CvssSeverity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) string { return v.CvssSeverity }).(pulumi.StringOutput)
}

// Indicates if the cvss severity is scanned.
func (o LookupImageAssurancePolicyResultOutput) CvssSeverityEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.CvssSeverityEnabled }).(pulumi.BoolOutput)
}

// Indicates that policy should ignore cvss cases that do not have a known fix.
func (o LookupImageAssurancePolicyResultOutput) CvssSeverityExcludeNoFix() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.CvssSeverityExcludeNoFix }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Indicates if malware should block the image.
func (o LookupImageAssurancePolicyResultOutput) DisallowMalware() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.DisallowMalware }).(pulumi.BoolOutput)
}

// Checks the host according to the Docker CIS benchmark, if Docker is found on the host.
func (o LookupImageAssurancePolicyResultOutput) DockerCisEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.DockerCisEnabled }).(pulumi.BoolOutput)
}

// Name of the container image.
func (o LookupImageAssurancePolicyResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o LookupImageAssurancePolicyResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o LookupImageAssurancePolicyResultOutput) DtaEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.DtaEnabled }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) DtaSeverity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) string { return v.DtaSeverity }).(pulumi.StringOutput)
}

func (o LookupImageAssurancePolicyResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) Enforce() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.Enforce }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) EnforceAfterDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) int { return v.EnforceAfterDays }).(pulumi.IntOutput)
}

func (o LookupImageAssurancePolicyResultOutput) EnforceExcessivePermissions() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.EnforceExcessivePermissions }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) ExceptionalMonitoredMalwarePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.ExceptionalMonitoredMalwarePaths }).(pulumi.StringArrayOutput)
}

// Indicates if cicd failures will fail the image.
func (o LookupImageAssurancePolicyResultOutput) FailCicd() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.FailCicd }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) ForbiddenLabels() GetImageAssurancePolicyForbiddenLabelArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []GetImageAssurancePolicyForbiddenLabel {
		return v.ForbiddenLabels
	}).(GetImageAssurancePolicyForbiddenLabelArrayOutput)
}

func (o LookupImageAssurancePolicyResultOutput) ForbiddenLabelsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.ForbiddenLabelsEnabled }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) ForceMicroenforcer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.ForceMicroenforcer }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) FunctionIntegrityEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.FunctionIntegrityEnabled }).(pulumi.BoolOutput)
}

// The ID of this resource.
func (o LookupImageAssurancePolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupImageAssurancePolicyResultOutput) IgnoreRecentlyPublishedVln() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.IgnoreRecentlyPublishedVln }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) IgnoreRecentlyPublishedVlnPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) int { return v.IgnoreRecentlyPublishedVlnPeriod }).(pulumi.IntOutput)
}

// Indicates if risk resources are ignored.
func (o LookupImageAssurancePolicyResultOutput) IgnoreRiskResourcesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.IgnoreRiskResourcesEnabled }).(pulumi.BoolOutput)
}

// List of ignored risk resources.
func (o LookupImageAssurancePolicyResultOutput) IgnoredRiskResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.IgnoredRiskResources }).(pulumi.StringArrayOutput)
}

// List of images.
func (o LookupImageAssurancePolicyResultOutput) Images() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.Images }).(pulumi.StringArrayOutput)
}

// Performs a Kubernetes CIS benchmark check for the host.
func (o LookupImageAssurancePolicyResultOutput) KubeCisEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.KubeCisEnabled }).(pulumi.BoolOutput)
}

// List of labels.
func (o LookupImageAssurancePolicyResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o LookupImageAssurancePolicyResultOutput) MalwareAction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) string { return v.MalwareAction }).(pulumi.StringOutput)
}

// Value of allowed maximum score.
func (o LookupImageAssurancePolicyResultOutput) MaximumScore() pulumi.Float64Output {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) float64 { return v.MaximumScore }).(pulumi.Float64Output)
}

// Indicates if exceeding the maximum score is scanned.
func (o LookupImageAssurancePolicyResultOutput) MaximumScoreEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.MaximumScoreEnabled }).(pulumi.BoolOutput)
}

// Indicates that policy should ignore cases that do not have a known fix.
func (o LookupImageAssurancePolicyResultOutput) MaximumScoreExcludeNoFix() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.MaximumScoreExcludeNoFix }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) MonitoredMalwarePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.MonitoredMalwarePaths }).(pulumi.StringArrayOutput)
}

func (o LookupImageAssurancePolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates if raise a warning for images that should only be run as root.
func (o LookupImageAssurancePolicyResultOutput) OnlyNoneRootUsers() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.OnlyNoneRootUsers }).(pulumi.BoolOutput)
}

// Indicates if packages blacklist is relevant.
func (o LookupImageAssurancePolicyResultOutput) PackagesBlackListEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.PackagesBlackListEnabled }).(pulumi.BoolOutput)
}

// List of blacklisted images.
func (o LookupImageAssurancePolicyResultOutput) PackagesBlackLists() GetImageAssurancePolicyPackagesBlackListArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []GetImageAssurancePolicyPackagesBlackList {
		return v.PackagesBlackLists
	}).(GetImageAssurancePolicyPackagesBlackListArrayOutput)
}

// Indicates if packages whitelist is relevant.
func (o LookupImageAssurancePolicyResultOutput) PackagesWhiteListEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.PackagesWhiteListEnabled }).(pulumi.BoolOutput)
}

// List of whitelisted images.
func (o LookupImageAssurancePolicyResultOutput) PackagesWhiteLists() GetImageAssurancePolicyPackagesWhiteListArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []GetImageAssurancePolicyPackagesWhiteList {
		return v.PackagesWhiteLists
	}).(GetImageAssurancePolicyPackagesWhiteListArrayOutput)
}

func (o LookupImageAssurancePolicyResultOutput) PartialResultsImageFail() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.PartialResultsImageFail }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.ReadOnly }).(pulumi.BoolOutput)
}

// List of registries.
func (o LookupImageAssurancePolicyResultOutput) Registries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.Registries }).(pulumi.StringArrayOutput)
}

func (o LookupImageAssurancePolicyResultOutput) Registry() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) string { return v.Registry }).(pulumi.StringOutput)
}

func (o LookupImageAssurancePolicyResultOutput) RequiredLabels() GetImageAssurancePolicyRequiredLabelArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []GetImageAssurancePolicyRequiredLabel {
		return v.RequiredLabels
	}).(GetImageAssurancePolicyRequiredLabelArrayOutput)
}

func (o LookupImageAssurancePolicyResultOutput) RequiredLabelsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.RequiredLabelsEnabled }).(pulumi.BoolOutput)
}

func (o LookupImageAssurancePolicyResultOutput) ScanNfsMounts() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.ScanNfsMounts }).(pulumi.BoolOutput)
}

// Indicates if scan should include sensitive data in the image.
func (o LookupImageAssurancePolicyResultOutput) ScanSensitiveData() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.ScanSensitiveData }).(pulumi.BoolOutput)
}

// Indicates if scanning should include scap.
func (o LookupImageAssurancePolicyResultOutput) ScapEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.ScapEnabled }).(pulumi.BoolOutput)
}

// List of SCAP user scripts for checks.
func (o LookupImageAssurancePolicyResultOutput) ScapFiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.ScapFiles }).(pulumi.StringArrayOutput)
}

func (o LookupImageAssurancePolicyResultOutput) Scopes() GetImageAssurancePolicyScopeArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []GetImageAssurancePolicyScope { return v.Scopes }).(GetImageAssurancePolicyScopeArrayOutput)
}

// List of trusted images.
func (o LookupImageAssurancePolicyResultOutput) TrustedBaseImages() GetImageAssurancePolicyTrustedBaseImageArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []GetImageAssurancePolicyTrustedBaseImage {
		return v.TrustedBaseImages
	}).(GetImageAssurancePolicyTrustedBaseImageArrayOutput)
}

// Indicates if list of trusted base images is relevant.
func (o LookupImageAssurancePolicyResultOutput) TrustedBaseImagesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.TrustedBaseImagesEnabled }).(pulumi.BoolOutput)
}

// List of whitelisted licenses.
func (o LookupImageAssurancePolicyResultOutput) WhitelistedLicenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) []string { return v.WhitelistedLicenses }).(pulumi.StringArrayOutput)
}

// Indicates if license blacklist is relevant.
func (o LookupImageAssurancePolicyResultOutput) WhitelistedLicensesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageAssurancePolicyResult) bool { return v.WhitelistedLicensesEnabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageAssurancePolicyResultOutput{})
}
