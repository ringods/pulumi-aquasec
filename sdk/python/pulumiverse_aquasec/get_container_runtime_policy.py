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
    'GetContainerRuntimePolicyResult',
    'AwaitableGetContainerRuntimePolicyResult',
    'get_container_runtime_policy',
    'get_container_runtime_policy_output',
]

@pulumi.output_type
class GetContainerRuntimePolicyResult:
    """
    A collection of values returned by getContainerRuntimePolicy.
    """
    def __init__(__self__, allowed_executables=None, allowed_registries=None, application_scopes=None, audit_all_network_activity=None, audit_all_processes_activity=None, audit_full_command_arguments=None, author=None, block_access_host_network=None, block_adding_capabilities=None, block_container_exec=None, block_cryptocurrency_mining=None, block_fileless_exec=None, block_low_port_binding=None, block_non_compliant_images=None, block_non_compliant_workloads=None, block_non_k8s_containers=None, block_privileged_containers=None, block_reverse_shell=None, block_root_user=None, block_unregistered_images=None, block_use_ipc_namespace=None, block_use_pid_namespace=None, block_use_user_namespace=None, block_use_uts_namespace=None, blocked_capabilities=None, blocked_executables=None, blocked_files=None, blocked_inbound_ports=None, blocked_outbound_ports=None, blocked_packages=None, blocked_volumes=None, container_exec_allowed_processes=None, description=None, enable_drift_prevention=None, enable_fork_guard=None, enable_ip_reputation_security=None, enable_port_scan_detection=None, enabled=None, enforce=None, enforce_after_days=None, exceptional_readonly_files_and_directories=None, file_integrity_monitorings=None, fork_guard_process_limit=None, id=None, limit_new_privileges=None, malware_scan_options=None, monitor_system_time_changes=None, name=None, readonly_files_and_directories=None, reverse_shell_allowed_ips=None, reverse_shell_allowed_processes=None, scope_expression=None, scope_variables=None):
        if allowed_executables and not isinstance(allowed_executables, list):
            raise TypeError("Expected argument 'allowed_executables' to be a list")
        pulumi.set(__self__, "allowed_executables", allowed_executables)
        if allowed_registries and not isinstance(allowed_registries, list):
            raise TypeError("Expected argument 'allowed_registries' to be a list")
        pulumi.set(__self__, "allowed_registries", allowed_registries)
        if application_scopes and not isinstance(application_scopes, list):
            raise TypeError("Expected argument 'application_scopes' to be a list")
        pulumi.set(__self__, "application_scopes", application_scopes)
        if audit_all_network_activity and not isinstance(audit_all_network_activity, bool):
            raise TypeError("Expected argument 'audit_all_network_activity' to be a bool")
        pulumi.set(__self__, "audit_all_network_activity", audit_all_network_activity)
        if audit_all_processes_activity and not isinstance(audit_all_processes_activity, bool):
            raise TypeError("Expected argument 'audit_all_processes_activity' to be a bool")
        pulumi.set(__self__, "audit_all_processes_activity", audit_all_processes_activity)
        if audit_full_command_arguments and not isinstance(audit_full_command_arguments, bool):
            raise TypeError("Expected argument 'audit_full_command_arguments' to be a bool")
        pulumi.set(__self__, "audit_full_command_arguments", audit_full_command_arguments)
        if author and not isinstance(author, str):
            raise TypeError("Expected argument 'author' to be a str")
        pulumi.set(__self__, "author", author)
        if block_access_host_network and not isinstance(block_access_host_network, bool):
            raise TypeError("Expected argument 'block_access_host_network' to be a bool")
        pulumi.set(__self__, "block_access_host_network", block_access_host_network)
        if block_adding_capabilities and not isinstance(block_adding_capabilities, bool):
            raise TypeError("Expected argument 'block_adding_capabilities' to be a bool")
        pulumi.set(__self__, "block_adding_capabilities", block_adding_capabilities)
        if block_container_exec and not isinstance(block_container_exec, bool):
            raise TypeError("Expected argument 'block_container_exec' to be a bool")
        pulumi.set(__self__, "block_container_exec", block_container_exec)
        if block_cryptocurrency_mining and not isinstance(block_cryptocurrency_mining, bool):
            raise TypeError("Expected argument 'block_cryptocurrency_mining' to be a bool")
        pulumi.set(__self__, "block_cryptocurrency_mining", block_cryptocurrency_mining)
        if block_fileless_exec and not isinstance(block_fileless_exec, bool):
            raise TypeError("Expected argument 'block_fileless_exec' to be a bool")
        pulumi.set(__self__, "block_fileless_exec", block_fileless_exec)
        if block_low_port_binding and not isinstance(block_low_port_binding, bool):
            raise TypeError("Expected argument 'block_low_port_binding' to be a bool")
        pulumi.set(__self__, "block_low_port_binding", block_low_port_binding)
        if block_non_compliant_images and not isinstance(block_non_compliant_images, bool):
            raise TypeError("Expected argument 'block_non_compliant_images' to be a bool")
        pulumi.set(__self__, "block_non_compliant_images", block_non_compliant_images)
        if block_non_compliant_workloads and not isinstance(block_non_compliant_workloads, bool):
            raise TypeError("Expected argument 'block_non_compliant_workloads' to be a bool")
        pulumi.set(__self__, "block_non_compliant_workloads", block_non_compliant_workloads)
        if block_non_k8s_containers and not isinstance(block_non_k8s_containers, bool):
            raise TypeError("Expected argument 'block_non_k8s_containers' to be a bool")
        pulumi.set(__self__, "block_non_k8s_containers", block_non_k8s_containers)
        if block_privileged_containers and not isinstance(block_privileged_containers, bool):
            raise TypeError("Expected argument 'block_privileged_containers' to be a bool")
        pulumi.set(__self__, "block_privileged_containers", block_privileged_containers)
        if block_reverse_shell and not isinstance(block_reverse_shell, bool):
            raise TypeError("Expected argument 'block_reverse_shell' to be a bool")
        pulumi.set(__self__, "block_reverse_shell", block_reverse_shell)
        if block_root_user and not isinstance(block_root_user, bool):
            raise TypeError("Expected argument 'block_root_user' to be a bool")
        pulumi.set(__self__, "block_root_user", block_root_user)
        if block_unregistered_images and not isinstance(block_unregistered_images, bool):
            raise TypeError("Expected argument 'block_unregistered_images' to be a bool")
        pulumi.set(__self__, "block_unregistered_images", block_unregistered_images)
        if block_use_ipc_namespace and not isinstance(block_use_ipc_namespace, bool):
            raise TypeError("Expected argument 'block_use_ipc_namespace' to be a bool")
        pulumi.set(__self__, "block_use_ipc_namespace", block_use_ipc_namespace)
        if block_use_pid_namespace and not isinstance(block_use_pid_namespace, bool):
            raise TypeError("Expected argument 'block_use_pid_namespace' to be a bool")
        pulumi.set(__self__, "block_use_pid_namespace", block_use_pid_namespace)
        if block_use_user_namespace and not isinstance(block_use_user_namespace, bool):
            raise TypeError("Expected argument 'block_use_user_namespace' to be a bool")
        pulumi.set(__self__, "block_use_user_namespace", block_use_user_namespace)
        if block_use_uts_namespace and not isinstance(block_use_uts_namespace, bool):
            raise TypeError("Expected argument 'block_use_uts_namespace' to be a bool")
        pulumi.set(__self__, "block_use_uts_namespace", block_use_uts_namespace)
        if blocked_capabilities and not isinstance(blocked_capabilities, list):
            raise TypeError("Expected argument 'blocked_capabilities' to be a list")
        pulumi.set(__self__, "blocked_capabilities", blocked_capabilities)
        if blocked_executables and not isinstance(blocked_executables, list):
            raise TypeError("Expected argument 'blocked_executables' to be a list")
        pulumi.set(__self__, "blocked_executables", blocked_executables)
        if blocked_files and not isinstance(blocked_files, list):
            raise TypeError("Expected argument 'blocked_files' to be a list")
        pulumi.set(__self__, "blocked_files", blocked_files)
        if blocked_inbound_ports and not isinstance(blocked_inbound_ports, list):
            raise TypeError("Expected argument 'blocked_inbound_ports' to be a list")
        pulumi.set(__self__, "blocked_inbound_ports", blocked_inbound_ports)
        if blocked_outbound_ports and not isinstance(blocked_outbound_ports, list):
            raise TypeError("Expected argument 'blocked_outbound_ports' to be a list")
        pulumi.set(__self__, "blocked_outbound_ports", blocked_outbound_ports)
        if blocked_packages and not isinstance(blocked_packages, list):
            raise TypeError("Expected argument 'blocked_packages' to be a list")
        pulumi.set(__self__, "blocked_packages", blocked_packages)
        if blocked_volumes and not isinstance(blocked_volumes, list):
            raise TypeError("Expected argument 'blocked_volumes' to be a list")
        pulumi.set(__self__, "blocked_volumes", blocked_volumes)
        if container_exec_allowed_processes and not isinstance(container_exec_allowed_processes, list):
            raise TypeError("Expected argument 'container_exec_allowed_processes' to be a list")
        pulumi.set(__self__, "container_exec_allowed_processes", container_exec_allowed_processes)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enable_drift_prevention and not isinstance(enable_drift_prevention, bool):
            raise TypeError("Expected argument 'enable_drift_prevention' to be a bool")
        pulumi.set(__self__, "enable_drift_prevention", enable_drift_prevention)
        if enable_fork_guard and not isinstance(enable_fork_guard, bool):
            raise TypeError("Expected argument 'enable_fork_guard' to be a bool")
        pulumi.set(__self__, "enable_fork_guard", enable_fork_guard)
        if enable_ip_reputation_security and not isinstance(enable_ip_reputation_security, bool):
            raise TypeError("Expected argument 'enable_ip_reputation_security' to be a bool")
        pulumi.set(__self__, "enable_ip_reputation_security", enable_ip_reputation_security)
        if enable_port_scan_detection and not isinstance(enable_port_scan_detection, bool):
            raise TypeError("Expected argument 'enable_port_scan_detection' to be a bool")
        pulumi.set(__self__, "enable_port_scan_detection", enable_port_scan_detection)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if enforce and not isinstance(enforce, bool):
            raise TypeError("Expected argument 'enforce' to be a bool")
        pulumi.set(__self__, "enforce", enforce)
        if enforce_after_days and not isinstance(enforce_after_days, int):
            raise TypeError("Expected argument 'enforce_after_days' to be a int")
        pulumi.set(__self__, "enforce_after_days", enforce_after_days)
        if exceptional_readonly_files_and_directories and not isinstance(exceptional_readonly_files_and_directories, list):
            raise TypeError("Expected argument 'exceptional_readonly_files_and_directories' to be a list")
        pulumi.set(__self__, "exceptional_readonly_files_and_directories", exceptional_readonly_files_and_directories)
        if file_integrity_monitorings and not isinstance(file_integrity_monitorings, list):
            raise TypeError("Expected argument 'file_integrity_monitorings' to be a list")
        pulumi.set(__self__, "file_integrity_monitorings", file_integrity_monitorings)
        if fork_guard_process_limit and not isinstance(fork_guard_process_limit, int):
            raise TypeError("Expected argument 'fork_guard_process_limit' to be a int")
        pulumi.set(__self__, "fork_guard_process_limit", fork_guard_process_limit)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if limit_new_privileges and not isinstance(limit_new_privileges, bool):
            raise TypeError("Expected argument 'limit_new_privileges' to be a bool")
        pulumi.set(__self__, "limit_new_privileges", limit_new_privileges)
        if malware_scan_options and not isinstance(malware_scan_options, list):
            raise TypeError("Expected argument 'malware_scan_options' to be a list")
        pulumi.set(__self__, "malware_scan_options", malware_scan_options)
        if monitor_system_time_changes and not isinstance(monitor_system_time_changes, bool):
            raise TypeError("Expected argument 'monitor_system_time_changes' to be a bool")
        pulumi.set(__self__, "monitor_system_time_changes", monitor_system_time_changes)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if readonly_files_and_directories and not isinstance(readonly_files_and_directories, list):
            raise TypeError("Expected argument 'readonly_files_and_directories' to be a list")
        pulumi.set(__self__, "readonly_files_and_directories", readonly_files_and_directories)
        if reverse_shell_allowed_ips and not isinstance(reverse_shell_allowed_ips, list):
            raise TypeError("Expected argument 'reverse_shell_allowed_ips' to be a list")
        pulumi.set(__self__, "reverse_shell_allowed_ips", reverse_shell_allowed_ips)
        if reverse_shell_allowed_processes and not isinstance(reverse_shell_allowed_processes, list):
            raise TypeError("Expected argument 'reverse_shell_allowed_processes' to be a list")
        pulumi.set(__self__, "reverse_shell_allowed_processes", reverse_shell_allowed_processes)
        if scope_expression and not isinstance(scope_expression, str):
            raise TypeError("Expected argument 'scope_expression' to be a str")
        pulumi.set(__self__, "scope_expression", scope_expression)
        if scope_variables and not isinstance(scope_variables, list):
            raise TypeError("Expected argument 'scope_variables' to be a list")
        pulumi.set(__self__, "scope_variables", scope_variables)

    @property
    @pulumi.getter(name="allowedExecutables")
    def allowed_executables(self) -> Sequence[str]:
        """
        List of executables that are allowed for the user.
        """
        return pulumi.get(self, "allowed_executables")

    @property
    @pulumi.getter(name="allowedRegistries")
    def allowed_registries(self) -> Sequence[str]:
        """
        List of registries that allowed for running containers.
        """
        return pulumi.get(self, "allowed_registries")

    @property
    @pulumi.getter(name="applicationScopes")
    def application_scopes(self) -> Sequence[str]:
        """
        Indicates the application scope of the service.
        """
        return pulumi.get(self, "application_scopes")

    @property
    @pulumi.getter(name="auditAllNetworkActivity")
    def audit_all_network_activity(self) -> bool:
        """
        If true, all network activity will be audited.
        """
        return pulumi.get(self, "audit_all_network_activity")

    @property
    @pulumi.getter(name="auditAllProcessesActivity")
    def audit_all_processes_activity(self) -> bool:
        """
        If true, all process activity will be audited.
        """
        return pulumi.get(self, "audit_all_processes_activity")

    @property
    @pulumi.getter(name="auditFullCommandArguments")
    def audit_full_command_arguments(self) -> bool:
        """
        If true, full command arguments will be audited.
        """
        return pulumi.get(self, "audit_full_command_arguments")

    @property
    @pulumi.getter
    def author(self) -> str:
        """
        Username of the account that created the service.
        """
        return pulumi.get(self, "author")

    @property
    @pulumi.getter(name="blockAccessHostNetwork")
    def block_access_host_network(self) -> bool:
        """
        If true, prevent containers from running with access to host network.
        """
        return pulumi.get(self, "block_access_host_network")

    @property
    @pulumi.getter(name="blockAddingCapabilities")
    def block_adding_capabilities(self) -> bool:
        """
        If true, prevent containers from running with adding capabilities with `--cap-add` privilege.
        """
        return pulumi.get(self, "block_adding_capabilities")

    @property
    @pulumi.getter(name="blockContainerExec")
    def block_container_exec(self) -> bool:
        """
        If true, exec into a container is prevented.
        """
        return pulumi.get(self, "block_container_exec")

    @property
    @pulumi.getter(name="blockCryptocurrencyMining")
    def block_cryptocurrency_mining(self) -> bool:
        """
        Detect and prevent communication to DNS/IP addresses known to be used for Cryptocurrency Mining
        """
        return pulumi.get(self, "block_cryptocurrency_mining")

    @property
    @pulumi.getter(name="blockFilelessExec")
    def block_fileless_exec(self) -> bool:
        """
        Detect and prevent running in-memory execution
        """
        return pulumi.get(self, "block_fileless_exec")

    @property
    @pulumi.getter(name="blockLowPortBinding")
    def block_low_port_binding(self) -> bool:
        """
        If true, prevent containers from running with the capability to bind in port lower than 1024.
        """
        return pulumi.get(self, "block_low_port_binding")

    @property
    @pulumi.getter(name="blockNonCompliantImages")
    def block_non_compliant_images(self) -> bool:
        """
        If true, running non-compliant image in the container is prevented.
        """
        return pulumi.get(self, "block_non_compliant_images")

    @property
    @pulumi.getter(name="blockNonCompliantWorkloads")
    def block_non_compliant_workloads(self) -> bool:
        """
        If true, running containers in non-compliant pods is prevented.
        """
        return pulumi.get(self, "block_non_compliant_workloads")

    @property
    @pulumi.getter(name="blockNonK8sContainers")
    def block_non_k8s_containers(self) -> bool:
        """
        If true, running non-kubernetes containers is prevented.
        """
        return pulumi.get(self, "block_non_k8s_containers")

    @property
    @pulumi.getter(name="blockPrivilegedContainers")
    def block_privileged_containers(self) -> bool:
        """
        If true, prevent containers from running with privileged container capability.
        """
        return pulumi.get(self, "block_privileged_containers")

    @property
    @pulumi.getter(name="blockReverseShell")
    def block_reverse_shell(self) -> bool:
        """
        If true, reverse shell is prevented.
        """
        return pulumi.get(self, "block_reverse_shell")

    @property
    @pulumi.getter(name="blockRootUser")
    def block_root_user(self) -> bool:
        """
        If true, prevent containers from running with root user.
        """
        return pulumi.get(self, "block_root_user")

    @property
    @pulumi.getter(name="blockUnregisteredImages")
    def block_unregistered_images(self) -> bool:
        """
        If true, running images in the container that are not registered in Aqua is prevented.
        """
        return pulumi.get(self, "block_unregistered_images")

    @property
    @pulumi.getter(name="blockUseIpcNamespace")
    def block_use_ipc_namespace(self) -> bool:
        """
        If true, prevent containers from running with the privilege to use the IPC namespace.
        """
        return pulumi.get(self, "block_use_ipc_namespace")

    @property
    @pulumi.getter(name="blockUsePidNamespace")
    def block_use_pid_namespace(self) -> bool:
        """
        If true, prevent containers from running with the privilege to use the PID namespace.
        """
        return pulumi.get(self, "block_use_pid_namespace")

    @property
    @pulumi.getter(name="blockUseUserNamespace")
    def block_use_user_namespace(self) -> bool:
        """
        If true, prevent containers from running with the privilege to use the user namespace.
        """
        return pulumi.get(self, "block_use_user_namespace")

    @property
    @pulumi.getter(name="blockUseUtsNamespace")
    def block_use_uts_namespace(self) -> bool:
        """
        If true, prevent containers from running with the privilege to use the UTS namespace.
        """
        return pulumi.get(self, "block_use_uts_namespace")

    @property
    @pulumi.getter(name="blockedCapabilities")
    def blocked_capabilities(self) -> Sequence[str]:
        """
        If true, prevents containers from using specific Unix capabilities.
        """
        return pulumi.get(self, "blocked_capabilities")

    @property
    @pulumi.getter(name="blockedExecutables")
    def blocked_executables(self) -> Sequence[str]:
        """
        List of executables that are prevented from running in containers.
        """
        return pulumi.get(self, "blocked_executables")

    @property
    @pulumi.getter(name="blockedFiles")
    def blocked_files(self) -> Sequence[str]:
        """
        List of files that are prevented from being read, modified and executed in the containers.
        """
        return pulumi.get(self, "blocked_files")

    @property
    @pulumi.getter(name="blockedInboundPorts")
    def blocked_inbound_ports(self) -> Sequence[str]:
        """
        List of blocked inbound ports.
        """
        return pulumi.get(self, "blocked_inbound_ports")

    @property
    @pulumi.getter(name="blockedOutboundPorts")
    def blocked_outbound_ports(self) -> Sequence[str]:
        """
        List of blocked outbound ports.
        """
        return pulumi.get(self, "blocked_outbound_ports")

    @property
    @pulumi.getter(name="blockedPackages")
    def blocked_packages(self) -> Sequence[str]:
        """
        Prevent containers from reading, writing, or executing all files in the list of packages.
        """
        return pulumi.get(self, "blocked_packages")

    @property
    @pulumi.getter(name="blockedVolumes")
    def blocked_volumes(self) -> Sequence[str]:
        """
        List of volumes that are prevented from being mounted in the containers.
        """
        return pulumi.get(self, "blocked_volumes")

    @property
    @pulumi.getter(name="containerExecAllowedProcesses")
    def container_exec_allowed_processes(self) -> Sequence[str]:
        """
        List of processes that will be allowed.
        """
        return pulumi.get(self, "container_exec_allowed_processes")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the container runtime policy
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableDriftPrevention")
    def enable_drift_prevention(self) -> bool:
        """
        If true, executables that are not in the original image is prevented from running.
        """
        return pulumi.get(self, "enable_drift_prevention")

    @property
    @pulumi.getter(name="enableForkGuard")
    def enable_fork_guard(self) -> bool:
        """
        If true, fork bombs are prevented in the containers.
        """
        return pulumi.get(self, "enable_fork_guard")

    @property
    @pulumi.getter(name="enableIpReputationSecurity")
    def enable_ip_reputation_security(self) -> bool:
        """
        If true, detect and prevent communication from containers to IP addresses known to have a bad reputation.
        """
        return pulumi.get(self, "enable_ip_reputation_security")

    @property
    @pulumi.getter(name="enablePortScanDetection")
    def enable_port_scan_detection(self) -> bool:
        """
        If true, detects port scanning behavior in the container.
        """
        return pulumi.get(self, "enable_port_scan_detection")

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
    @pulumi.getter(name="exceptionalReadonlyFilesAndDirectories")
    def exceptional_readonly_files_and_directories(self) -> Sequence[str]:
        """
        List of files and directories to be excluded from the read-only list.
        """
        return pulumi.get(self, "exceptional_readonly_files_and_directories")

    @property
    @pulumi.getter(name="fileIntegrityMonitorings")
    def file_integrity_monitorings(self) -> Sequence['outputs.GetContainerRuntimePolicyFileIntegrityMonitoringResult']:
        """
        Configuration for file integrity monitoring.
        """
        return pulumi.get(self, "file_integrity_monitorings")

    @property
    @pulumi.getter(name="forkGuardProcessLimit")
    def fork_guard_process_limit(self) -> int:
        """
        Process limit for the fork guard.
        """
        return pulumi.get(self, "fork_guard_process_limit")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="limitNewPrivileges")
    def limit_new_privileges(self) -> bool:
        """
        If true, prevents the container from obtaining new privileges at runtime. (only enabled in enforce mode)
        """
        return pulumi.get(self, "limit_new_privileges")

    @property
    @pulumi.getter(name="malwareScanOptions")
    def malware_scan_options(self) -> Sequence['outputs.GetContainerRuntimePolicyMalwareScanOptionResult']:
        """
        Configuration for Real-Time Malware Protection.
        """
        return pulumi.get(self, "malware_scan_options")

    @property
    @pulumi.getter(name="monitorSystemTimeChanges")
    def monitor_system_time_changes(self) -> bool:
        """
        If true, system time changes will be monitored.
        """
        return pulumi.get(self, "monitor_system_time_changes")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the container runtime policy
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="readonlyFilesAndDirectories")
    def readonly_files_and_directories(self) -> Sequence[str]:
        """
        List of files and directories to be restricted as read-only
        """
        return pulumi.get(self, "readonly_files_and_directories")

    @property
    @pulumi.getter(name="reverseShellAllowedIps")
    def reverse_shell_allowed_ips(self) -> Sequence[str]:
        """
        List of IPs/ CIDRs that will be allowed
        """
        return pulumi.get(self, "reverse_shell_allowed_ips")

    @property
    @pulumi.getter(name="reverseShellAllowedProcesses")
    def reverse_shell_allowed_processes(self) -> Sequence[str]:
        """
        List of processes that will be allowed
        """
        return pulumi.get(self, "reverse_shell_allowed_processes")

    @property
    @pulumi.getter(name="scopeExpression")
    def scope_expression(self) -> str:
        """
        Logical expression of how to compute the dependency of the scope variables.
        """
        return pulumi.get(self, "scope_expression")

    @property
    @pulumi.getter(name="scopeVariables")
    def scope_variables(self) -> Sequence['outputs.GetContainerRuntimePolicyScopeVariableResult']:
        """
        List of scope attributes.
        """
        return pulumi.get(self, "scope_variables")


