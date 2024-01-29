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
from ._inputs import *

__all__ = [
    'GetHostRuntimePolicyResult',
    'AwaitableGetHostRuntimePolicyResult',
    'get_host_runtime_policy',
    'get_host_runtime_policy_output',
]

@pulumi.output_type
class GetHostRuntimePolicyResult:
    """
    A collection of values returned by getHostRuntimePolicy.
    """
    def __init__(__self__, application_scopes=None, audit_all_os_user_activity=None, audit_brute_force_login=None, audit_full_command_arguments=None, audit_host_failed_login_events=None, audit_host_successful_login_events=None, audit_user_account_management=None, auditing=None, author=None, block_cryptocurrency_mining=None, blocked_files=None, description=None, enable_ip_reputation=None, enabled=None, enforce=None, enforce_after_days=None, file_integrity_monitorings=None, id=None, malware_scan_options=None, monitor_system_log_integrity=None, monitor_system_time_changes=None, monitor_windows_services=None, name=None, os_groups_alloweds=None, os_groups_blockeds=None, os_users_alloweds=None, os_users_blockeds=None, package_blocks=None, port_scanning_detection=None, scope_expression=None, scope_variables=None, windows_registry_monitorings=None, windows_registry_protections=None):
        if application_scopes and not isinstance(application_scopes, list):
            raise TypeError("Expected argument 'application_scopes' to be a list")
        pulumi.set(__self__, "application_scopes", application_scopes)
        if audit_all_os_user_activity and not isinstance(audit_all_os_user_activity, bool):
            raise TypeError("Expected argument 'audit_all_os_user_activity' to be a bool")
        pulumi.set(__self__, "audit_all_os_user_activity", audit_all_os_user_activity)
        if audit_brute_force_login and not isinstance(audit_brute_force_login, bool):
            raise TypeError("Expected argument 'audit_brute_force_login' to be a bool")
        pulumi.set(__self__, "audit_brute_force_login", audit_brute_force_login)
        if audit_full_command_arguments and not isinstance(audit_full_command_arguments, bool):
            raise TypeError("Expected argument 'audit_full_command_arguments' to be a bool")
        pulumi.set(__self__, "audit_full_command_arguments", audit_full_command_arguments)
        if audit_host_failed_login_events and not isinstance(audit_host_failed_login_events, bool):
            raise TypeError("Expected argument 'audit_host_failed_login_events' to be a bool")
        pulumi.set(__self__, "audit_host_failed_login_events", audit_host_failed_login_events)
        if audit_host_successful_login_events and not isinstance(audit_host_successful_login_events, bool):
            raise TypeError("Expected argument 'audit_host_successful_login_events' to be a bool")
        pulumi.set(__self__, "audit_host_successful_login_events", audit_host_successful_login_events)
        if audit_user_account_management and not isinstance(audit_user_account_management, bool):
            raise TypeError("Expected argument 'audit_user_account_management' to be a bool")
        pulumi.set(__self__, "audit_user_account_management", audit_user_account_management)
        if auditing and not isinstance(auditing, dict):
            raise TypeError("Expected argument 'auditing' to be a dict")
        pulumi.set(__self__, "auditing", auditing)
        if author and not isinstance(author, str):
            raise TypeError("Expected argument 'author' to be a str")
        pulumi.set(__self__, "author", author)
        if block_cryptocurrency_mining and not isinstance(block_cryptocurrency_mining, bool):
            raise TypeError("Expected argument 'block_cryptocurrency_mining' to be a bool")
        pulumi.set(__self__, "block_cryptocurrency_mining", block_cryptocurrency_mining)
        if blocked_files and not isinstance(blocked_files, list):
            raise TypeError("Expected argument 'blocked_files' to be a list")
        pulumi.set(__self__, "blocked_files", blocked_files)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enable_ip_reputation and not isinstance(enable_ip_reputation, bool):
            raise TypeError("Expected argument 'enable_ip_reputation' to be a bool")
        pulumi.set(__self__, "enable_ip_reputation", enable_ip_reputation)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if enforce and not isinstance(enforce, bool):
            raise TypeError("Expected argument 'enforce' to be a bool")
        pulumi.set(__self__, "enforce", enforce)
        if enforce_after_days and not isinstance(enforce_after_days, int):
            raise TypeError("Expected argument 'enforce_after_days' to be a int")
        pulumi.set(__self__, "enforce_after_days", enforce_after_days)
        if file_integrity_monitorings and not isinstance(file_integrity_monitorings, list):
            raise TypeError("Expected argument 'file_integrity_monitorings' to be a list")
        pulumi.set(__self__, "file_integrity_monitorings", file_integrity_monitorings)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if malware_scan_options and not isinstance(malware_scan_options, list):
            raise TypeError("Expected argument 'malware_scan_options' to be a list")
        pulumi.set(__self__, "malware_scan_options", malware_scan_options)
        if monitor_system_log_integrity and not isinstance(monitor_system_log_integrity, bool):
            raise TypeError("Expected argument 'monitor_system_log_integrity' to be a bool")
        pulumi.set(__self__, "monitor_system_log_integrity", monitor_system_log_integrity)
        if monitor_system_time_changes and not isinstance(monitor_system_time_changes, bool):
            raise TypeError("Expected argument 'monitor_system_time_changes' to be a bool")
        pulumi.set(__self__, "monitor_system_time_changes", monitor_system_time_changes)
        if monitor_windows_services and not isinstance(monitor_windows_services, bool):
            raise TypeError("Expected argument 'monitor_windows_services' to be a bool")
        pulumi.set(__self__, "monitor_windows_services", monitor_windows_services)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if os_groups_alloweds and not isinstance(os_groups_alloweds, list):
            raise TypeError("Expected argument 'os_groups_alloweds' to be a list")
        pulumi.set(__self__, "os_groups_alloweds", os_groups_alloweds)
        if os_groups_blockeds and not isinstance(os_groups_blockeds, list):
            raise TypeError("Expected argument 'os_groups_blockeds' to be a list")
        pulumi.set(__self__, "os_groups_blockeds", os_groups_blockeds)
        if os_users_alloweds and not isinstance(os_users_alloweds, list):
            raise TypeError("Expected argument 'os_users_alloweds' to be a list")
        pulumi.set(__self__, "os_users_alloweds", os_users_alloweds)
        if os_users_blockeds and not isinstance(os_users_blockeds, list):
            raise TypeError("Expected argument 'os_users_blockeds' to be a list")
        pulumi.set(__self__, "os_users_blockeds", os_users_blockeds)
        if package_blocks and not isinstance(package_blocks, list):
            raise TypeError("Expected argument 'package_blocks' to be a list")
        pulumi.set(__self__, "package_blocks", package_blocks)
        if port_scanning_detection and not isinstance(port_scanning_detection, bool):
            raise TypeError("Expected argument 'port_scanning_detection' to be a bool")
        pulumi.set(__self__, "port_scanning_detection", port_scanning_detection)
        if scope_expression and not isinstance(scope_expression, str):
            raise TypeError("Expected argument 'scope_expression' to be a str")
        pulumi.set(__self__, "scope_expression", scope_expression)
        if scope_variables and not isinstance(scope_variables, list):
            raise TypeError("Expected argument 'scope_variables' to be a list")
        pulumi.set(__self__, "scope_variables", scope_variables)
        if windows_registry_monitorings and not isinstance(windows_registry_monitorings, list):
            raise TypeError("Expected argument 'windows_registry_monitorings' to be a list")
        pulumi.set(__self__, "windows_registry_monitorings", windows_registry_monitorings)
        if windows_registry_protections and not isinstance(windows_registry_protections, list):
            raise TypeError("Expected argument 'windows_registry_protections' to be a list")
        pulumi.set(__self__, "windows_registry_protections", windows_registry_protections)

    @property
    @pulumi.getter(name="applicationScopes")
    def application_scopes(self) -> Sequence[str]:
        """
        Indicates the application scope of the service.
        """
        return pulumi.get(self, "application_scopes")

    @property
    @pulumi.getter(name="auditAllOsUserActivity")
    def audit_all_os_user_activity(self) -> bool:
        """
        If true, all process activity will be audited.
        """
        return pulumi.get(self, "audit_all_os_user_activity")

    @property
    @pulumi.getter(name="auditBruteForceLogin")
    def audit_brute_force_login(self) -> bool:
        """
        Detects brute force login attempts
        """
        return pulumi.get(self, "audit_brute_force_login")

    @property
    @pulumi.getter(name="auditFullCommandArguments")
    def audit_full_command_arguments(self) -> bool:
        """
        If true, full command arguments will be audited.
        """
        return pulumi.get(self, "audit_full_command_arguments")

    @property
    @pulumi.getter(name="auditHostFailedLoginEvents")
    def audit_host_failed_login_events(self) -> bool:
        """
        If true, host failed logins will be audited.
        """
        return pulumi.get(self, "audit_host_failed_login_events")

    @property
    @pulumi.getter(name="auditHostSuccessfulLoginEvents")
    def audit_host_successful_login_events(self) -> bool:
        """
        If true, host successful logins will be audited.
        """
        return pulumi.get(self, "audit_host_successful_login_events")

    @property
    @pulumi.getter(name="auditUserAccountManagement")
    def audit_user_account_management(self) -> bool:
        """
        If true, account management will be audited.
        """
        return pulumi.get(self, "audit_user_account_management")

    @property
    @pulumi.getter
    def auditing(self) -> Optional['outputs.GetHostRuntimePolicyAuditingResult']:
        return pulumi.get(self, "auditing")

    @property
    @pulumi.getter
    def author(self) -> str:
        """
        Username of the account that created the service.
        """
        return pulumi.get(self, "author")

    @property
    @pulumi.getter(name="blockCryptocurrencyMining")
    def block_cryptocurrency_mining(self) -> bool:
        """
        Detect and prevent communication to DNS/IP addresses known to be used for Cryptocurrency Mining
        """
        return pulumi.get(self, "block_cryptocurrency_mining")

    @property
    @pulumi.getter(name="blockedFiles")
    def blocked_files(self) -> Sequence[str]:
        """
        List of files that are prevented from being read, modified and executed in the containers.
        """
        return pulumi.get(self, "blocked_files")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the host runtime policy
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableIpReputation")
    def enable_ip_reputation(self) -> bool:
        """
        If true, detect and prevent communication from containers to IP addresses known to have a bad reputation.
        """
        return pulumi.get(self, "enable_ip_reputation")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Indicates if the runtime policy is enabled or not.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def enforce(self) -> bool:
        """
        Indicates that policy should effect container execution (not just for audit).
        """
        return pulumi.get(self, "enforce")

    @property
    @pulumi.getter(name="enforceAfterDays")
    def enforce_after_days(self) -> int:
        """
        Indicates the number of days after which the runtime policy will be changed to enforce mode.
        """
        return pulumi.get(self, "enforce_after_days")

    @property
    @pulumi.getter(name="fileIntegrityMonitorings")
    def file_integrity_monitorings(self) -> Optional[Sequence['outputs.GetHostRuntimePolicyFileIntegrityMonitoringResult']]:
        """
        Configuration for file integrity monitoring.
        """
        return pulumi.get(self, "file_integrity_monitorings")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="malwareScanOptions")
    def malware_scan_options(self) -> Optional[Sequence['outputs.GetHostRuntimePolicyMalwareScanOptionResult']]:
        """
        Configuration for Real-Time Malware Protection.
        """
        return pulumi.get(self, "malware_scan_options")

    @property
    @pulumi.getter(name="monitorSystemLogIntegrity")
    def monitor_system_log_integrity(self) -> bool:
        """
        If true, system log will be monitored.
        """
        return pulumi.get(self, "monitor_system_log_integrity")

    @property
    @pulumi.getter(name="monitorSystemTimeChanges")
    def monitor_system_time_changes(self) -> bool:
        """
        If true, system time changes will be monitored.
        """
        return pulumi.get(self, "monitor_system_time_changes")

    @property
    @pulumi.getter(name="monitorWindowsServices")
    def monitor_windows_services(self) -> bool:
        """
        If true, windows service operations will be monitored.
        """
        return pulumi.get(self, "monitor_windows_services")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the host runtime policy
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osGroupsAlloweds")
    def os_groups_alloweds(self) -> Sequence[str]:
        """
        List of OS (Linux or Windows) groups that are allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
        """
        return pulumi.get(self, "os_groups_alloweds")

    @property
    @pulumi.getter(name="osGroupsBlockeds")
    def os_groups_blockeds(self) -> Sequence[str]:
        """
        List of OS (Linux or Windows) groups that are not allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
        """
        return pulumi.get(self, "os_groups_blockeds")

    @property
    @pulumi.getter(name="osUsersAlloweds")
    def os_users_alloweds(self) -> Sequence[str]:
        """
        List of OS (Linux or Windows) users that are allowed to authenticate to the host, and block authentication requests from all others.
        """
        return pulumi.get(self, "os_users_alloweds")

    @property
    @pulumi.getter(name="osUsersBlockeds")
    def os_users_blockeds(self) -> Sequence[str]:
        """
        List of OS (Linux or Windows) users that are not allowed to authenticate to the host, and block authentication requests from all others.
        """
        return pulumi.get(self, "os_users_blockeds")

    @property
    @pulumi.getter(name="packageBlocks")
    def package_blocks(self) -> Optional[Sequence['outputs.GetHostRuntimePolicyPackageBlockResult']]:
        return pulumi.get(self, "package_blocks")

    @property
    @pulumi.getter(name="portScanningDetection")
    def port_scanning_detection(self) -> bool:
        """
        If true, port scanning behaviors will be audited.
        """
        return pulumi.get(self, "port_scanning_detection")

    @property
    @pulumi.getter(name="scopeExpression")
    def scope_expression(self) -> str:
        """
        Logical expression of how to compute the dependency of the scope variables.
        """
        return pulumi.get(self, "scope_expression")

    @property
    @pulumi.getter(name="scopeVariables")
    def scope_variables(self) -> Sequence['outputs.GetHostRuntimePolicyScopeVariableResult']:
        """
        List of scope attributes.
        """
        return pulumi.get(self, "scope_variables")

    @property
    @pulumi.getter(name="windowsRegistryMonitorings")
    def windows_registry_monitorings(self) -> Sequence['outputs.GetHostRuntimePolicyWindowsRegistryMonitoringResult']:
        """
        Configuration for windows registry monitoring.
        """
        return pulumi.get(self, "windows_registry_monitorings")

    @property
    @pulumi.getter(name="windowsRegistryProtections")
    def windows_registry_protections(self) -> Sequence['outputs.GetHostRuntimePolicyWindowsRegistryProtectionResult']:
        """
        Configuration for windows registry protection.
        """
        return pulumi.get(self, "windows_registry_protections")


