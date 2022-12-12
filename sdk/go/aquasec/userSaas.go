// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `UserSaas` resource manages your saas users within Aqua.
//
// The users created must have at least one Csp Role that is already present within Aqua.
//
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
//			_, err := aquasec.NewUserSaas(ctx, "iaC1", &aquasec.UserSaasArgs{
//				AccountAdmin: pulumi.Bool(true),
//				CspRoles:     pulumi.StringArray{},
//				Email:        pulumi.String("infrastructure1@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = aquasec.NewUserSaas(ctx, "iaC2", &aquasec.UserSaasArgs{
//				AccountAdmin: pulumi.Bool(false),
//				CspRoles: pulumi.StringArray{
//					pulumi.String("Default"),
//				},
//				Email: pulumi.String("infrastructure2@example.com"),
//				Groups: aquasec.UserSaasGroupArray{
//					&aquasec.UserSaasGroupArgs{
//						GroupAdmin: pulumi.Bool(false),
//						Name:       pulumi.String("IacGroupName"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type UserSaas struct {
	pulumi.CustomResourceState

	AccountAdmin      pulumi.BoolOutput        `pulumi:"accountAdmin"`
	Confirmed         pulumi.BoolOutput        `pulumi:"confirmed"`
	Created           pulumi.StringOutput      `pulumi:"created"`
	CspRoles          pulumi.StringArrayOutput `pulumi:"cspRoles"`
	Email             pulumi.StringOutput      `pulumi:"email"`
	Groups            UserSaasGroupArrayOutput `pulumi:"groups"`
	Logins            UserSaasLoginArrayOutput `pulumi:"logins"`
	Multiaccount      pulumi.BoolOutput        `pulumi:"multiaccount"`
	PasswordReset     pulumi.BoolOutput        `pulumi:"passwordReset"`
	SendAnnouncements pulumi.BoolOutput        `pulumi:"sendAnnouncements"`
	SendNewPlugins    pulumi.BoolOutput        `pulumi:"sendNewPlugins"`
	SendNewRisks      pulumi.BoolOutput        `pulumi:"sendNewRisks"`
	SendScanResults   pulumi.BoolOutput        `pulumi:"sendScanResults"`
	UserId            pulumi.StringOutput      `pulumi:"userId"`
}

