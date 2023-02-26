// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Service struct {
	pulumi.CustomResourceState

	// Indicates the application scope of the service.
	ApplicationScopes pulumi.StringArrayOutput `pulumi:"applicationScopes"`
	// Username of the account that created the service.
	Author pulumi.StringOutput `pulumi:"author"`
	// The number of containers associated with the service.
	ContainersCount pulumi.IntOutput `pulumi:"containersCount"`
	// A textual description of the service record; maximum 500 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Enforcement status of the service.
	Enforce pulumi.BoolPtrOutput `pulumi:"enforce"`
	// Whether the service has been evaluated for security vulnerabilities.
	Evaluated pulumi.BoolOutput `pulumi:"evaluated"`
	// Indicates if registered or not.
	IsRegistered pulumi.BoolOutput `pulumi:"isRegistered"`
	// Timestamp of the last update in Unix time format.
	Lastupdate pulumi.IntOutput `pulumi:"lastupdate"`
	// Indicates if monitoring is enabled or not
	Monitoring pulumi.BoolPtrOutput `pulumi:"monitoring"`
	// The name of the service. It is recommended not to use whitespace characters in the name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of container that are not evaluated.
	NotEvaluatedCount pulumi.IntOutput `pulumi:"notEvaluatedCount"`
	// The service's policies; an array of container firewall policy names.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// Rules priority, must be between 1-100.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression pulumi.StringPtrOutput `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables ServiceScopeVariableArrayOutput `pulumi:"scopeVariables"`
	// Type of the workload. container or host.
	Target pulumi.StringOutput `pulumi:"target"`
	// The number of containers allocated to the service that are not registered.
	UnregisteredCount pulumi.IntOutput `pulumi:"unregisteredCount"`
	// Number of high severity vulnerabilities.
	VulnerabilitiesHigh pulumi.IntOutput `pulumi:"vulnerabilitiesHigh"`
	// Number of low severity vulnerabilities.
	VulnerabilitiesLow pulumi.IntOutput `pulumi:"vulnerabilitiesLow"`
	// Number of malware.
	VulnerabilitiesMalware pulumi.IntOutput `pulumi:"vulnerabilitiesMalware"`
	// Number of medium severity vulnerabilities.
	VulnerabilitiesMedium pulumi.IntOutput `pulumi:"vulnerabilitiesMedium"`
	// Number of negligible vulnerabilities.
	VulnerabilitiesNegligible pulumi.IntOutput `pulumi:"vulnerabilitiesNegligible"`
	// The CVSS average vulnerabilities score.
	VulnerabilitiesScoreAverage pulumi.IntOutput `pulumi:"vulnerabilitiesScoreAverage"`
	// Number of sensitive vulnerabilities.
	VulnerabilitiesSensitive pulumi.IntOutput `pulumi:"vulnerabilitiesSensitive"`
	// Total number of vulnerabilities.
	VulnerabilitiesTotal pulumi.IntOutput `pulumi:"vulnerabilitiesTotal"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationScopes == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationScopes'")
	}
	if args.Policies == nil {
		return nil, errors.New("invalid value for required argument 'Policies'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("aquasec:index/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("aquasec:index/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// Indicates the application scope of the service.
	ApplicationScopes []string `pulumi:"applicationScopes"`
	// Username of the account that created the service.
	Author *string `pulumi:"author"`
	// The number of containers associated with the service.
	ContainersCount *int `pulumi:"containersCount"`
	// A textual description of the service record; maximum 500 characters.
	Description *string `pulumi:"description"`
	// Enforcement status of the service.
	Enforce *bool `pulumi:"enforce"`
	// Whether the service has been evaluated for security vulnerabilities.
	Evaluated *bool `pulumi:"evaluated"`
	// Indicates if registered or not.
	IsRegistered *bool `pulumi:"isRegistered"`
	// Timestamp of the last update in Unix time format.
	Lastupdate *int `pulumi:"lastupdate"`
	// Indicates if monitoring is enabled or not
	Monitoring *bool `pulumi:"monitoring"`
	// The name of the service. It is recommended not to use whitespace characters in the name.
	Name *string `pulumi:"name"`
	// The number of container that are not evaluated.
	NotEvaluatedCount *int `pulumi:"notEvaluatedCount"`
	// The service's policies; an array of container firewall policy names.
	Policies []string `pulumi:"policies"`
	// Rules priority, must be between 1-100.
	Priority *int `pulumi:"priority"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression *string `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables []ServiceScopeVariable `pulumi:"scopeVariables"`
	// Type of the workload. container or host.
	Target *string `pulumi:"target"`
	// The number of containers allocated to the service that are not registered.
	UnregisteredCount *int `pulumi:"unregisteredCount"`
	// Number of high severity vulnerabilities.
	VulnerabilitiesHigh *int `pulumi:"vulnerabilitiesHigh"`
	// Number of low severity vulnerabilities.
	VulnerabilitiesLow *int `pulumi:"vulnerabilitiesLow"`
	// Number of malware.
	VulnerabilitiesMalware *int `pulumi:"vulnerabilitiesMalware"`
	// Number of medium severity vulnerabilities.
	VulnerabilitiesMedium *int `pulumi:"vulnerabilitiesMedium"`
	// Number of negligible vulnerabilities.
	VulnerabilitiesNegligible *int `pulumi:"vulnerabilitiesNegligible"`
	// The CVSS average vulnerabilities score.
	VulnerabilitiesScoreAverage *int `pulumi:"vulnerabilitiesScoreAverage"`
	// Number of sensitive vulnerabilities.
	VulnerabilitiesSensitive *int `pulumi:"vulnerabilitiesSensitive"`
	// Total number of vulnerabilities.
	VulnerabilitiesTotal *int `pulumi:"vulnerabilitiesTotal"`
}

type ServiceState struct {
	// Indicates the application scope of the service.
	ApplicationScopes pulumi.StringArrayInput
	// Username of the account that created the service.
	Author pulumi.StringPtrInput
	// The number of containers associated with the service.
	ContainersCount pulumi.IntPtrInput
	// A textual description of the service record; maximum 500 characters.
	Description pulumi.StringPtrInput
	// Enforcement status of the service.
	Enforce pulumi.BoolPtrInput
	// Whether the service has been evaluated for security vulnerabilities.
	Evaluated pulumi.BoolPtrInput
	// Indicates if registered or not.
	IsRegistered pulumi.BoolPtrInput
	// Timestamp of the last update in Unix time format.
	Lastupdate pulumi.IntPtrInput
	// Indicates if monitoring is enabled or not
	Monitoring pulumi.BoolPtrInput
	// The name of the service. It is recommended not to use whitespace characters in the name.
	Name pulumi.StringPtrInput
	// The number of container that are not evaluated.
	NotEvaluatedCount pulumi.IntPtrInput
	// The service's policies; an array of container firewall policy names.
	Policies pulumi.StringArrayInput
	// Rules priority, must be between 1-100.
	Priority pulumi.IntPtrInput
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression pulumi.StringPtrInput
	// List of scope attributes.
	ScopeVariables ServiceScopeVariableArrayInput
	// Type of the workload. container or host.
	Target pulumi.StringPtrInput
	// The number of containers allocated to the service that are not registered.
	UnregisteredCount pulumi.IntPtrInput
	// Number of high severity vulnerabilities.
	VulnerabilitiesHigh pulumi.IntPtrInput
	// Number of low severity vulnerabilities.
	VulnerabilitiesLow pulumi.IntPtrInput
	// Number of malware.
	VulnerabilitiesMalware pulumi.IntPtrInput
	// Number of medium severity vulnerabilities.
	VulnerabilitiesMedium pulumi.IntPtrInput
	// Number of negligible vulnerabilities.
	VulnerabilitiesNegligible pulumi.IntPtrInput
	// The CVSS average vulnerabilities score.
	VulnerabilitiesScoreAverage pulumi.IntPtrInput
	// Number of sensitive vulnerabilities.
	VulnerabilitiesSensitive pulumi.IntPtrInput
	// Total number of vulnerabilities.
	VulnerabilitiesTotal pulumi.IntPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Indicates the application scope of the service.
	ApplicationScopes []string `pulumi:"applicationScopes"`
	// A textual description of the service record; maximum 500 characters.
	Description *string `pulumi:"description"`
	// Enforcement status of the service.
	Enforce *bool `pulumi:"enforce"`
	// Indicates if monitoring is enabled or not
	Monitoring *bool `pulumi:"monitoring"`
	// The name of the service. It is recommended not to use whitespace characters in the name.
	Name *string `pulumi:"name"`
	// The service's policies; an array of container firewall policy names.
	Policies []string `pulumi:"policies"`
	// Rules priority, must be between 1-100.
	Priority *int `pulumi:"priority"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression *string `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables []ServiceScopeVariable `pulumi:"scopeVariables"`
	// Type of the workload. container or host.
	Target string `pulumi:"target"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Indicates the application scope of the service.
	ApplicationScopes pulumi.StringArrayInput
	// A textual description of the service record; maximum 500 characters.
	Description pulumi.StringPtrInput
	// Enforcement status of the service.
	Enforce pulumi.BoolPtrInput
	// Indicates if monitoring is enabled or not
	Monitoring pulumi.BoolPtrInput
	// The name of the service. It is recommended not to use whitespace characters in the name.
	Name pulumi.StringPtrInput
	// The service's policies; an array of container firewall policy names.
	Policies pulumi.StringArrayInput
	// Rules priority, must be between 1-100.
	Priority pulumi.IntPtrInput
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression pulumi.StringPtrInput
	// List of scope attributes.
	ScopeVariables ServiceScopeVariableArrayInput
	// Type of the workload. container or host.
	Target pulumi.StringInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//	ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//	ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// Indicates the application scope of the service.
