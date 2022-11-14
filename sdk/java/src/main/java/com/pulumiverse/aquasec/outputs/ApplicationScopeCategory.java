// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.aquasec.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumiverse.aquasec.outputs.ApplicationScopeCategoryArtifact;
import com.pulumiverse.aquasec.outputs.ApplicationScopeCategoryEntityScope;
import com.pulumiverse.aquasec.outputs.ApplicationScopeCategoryInfrastructure;
import com.pulumiverse.aquasec.outputs.ApplicationScopeCategoryWorkload;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ApplicationScopeCategory {
    /**
     * @return An artifact is an application. It can be an image (for a container, not a CF application); a serverless function; or a Tanzu Application Service (TAS) droplet.
     * 
     */
    private @Nullable List<ApplicationScopeCategoryArtifact> artifacts;
    private @Nullable List<ApplicationScopeCategoryEntityScope> entityScopes;
    /**
     * @return An infrastructure resource is an element of a computing environment on which a workload is orchestrated and run. It can be a host (VM) or a Kubernetes cluster.
     * 
     */
    private @Nullable List<ApplicationScopeCategoryInfrastructure> infrastructures;
    /**
     * @return A workload is a running container. It can run in a Kubernetes cluster, on a VM (no orchestrator), or under Tanzu Application Service (TAS).
     * 
     */
    private @Nullable List<ApplicationScopeCategoryWorkload> workloads;

    private ApplicationScopeCategory() {}
    /**
     * @return An artifact is an application. It can be an image (for a container, not a CF application); a serverless function; or a Tanzu Application Service (TAS) droplet.
     * 
     */
    public List<ApplicationScopeCategoryArtifact> artifacts() {
        return this.artifacts == null ? List.of() : this.artifacts;
    }
    public List<ApplicationScopeCategoryEntityScope> entityScopes() {
        return this.entityScopes == null ? List.of() : this.entityScopes;
    }
    /**
     * @return An infrastructure resource is an element of a computing environment on which a workload is orchestrated and run. It can be a host (VM) or a Kubernetes cluster.
     * 
     */
    public List<ApplicationScopeCategoryInfrastructure> infrastructures() {
        return this.infrastructures == null ? List.of() : this.infrastructures;
    }
    /**
     * @return A workload is a running container. It can run in a Kubernetes cluster, on a VM (no orchestrator), or under Tanzu Application Service (TAS).
     * 
     */
    public List<ApplicationScopeCategoryWorkload> workloads() {
        return this.workloads == null ? List.of() : this.workloads;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationScopeCategory defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<ApplicationScopeCategoryArtifact> artifacts;
        private @Nullable List<ApplicationScopeCategoryEntityScope> entityScopes;
        private @Nullable List<ApplicationScopeCategoryInfrastructure> infrastructures;
        private @Nullable List<ApplicationScopeCategoryWorkload> workloads;
        public Builder() {}
        public Builder(ApplicationScopeCategory defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.artifacts = defaults.artifacts;
    	      this.entityScopes = defaults.entityScopes;
    	      this.infrastructures = defaults.infrastructures;
    	      this.workloads = defaults.workloads;
        }

        @CustomType.Setter
        public Builder artifacts(@Nullable List<ApplicationScopeCategoryArtifact> artifacts) {
            this.artifacts = artifacts;
            return this;
        }
        public Builder artifacts(ApplicationScopeCategoryArtifact... artifacts) {
            return artifacts(List.of(artifacts));
        }
        @CustomType.Setter
        public Builder entityScopes(@Nullable List<ApplicationScopeCategoryEntityScope> entityScopes) {
            this.entityScopes = entityScopes;
            return this;
        }
        public Builder entityScopes(ApplicationScopeCategoryEntityScope... entityScopes) {
            return entityScopes(List.of(entityScopes));
        }
        @CustomType.Setter
        public Builder infrastructures(@Nullable List<ApplicationScopeCategoryInfrastructure> infrastructures) {
            this.infrastructures = infrastructures;
            return this;
        }
        public Builder infrastructures(ApplicationScopeCategoryInfrastructure... infrastructures) {
            return infrastructures(List.of(infrastructures));
        }
        @CustomType.Setter
        public Builder workloads(@Nullable List<ApplicationScopeCategoryWorkload> workloads) {
            this.workloads = workloads;
            return this;
        }
        public Builder workloads(ApplicationScopeCategoryWorkload... workloads) {
            return workloads(List.of(workloads));
        }
        public ApplicationScopeCategory build() {
            final var o = new ApplicationScopeCategory();
            o.artifacts = artifacts;
            o.entityScopes = entityScopes;
            o.infrastructures = infrastructures;
            o.workloads = workloads;
            return o;
        }
    }
}