// NewUserSaas registers a new resource with the given unique name, arguments, and options.
func NewUserSaas(ctx *pulumi.Context,
	name string, args *UserSaasArgs, opts ...pulumi.ResourceOption) (*UserSaas, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountAdmin == nil {
		return nil, errors.New("invalid value for required argument 'AccountAdmin'")
	}
	if args.CspRoles == nil {
		return nil, errors.New("invalid value for required argument 'CspRoles'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UserSaas
	err := ctx.RegisterResource("aquasec:index/userSaas:UserSaas", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserSaas gets an existing UserSaas resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserSaas(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserSaasState, opts ...pulumi.ResourceOption) (*UserSaas, error) {
	var resource UserSaas
	err := ctx.ReadResource("aquasec:index/userSaas:UserSaas", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserSaas resources.
type userSaasState struct {
	AccountAdmin      *bool           `pulumi:"accountAdmin"`
	Confirmed         *bool           `pulumi:"confirmed"`
	Created           *string         `pulumi:"created"`
	CspRoles          []string        `pulumi:"cspRoles"`
	Email             *string         `pulumi:"email"`
	Groups            []UserSaasGroup `pulumi:"groups"`
	Logins            []UserSaasLogin `pulumi:"logins"`
	Multiaccount      *bool           `pulumi:"multiaccount"`
	PasswordReset     *bool           `pulumi:"passwordReset"`
	SendAnnouncements *bool           `pulumi:"sendAnnouncements"`
	SendNewPlugins    *bool           `pulumi:"sendNewPlugins"`
	SendNewRisks      *bool           `pulumi:"sendNewRisks"`
	SendScanResults   *bool           `pulumi:"sendScanResults"`
	UserId            *string         `pulumi:"userId"`
}

type UserSaasState struct {
	AccountAdmin      pulumi.BoolPtrInput
	Confirmed         pulumi.BoolPtrInput
	Created           pulumi.StringPtrInput
	CspRoles          pulumi.StringArrayInput
	Email             pulumi.StringPtrInput
	Groups            UserSaasGroupArrayInput
	Logins            UserSaasLoginArrayInput
	Multiaccount      pulumi.BoolPtrInput
	PasswordReset     pulumi.BoolPtrInput
	SendAnnouncements pulumi.BoolPtrInput
	SendNewPlugins    pulumi.BoolPtrInput
	SendNewRisks      pulumi.BoolPtrInput
	SendScanResults   pulumi.BoolPtrInput
	UserId            pulumi.StringPtrInput
}

func (UserSaasState) ElementType() reflect.Type {
	return reflect.TypeOf((*userSaasState)(nil)).Elem()
}

type userSaasArgs struct {
	AccountAdmin bool            `pulumi:"accountAdmin"`
	CspRoles     []string        `pulumi:"cspRoles"`
	Email        string          `pulumi:"email"`
	Groups       []UserSaasGroup `pulumi:"groups"`
}

// The set of arguments for constructing a UserSaas resource.
type UserSaasArgs struct {
	AccountAdmin pulumi.BoolInput
	CspRoles     pulumi.StringArrayInput
	Email        pulumi.StringInput
	Groups       UserSaasGroupArrayInput
}

func (UserSaasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userSaasArgs)(nil)).Elem()
}

type UserSaasInput interface {
	pulumi.Input

	ToUserSaasOutput() UserSaasOutput
	ToUserSaasOutputWithContext(ctx context.Context) UserSaasOutput
}

func (*UserSaas) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSaas)(nil)).Elem()
}

func (i *UserSaas) ToUserSaasOutput() UserSaasOutput {
	return i.ToUserSaasOutputWithContext(context.Background())
}

func (i *UserSaas) ToUserSaasOutputWithContext(ctx context.Context) UserSaasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSaasOutput)
}

// UserSaasArrayInput is an input type that accepts UserSaasArray and UserSaasArrayOutput values.
// You can construct a concrete instance of `UserSaasArrayInput` via:
//
//	UserSaasArray{ UserSaasArgs{...} }
type UserSaasArrayInput interface {
	pulumi.Input

	ToUserSaasArrayOutput() UserSaasArrayOutput
	ToUserSaasArrayOutputWithContext(context.Context) UserSaasArrayOutput
}

type UserSaasArray []UserSaasInput

func (UserSaasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserSaas)(nil)).Elem()
}

func (i UserSaasArray) ToUserSaasArrayOutput() UserSaasArrayOutput {
	return i.ToUserSaasArrayOutputWithContext(context.Background())
}

func (i UserSaasArray) ToUserSaasArrayOutputWithContext(ctx context.Context) UserSaasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSaasArrayOutput)
}

// UserSaasMapInput is an input type that accepts UserSaasMap and UserSaasMapOutput values.
// You can construct a concrete instance of `UserSaasMapInput` via:
//
//	UserSaasMap{ "key": UserSaasArgs{...} }
type UserSaasMapInput interface {
	pulumi.Input

	ToUserSaasMapOutput() UserSaasMapOutput
	ToUserSaasMapOutputWithContext(context.Context) UserSaasMapOutput
}

type UserSaasMap map[string]UserSaasInput

func (UserSaasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserSaas)(nil)).Elem()
}

func (i UserSaasMap) ToUserSaasMapOutput() UserSaasMapOutput {
	return i.ToUserSaasMapOutputWithContext(context.Background())
}

func (i UserSaasMap) ToUserSaasMapOutputWithContext(ctx context.Context) UserSaasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSaasMapOutput)
}

type UserSaasOutput struct{ *pulumi.OutputState }

func (UserSaasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSaas)(nil)).Elem()
}

