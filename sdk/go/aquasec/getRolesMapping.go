// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-aquasec/sdk/go/aquasec/internal"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-aquasec/sdk/go/aquasec"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			rolesMapping, err := aquasec.GetRolesMapping(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("roleMappingAll", rolesMapping)
//			ctx.Export("roleMappingSaml", rolesMapping.Samls)
//			return nil
//		})
//	}
//
// ```
func GetRolesMapping(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetRolesMappingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRolesMappingResult
	err := ctx.Invoke("aquasec:index/getRolesMapping:getRolesMapping", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getRolesMapping.
type GetRolesMappingResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// LDAP Authentication
	Ldaps []GetRolesMappingLdap `pulumi:"ldaps"`
	// Oauth2 Authentication
	Oauth2s []GetRolesMappingOauth2 `pulumi:"oauth2s"`
	// OpenId Authentication
	Openids []GetRolesMappingOpenid `pulumi:"openids"`
	// SAML Authentication
	Samls []GetRolesMappingSaml `pulumi:"samls"`
}

func GetRolesMappingOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetRolesMappingResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetRolesMappingResult, error) {
		r, err := GetRolesMapping(ctx, opts...)
		var s GetRolesMappingResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetRolesMappingResultOutput)
}

// A collection of values returned by getRolesMapping.
type GetRolesMappingResultOutput struct{ *pulumi.OutputState }

func (GetRolesMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolesMappingResult)(nil)).Elem()
}

func (o GetRolesMappingResultOutput) ToGetRolesMappingResultOutput() GetRolesMappingResultOutput {
	return o
}

func (o GetRolesMappingResultOutput) ToGetRolesMappingResultOutputWithContext(ctx context.Context) GetRolesMappingResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetRolesMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRolesMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// LDAP Authentication
func (o GetRolesMappingResultOutput) Ldaps() GetRolesMappingLdapArrayOutput {
	return o.ApplyT(func(v GetRolesMappingResult) []GetRolesMappingLdap { return v.Ldaps }).(GetRolesMappingLdapArrayOutput)
}

// Oauth2 Authentication
func (o GetRolesMappingResultOutput) Oauth2s() GetRolesMappingOauth2ArrayOutput {
	return o.ApplyT(func(v GetRolesMappingResult) []GetRolesMappingOauth2 { return v.Oauth2s }).(GetRolesMappingOauth2ArrayOutput)
}

// OpenId Authentication
func (o GetRolesMappingResultOutput) Openids() GetRolesMappingOpenidArrayOutput {
	return o.ApplyT(func(v GetRolesMappingResult) []GetRolesMappingOpenid { return v.Openids }).(GetRolesMappingOpenidArrayOutput)
}

// SAML Authentication
func (o GetRolesMappingResultOutput) Samls() GetRolesMappingSamlArrayOutput {
	return o.ApplyT(func(v GetRolesMappingResult) []GetRolesMappingSaml { return v.Samls }).(GetRolesMappingSamlArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRolesMappingResultOutput{})
}
