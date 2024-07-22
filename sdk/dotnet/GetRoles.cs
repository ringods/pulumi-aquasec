// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Aquasec
{
    public static class GetRoles
    {
        /// <summary>
        /// The data source `aquasec.getRoles` provides a method to query all roles within the Aqua account managementrole database. The fields returned from this query are detailed in the Schema section below.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aquasec = Pulumi.Aquasec;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var roles = Aquasec.GetRoles.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstUserName"] = roles.Apply(getRolesResult =&gt; getRolesResult.Roles[0]),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetRolesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRolesResult>("aquasec:index/getRoles:getRoles", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// The data source `aquasec.getRoles` provides a method to query all roles within the Aqua account managementrole database. The fields returned from this query are detailed in the Schema section below.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aquasec = Pulumi.Aquasec;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var roles = Aquasec.GetRoles.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstUserName"] = roles.Apply(getRolesResult =&gt; getRolesResult.Roles[0]),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetRolesResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRolesResult>("aquasec:index/getRoles:getRoles", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetRolesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetRolesRoleResult> Roles;

        [OutputConstructor]
        private GetRolesResult(
            string id,

            ImmutableArray<Outputs.GetRolesRoleResult> roles)
        {
            Id = id;
            Roles = roles;
        }
    }
}