func (o ServiceOutput) ApplicationScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Service) pulumi.StringArrayOutput { return v.ApplicationScopes }).(pulumi.StringArrayOutput)
}

// Username of the account that created the service.
func (o ServiceOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Author }).(pulumi.StringOutput)
}

// The number of containers associated with the service.
func (o ServiceOutput) ContainersCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.ContainersCount }).(pulumi.IntOutput)
}

// A textual description of the service record; maximum 500 characters.
func (o ServiceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Enforcement status of the service.
func (o ServiceOutput) Enforce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolPtrOutput { return v.Enforce }).(pulumi.BoolPtrOutput)
}

// Whether the service has been evaluated for security vulnerabilities.
func (o ServiceOutput) Evaluated() pulumi.BoolOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolOutput { return v.Evaluated }).(pulumi.BoolOutput)
}

// Indicates if registered or not.
func (o ServiceOutput) IsRegistered() pulumi.BoolOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolOutput { return v.IsRegistered }).(pulumi.BoolOutput)
}

// Timestamp of the last update in Unix time format.
func (o ServiceOutput) Lastupdate() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.Lastupdate }).(pulumi.IntOutput)
}

// Indicates if monitoring is enabled or not
func (o ServiceOutput) Monitoring() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolPtrOutput { return v.Monitoring }).(pulumi.BoolPtrOutput)
}