class AwaitableGetContainerRuntimePolicyResult(GetContainerRuntimePolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContainerRuntimePolicyResult(
            allowed_executables=self.allowed_executables,
            allowed_registries=self.allowed_registries,
            application_scopes=self.application_scopes,
            audit_all_network_activity=self.audit_all_network_activity,
            audit_all_processes_activity=self.audit_all_processes_activity,
            audit_full_command_arguments=self.audit_full_command_arguments,
            author=self.author,
            block_access_host_network=self.block_access_host_network,
            block_adding_capabilities=self.block_adding_capabilities,
            block_container_exec=self.block_container_exec,
            block_cryptocurrency_mining=self.block_cryptocurrency_mining,
            block_fileless_exec=self.block_fileless_exec,
            block_low_port_binding=self.block_low_port_binding,
            block_non_compliant_images=self.block_non_compliant_images,
            block_non_compliant_workloads=self.block_non_compliant_workloads,
            block_non_k8s_containers=self.block_non_k8s_containers,
            block_privileged_containers=self.block_privileged_containers,
            block_reverse_shell=self.block_reverse_shell,
            block_root_user=self.block_root_user,
            block_unregistered_images=self.block_unregistered_images,
            block_use_ipc_namespace=self.block_use_ipc_namespace,
            block_use_pid_namespace=self.block_use_pid_namespace,
            block_use_user_namespace=self.block_use_user_namespace,
            block_use_uts_namespace=self.block_use_uts_namespace,
            blocked_capabilities=self.blocked_capabilities,
            blocked_executables=self.blocked_executables,
            blocked_files=self.blocked_files,
            blocked_inbound_ports=self.blocked_inbound_ports,
            blocked_outbound_ports=self.blocked_outbound_ports,
            blocked_packages=self.blocked_packages,
            blocked_volumes=self.blocked_volumes,
            container_exec_allowed_processes=self.container_exec_allowed_processes,
            description=self.description,
            enable_drift_prevention=self.enable_drift_prevention,
            enable_fork_guard=self.enable_fork_guard,
            enable_ip_reputation_security=self.enable_ip_reputation_security,
            enable_port_scan_detection=self.enable_port_scan_detection,
            enabled=self.enabled,
            enforce=self.enforce,
            enforce_after_days=self.enforce_after_days,
            exceptional_readonly_files_and_directories=self.exceptional_readonly_files_and_directories,
            file_integrity_monitorings=self.file_integrity_monitorings,
            fork_guard_process_limit=self.fork_guard_process_limit,
            id=self.id,
            limit_new_privileges=self.limit_new_privileges,
            malware_scan_options=self.malware_scan_options,
            monitor_system_time_changes=self.monitor_system_time_changes,
            name=self.name,
            readonly_files_and_directories=self.readonly_files_and_directories,
            reverse_shell_allowed_ips=self.reverse_shell_allowed_ips,
            reverse_shell_allowed_processes=self.reverse_shell_allowed_processes,
            scope_expression=self.scope_expression,
            scope_variables=self.scope_variables)


def get_container_runtime_policy(malware_scan_options: Optional[Sequence[pulumi.InputType['GetContainerRuntimePolicyMalwareScanOptionArgs']]] = None,
                                 name: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContainerRuntimePolicyResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_aquasec as aquasec

    container_runtime_policy = aquasec.get_container_runtime_policy(name="FunctionRuntimePolicyName")
    pulumi.export("containerRuntimePolicyDetails", container_runtime_policy)
    ```


    :param Sequence[pulumi.InputType['GetContainerRuntimePolicyMalwareScanOptionArgs']] malware_scan_options: Configuration for Real-Time Malware Protection.
    :param str name: Name of the container runtime policy
    """
    __args__ = dict()
    __args__['malwareScanOptions'] = malware_scan_options
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aquasec:index/getContainerRuntimePolicy:getContainerRuntimePolicy', __args__, opts=opts, typ=GetContainerRuntimePolicyResult).value

    return AwaitableGetContainerRuntimePolicyResult(
        allowed_executables=__ret__.allowed_executables,
        allowed_registries=__ret__.allowed_registries,
        application_scopes=__ret__.application_scopes,
        audit_all_network_activity=__ret__.audit_all_network_activity,
        audit_all_processes_activity=__ret__.audit_all_processes_activity,
        audit_full_command_arguments=__ret__.audit_full_command_arguments,
        author=__ret__.author,
        block_access_host_network=__ret__.block_access_host_network,
        block_adding_capabilities=__ret__.block_adding_capabilities,
        block_container_exec=__ret__.block_container_exec,
        block_cryptocurrency_mining=__ret__.block_cryptocurrency_mining,
        block_fileless_exec=__ret__.block_fileless_exec,
        block_low_port_binding=__ret__.block_low_port_binding,
        block_non_compliant_images=__ret__.block_non_compliant_images,
        block_non_compliant_workloads=__ret__.block_non_compliant_workloads,
        block_non_k8s_containers=__ret__.block_non_k8s_containers,
        block_privileged_containers=__ret__.block_privileged_containers,
        block_reverse_shell=__ret__.block_reverse_shell,
        block_root_user=__ret__.block_root_user,
        block_unregistered_images=__ret__.block_unregistered_images,
        block_use_ipc_namespace=__ret__.block_use_ipc_namespace,
        block_use_pid_namespace=__ret__.block_use_pid_namespace,
        block_use_user_namespace=__ret__.block_use_user_namespace,
        block_use_uts_namespace=__ret__.block_use_uts_namespace,
        blocked_capabilities=__ret__.blocked_capabilities,
        blocked_executables=__ret__.blocked_executables,
        blocked_files=__ret__.blocked_files,
        blocked_inbound_ports=__ret__.blocked_inbound_ports,
        blocked_outbound_ports=__ret__.blocked_outbound_ports,
        blocked_packages=__ret__.blocked_packages,
        blocked_volumes=__ret__.blocked_volumes,
        container_exec_allowed_processes=__ret__.container_exec_allowed_processes,
        description=__ret__.description,
        enable_drift_prevention=__ret__.enable_drift_prevention,
        enable_fork_guard=__ret__.enable_fork_guard,
        enable_ip_reputation_security=__ret__.enable_ip_reputation_security,
        enable_port_scan_detection=__ret__.enable_port_scan_detection,
        enabled=__ret__.enabled,
        enforce=__ret__.enforce,
        enforce_after_days=__ret__.enforce_after_days,
        exceptional_readonly_files_and_directories=__ret__.exceptional_readonly_files_and_directories,
        file_integrity_monitorings=__ret__.file_integrity_monitorings,
        fork_guard_process_limit=__ret__.fork_guard_process_limit,
        id=__ret__.id,
        limit_new_privileges=__ret__.limit_new_privileges,
        malware_scan_options=__ret__.malware_scan_options,
        monitor_system_time_changes=__ret__.monitor_system_time_changes,
        name=__ret__.name,
        readonly_files_and_directories=__ret__.readonly_files_and_directories,
        reverse_shell_allowed_ips=__ret__.reverse_shell_allowed_ips,
        reverse_shell_allowed_processes=__ret__.reverse_shell_allowed_processes,
        scope_expression=__ret__.scope_expression,
        scope_variables=__ret__.scope_variables)


@_utilities.lift_output_func(get_container_runtime_policy)
def get_container_runtime_policy_output(malware_scan_options: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetContainerRuntimePolicyMalwareScanOptionArgs']]]]] = None,
                                        name: Optional[pulumi.Input[str]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetContainerRuntimePolicyResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_aquasec as aquasec

    container_runtime_policy = aquasec.get_container_runtime_policy(name="FunctionRuntimePolicyName")
    pulumi.export("containerRuntimePolicyDetails", container_runtime_policy)
    ```


    :param Sequence[pulumi.InputType['GetContainerRuntimePolicyMalwareScanOptionArgs']] malware_scan_options: Configuration for Real-Time Malware Protection.
    :param str name: Name of the container runtime policy
    """
    ...
