# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetRolesMappingResult',
    'AwaitableGetRolesMappingResult',
    'get_roles_mapping',
]

@pulumi.output_type
class GetRolesMappingResult:
    """
    A collection of values returned by getRolesMapping.
    """
    def __init__(__self__, id=None, oauth2s=None, openids=None, samls=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if oauth2s and not isinstance(oauth2s, list):
            raise TypeError("Expected argument 'oauth2s' to be a list")
        pulumi.set(__self__, "oauth2s", oauth2s)
        if openids and not isinstance(openids, list):
            raise TypeError("Expected argument 'openids' to be a list")
        pulumi.set(__self__, "openids", openids)
        if samls and not isinstance(samls, list):
            raise TypeError("Expected argument 'samls' to be a list")
        pulumi.set(__self__, "samls", samls)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def oauth2s(self) -> Sequence['outputs.GetRolesMappingOauth2Result']:
        """
        Oauth2 Authentication
        """
        return pulumi.get(self, "oauth2s")

    @property
    @pulumi.getter
    def openids(self) -> Sequence['outputs.GetRolesMappingOpenidResult']:
        """
        OpenId Authentication
        """
        return pulumi.get(self, "openids")

    @property
    @pulumi.getter
    def samls(self) -> Sequence['outputs.GetRolesMappingSamlResult']:
        """
        SAML Authentication
        """
        return pulumi.get(self, "samls")


class AwaitableGetRolesMappingResult(GetRolesMappingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRolesMappingResult(
            id=self.id,
            oauth2s=self.oauth2s,
            openids=self.openids,
            samls=self.samls)


def get_roles_mapping(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRolesMappingResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aquasec:index/getRolesMapping:getRolesMapping', __args__, opts=opts, typ=GetRolesMappingResult).value

    return AwaitableGetRolesMappingResult(
        id=__ret__.id,
        oauth2s=__ret__.oauth2s,
        openids=__ret__.openids,
        samls=__ret__.samls)