func (o UserSaasOutput) ToUserSaasOutput() UserSaasOutput {
	return o
}

func (o UserSaasOutput) ToUserSaasOutputWithContext(ctx context.Context) UserSaasOutput {
	return o
}

func (o UserSaasOutput) AccountAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserSaas) pulumi.BoolOutput { return v.AccountAdmin }).(pulumi.BoolOutput)
}

func (o UserSaasOutput) Confirmed() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserSaas) pulumi.BoolOutput { return v.Confirmed }).(pulumi.BoolOutput)
}

func (o UserSaasOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *UserSaas) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

func (o UserSaasOutput) CspRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserSaas) pulumi.StringArrayOutput { return v.CspRoles }).(pulumi.StringArrayOutput)
}

func (o UserSaasOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *UserSaas) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

func (o UserSaasOutput) Groups() UserSaasGroupArrayOutput {
	return o.ApplyT(func(v *UserSaas) UserSaasGroupArrayOutput { return v.Groups }).(UserSaasGroupArrayOutput)
}

func (o UserSaasOutput) Logins() UserSaasLoginArrayOutput {
	return o.ApplyT(func(v *UserSaas) UserSaasLoginArrayOutput { return v.Logins }).(UserSaasLoginArrayOutput)
}

func (o UserSaasOutput) Multiaccount() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserSaas) pulumi.BoolOutput { return v.Multiaccount }).(pulumi.BoolOutput)
}

func (o UserSaasOutput) PasswordReset() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserSaas) pulumi.BoolOutput { return v.PasswordReset }).(pulumi.BoolOutput)
}

func (o UserSaasOutput) SendAnnouncements() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserSaas) pulumi.BoolOutput { return v.SendAnnouncements }).(pulumi.BoolOutput)
}

func (o UserSaasOutput) SendNewPlugins() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserSaas) pulumi.BoolOutput { return v.SendNewPlugins }).(pulumi.BoolOutput)
}

func (o UserSaasOutput) SendNewRisks() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserSaas) pulumi.BoolOutput { return v.SendNewRisks }).(pulumi.BoolOutput)
}

func (o UserSaasOutput) SendScanResults() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserSaas) pulumi.BoolOutput { return v.SendScanResults }).(pulumi.BoolOutput)
}

func (o UserSaasOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserSaas) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type UserSaasArrayOutput struct{ *pulumi.OutputState }

func (UserSaasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserSaas)(nil)).Elem()
}

func (o UserSaasArrayOutput) ToUserSaasArrayOutput() UserSaasArrayOutput {
	return o
}

func (o UserSaasArrayOutput) ToUserSaasArrayOutputWithContext(ctx context.Context) UserSaasArrayOutput {
	return o
}

func (o UserSaasArrayOutput) Index(i pulumi.IntInput) UserSaasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserSaas {
		return vs[0].([]*UserSaas)[vs[1].(int)]
	}).(UserSaasOutput)
}

type UserSaasMapOutput struct{ *pulumi.OutputState }

func (UserSaasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserSaas)(nil)).Elem()
}

func (o UserSaasMapOutput) ToUserSaasMapOutput() UserSaasMapOutput {
	return o
}

func (o UserSaasMapOutput) ToUserSaasMapOutputWithContext(ctx context.Context) UserSaasMapOutput {
	return o
}

func (o UserSaasMapOutput) MapIndex(k pulumi.StringInput) UserSaasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserSaas {
		return vs[0].(map[string]*UserSaas)[vs[1].(string)]
	}).(UserSaasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserSaasInput)(nil)).Elem(), &UserSaas{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserSaasArrayInput)(nil)).Elem(), UserSaasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserSaasMapInput)(nil)).Elem(), UserSaasMap{})
	pulumi.RegisterOutputType(UserSaasOutput{})
	pulumi.RegisterOutputType(UserSaasArrayOutput{})
	pulumi.RegisterOutputType(UserSaasMapOutput{})
}
