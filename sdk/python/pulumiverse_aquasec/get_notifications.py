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
    'GetNotificationsResult',
    'AwaitableGetNotificationsResult',
    'get_notifications',
    'get_notifications_output',
]

@pulumi.output_type
class GetNotificationsResult:
    """
    A collection of values returned by getNotifications.
    """
    def __init__(__self__, emails=None, id=None, jiras=None, servicenows=None, slacks=None, splunks=None, teams=None, webhooks=None):
        if emails and not isinstance(emails, list):
            raise TypeError("Expected argument 'emails' to be a list")
        pulumi.set(__self__, "emails", emails)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if jiras and not isinstance(jiras, list):
            raise TypeError("Expected argument 'jiras' to be a list")
        pulumi.set(__self__, "jiras", jiras)
        if servicenows and not isinstance(servicenows, list):
            raise TypeError("Expected argument 'servicenows' to be a list")
        pulumi.set(__self__, "servicenows", servicenows)
        if slacks and not isinstance(slacks, list):
            raise TypeError("Expected argument 'slacks' to be a list")
        pulumi.set(__self__, "slacks", slacks)
        if splunks and not isinstance(splunks, list):
            raise TypeError("Expected argument 'splunks' to be a list")
        pulumi.set(__self__, "splunks", splunks)
        if teams and not isinstance(teams, list):
            raise TypeError("Expected argument 'teams' to be a list")
        pulumi.set(__self__, "teams", teams)
        if webhooks and not isinstance(webhooks, list):
            raise TypeError("Expected argument 'webhooks' to be a list")
        pulumi.set(__self__, "webhooks", webhooks)

    @property
    @pulumi.getter
    def emails(self) -> Sequence['outputs.GetNotificationsEmailResult']:
        return pulumi.get(self, "emails")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def jiras(self) -> Sequence['outputs.GetNotificationsJiraResult']:
        return pulumi.get(self, "jiras")

    @property
    @pulumi.getter
    def servicenows(self) -> Sequence['outputs.GetNotificationsServicenowResult']:
        return pulumi.get(self, "servicenows")

    @property
    @pulumi.getter
    def slacks(self) -> Sequence['outputs.GetNotificationsSlackResult']:
        return pulumi.get(self, "slacks")

    @property
    @pulumi.getter
    def splunks(self) -> Sequence['outputs.GetNotificationsSplunkResult']:
        return pulumi.get(self, "splunks")

    @property
    @pulumi.getter
    def teams(self) -> Sequence['outputs.GetNotificationsTeamResult']:
        return pulumi.get(self, "teams")

    @property
    @pulumi.getter
    def webhooks(self) -> Sequence['outputs.GetNotificationsWebhookResult']:
        return pulumi.get(self, "webhooks")


class AwaitableGetNotificationsResult(GetNotificationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNotificationsResult(
            emails=self.emails,
            id=self.id,
            jiras=self.jiras,
            servicenows=self.servicenows,
            slacks=self.slacks,
            splunks=self.splunks,
            teams=self.teams,
            webhooks=self.webhooks)


def get_notifications(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNotificationsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aquasec:index/getNotifications:getNotifications', __args__, opts=opts, typ=GetNotificationsResult).value

    return AwaitableGetNotificationsResult(
        emails=pulumi.get(__ret__, 'emails'),
        id=pulumi.get(__ret__, 'id'),
        jiras=pulumi.get(__ret__, 'jiras'),
        servicenows=pulumi.get(__ret__, 'servicenows'),
        slacks=pulumi.get(__ret__, 'slacks'),
        splunks=pulumi.get(__ret__, 'splunks'),
        teams=pulumi.get(__ret__, 'teams'),
        webhooks=pulumi.get(__ret__, 'webhooks'))


@_utilities.lift_output_func(get_notifications)
def get_notifications_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNotificationsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