class AwaitableGetHostRuntimePolicyResult(GetHostRuntimePolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHostRuntimePolicyResult(
            application_scopes=self.application_scopes,
            audit_all_os_user_activity=self.audit_all_os_user_activity,
            audit_brute_force_login=self.audit_brute_force_login,
            audit_full_command_arguments=self.audit_full_command_arguments,
            audit_host_failed_login_events=self.audit_host_failed_login_events,
            audit_host_successful_login_events=self.audit_host_successful_login_events,
            audit_user_account_management=self.audit_user_account_management,
            auditing=self.auditing,
            author=self.author,
            block_cryptocurrency_mining=self.block_cryptocurrency_mining,
            blocked_files=self.blocked_files,
            description=self.description,
            enable_ip_reputation=self.enable_ip_reputation,
            enabled=self.enabled,
            enforce=self.enforce,
            enforce_after_days=self.enforce_after_days,
            file_integrity_monitorings=self.file_integrity_monitorings,
            id=self.id,
            malware_scan_options=self.malware_scan_options,
            monitor_system_log_integrity=self.monitor_system_log_integrity,
            monitor_system_time_changes=self.monitor_system_time_changes,
            monitor_windows_services=self.monitor_windows_services,
            name=self.name,
            os_groups_alloweds=self.os_groups_alloweds,
            os_groups_blockeds=self.os_groups_blockeds,
            os_users_alloweds=self.os_users_alloweds,
            os_users_blockeds=self.os_users_blockeds,
            package_blocks=self.package_blocks,
            port_scanning_detection=self.port_scanning_detection,
            scope_expression=self.scope_expression,
            scope_variables=self.scope_variables,
            windows_registry_monitorings=self.windows_registry_monitorings,
            windows_registry_protections=self.windows_registry_protections)


def get_host_runtime_policy(auditing: Optional[pulumi.InputType['GetHostRuntimePolicyAuditingArgs']] = None,
                            file_integrity_monitorings: Optional[Sequence[pulumi.InputType['GetHostRuntimePolicyFileIntegrityMonitoringArgs']]] = None,
                            malware_scan_options: Optional[Sequence[pulumi.InputType['GetHostRuntimePolicyMalwareScanOptionArgs']]] = None,
                            name: Optional[str] = None,
                            package_blocks: Optional[Sequence[pulumi.InputType['GetHostRuntimePolicyPackageBlockArgs']]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHostRuntimePolicyResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_aquasec as aquasec

    host_runtime_policy = aquasec.get_host_runtime_policy(name="hostRuntimePolicyName")
    pulumi.export("hostRuntimePolicyDetails", host_runtime_policy)
    ```


    :param Sequence[pulumi.InputType['GetHostRuntimePolicyFileIntegrityMonitoringArgs']] file_integrity_monitorings: Configuration for file integrity monitoring.
    :param Sequence[pulumi.InputType['GetHostRuntimePolicyMalwareScanOptionArgs']] malware_scan_options: Configuration for Real-Time Malware Protection.
    """
    __args__ = dict()
    __args__['auditing'] = auditing
    __args__['fileIntegrityMonitorings'] = file_integrity_monitorings
    __args__['malwareScanOptions'] = malware_scan_options
    __args__['name'] = name
    __args__['packageBlocks'] = package_blocks
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aquasec:index/getHostRuntimePolicy:getHostRuntimePolicy', __args__, opts=opts, typ=GetHostRuntimePolicyResult).value

    return AwaitableGetHostRuntimePolicyResult(
        application_scopes=pulumi.get(__ret__, 'application_scopes'),
        audit_all_os_user_activity=pulumi.get(__ret__, 'audit_all_os_user_activity'),
        audit_brute_force_login=pulumi.get(__ret__, 'audit_brute_force_login'),
        audit_full_command_arguments=pulumi.get(__ret__, 'audit_full_command_arguments'),
        audit_host_failed_login_events=pulumi.get(__ret__, 'audit_host_failed_login_events'),
        audit_host_successful_login_events=pulumi.get(__ret__, 'audit_host_successful_login_events'),
        audit_user_account_management=pulumi.get(__ret__, 'audit_user_account_management'),
        auditing=pulumi.get(__ret__, 'auditing'),
        author=pulumi.get(__ret__, 'author'),
        block_cryptocurrency_mining=pulumi.get(__ret__, 'block_cryptocurrency_mining'),
        blocked_files=pulumi.get(__ret__, 'blocked_files'),
        description=pulumi.get(__ret__, 'description'),
        enable_ip_reputation=pulumi.get(__ret__, 'enable_ip_reputation'),
        enabled=pulumi.get(__ret__, 'enabled'),
        enforce=pulumi.get(__ret__, 'enforce'),
        enforce_after_days=pulumi.get(__ret__, 'enforce_after_days'),
        file_integrity_monitorings=pulumi.get(__ret__, 'file_integrity_monitorings'),
        id=pulumi.get(__ret__, 'id'),
        malware_scan_options=pulumi.get(__ret__, 'malware_scan_options'),
        monitor_system_log_integrity=pulumi.get(__ret__, 'monitor_system_log_integrity'),
        monitor_system_time_changes=pulumi.get(__ret__, 'monitor_system_time_changes'),
        monitor_windows_services=pulumi.get(__ret__, 'monitor_windows_services'),
        name=pulumi.get(__ret__, 'name'),
        os_groups_alloweds=pulumi.get(__ret__, 'os_groups_alloweds'),
        os_groups_blockeds=pulumi.get(__ret__, 'os_groups_blockeds'),
        os_users_alloweds=pulumi.get(__ret__, 'os_users_alloweds'),
        os_users_blockeds=pulumi.get(__ret__, 'os_users_blockeds'),
        package_blocks=pulumi.get(__ret__, 'package_blocks'),
        port_scanning_detection=pulumi.get(__ret__, 'port_scanning_detection'),
        scope_expression=pulumi.get(__ret__, 'scope_expression'),
        scope_variables=pulumi.get(__ret__, 'scope_variables'),
        windows_registry_monitorings=pulumi.get(__ret__, 'windows_registry_monitorings'),
        windows_registry_protections=pulumi.get(__ret__, 'windows_registry_protections'))


@_utilities.lift_output_func(get_host_runtime_policy)
def get_host_runtime_policy_output(auditing: Optional[pulumi.Input[Optional[pulumi.InputType['GetHostRuntimePolicyAuditingArgs']]]] = None,
                                   file_integrity_monitorings: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetHostRuntimePolicyFileIntegrityMonitoringArgs']]]]] = None,
                                   malware_scan_options: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetHostRuntimePolicyMalwareScanOptionArgs']]]]] = None,
                                   name: Optional[pulumi.Input[str]] = None,
                                   package_blocks: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetHostRuntimePolicyPackageBlockArgs']]]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHostRuntimePolicyResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_aquasec as aquasec

    host_runtime_policy = aquasec.get_host_runtime_policy(name="hostRuntimePolicyName")
    pulumi.export("hostRuntimePolicyDetails", host_runtime_policy)
    ```


    :param Sequence[pulumi.InputType['GetHostRuntimePolicyFileIntegrityMonitoringArgs']] file_integrity_monitorings: Configuration for file integrity monitoring.
    :param Sequence[pulumi.InputType['GetHostRuntimePolicyMalwareScanOptionArgs']] malware_scan_options: Configuration for Real-Time Malware Protection.
    """
    ...