// The name of the service. It is recommended not to use whitespace characters in the name.
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of container that are not evaluated.
func (o ServiceOutput) NotEvaluatedCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.NotEvaluatedCount }).(pulumi.IntOutput)
}

// The service's policies; an array of container firewall policy names.
func (o ServiceOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Service) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

// Rules priority, must be between 1-100.
func (o ServiceOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// Logical expression of how to compute the dependency of the scope variables.
func (o ServiceOutput) ScopeExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.ScopeExpression }).(pulumi.StringPtrOutput)
}

// List of scope attributes.
func (o ServiceOutput) ScopeVariables() ServiceScopeVariableArrayOutput {
	return o.ApplyT(func(v *Service) ServiceScopeVariableArrayOutput { return v.ScopeVariables }).(ServiceScopeVariableArrayOutput)
}

// Type of the workload. container or host.
func (o ServiceOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

// The number of containers allocated to the service that are not registered.
func (o ServiceOutput) UnregisteredCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.UnregisteredCount }).(pulumi.IntOutput)
}

// Number of high severity vulnerabilities.
func (o ServiceOutput) VulnerabilitiesHigh() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.VulnerabilitiesHigh }).(pulumi.IntOutput)
}

// Number of low severity vulnerabilities.
func (o ServiceOutput) VulnerabilitiesLow() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.VulnerabilitiesLow }).(pulumi.IntOutput)
}

// Number of malware.
func (o ServiceOutput) VulnerabilitiesMalware() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.VulnerabilitiesMalware }).(pulumi.IntOutput)
}

// Number of medium severity vulnerabilities.
func (o ServiceOutput) VulnerabilitiesMedium() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.VulnerabilitiesMedium }).(pulumi.IntOutput)
}

// Number of negligible vulnerabilities.
func (o ServiceOutput) VulnerabilitiesNegligible() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.VulnerabilitiesNegligible }).(pulumi.IntOutput)
}

// The CVSS average vulnerabilities score.
func (o ServiceOutput) VulnerabilitiesScoreAverage() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.VulnerabilitiesScoreAverage }).(pulumi.IntOutput)
}

// Number of sensitive vulnerabilities.
func (o ServiceOutput) VulnerabilitiesSensitive() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.VulnerabilitiesSensitive }).(pulumi.IntOutput)
}

// Total number of vulnerabilities.
func (o ServiceOutput) VulnerabilitiesTotal() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.VulnerabilitiesTotal }).(pulumi.IntOutput)
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Service {
		return vs[0].([]*Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Service {
		return vs[0].(map[string]*Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceArrayInput)(nil)).Elem(), ServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMapInput)(nil)).Elem(), ServiceMap{})
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
