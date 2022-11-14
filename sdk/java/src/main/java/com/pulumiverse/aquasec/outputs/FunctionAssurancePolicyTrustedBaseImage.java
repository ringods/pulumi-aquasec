// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.aquasec.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FunctionAssurancePolicyTrustedBaseImage {
    private @Nullable String imagename;
    private @Nullable String registry;

    private FunctionAssurancePolicyTrustedBaseImage() {}
    public Optional<String> imagename() {
        return Optional.ofNullable(this.imagename);
    }
    public Optional<String> registry() {
        return Optional.ofNullable(this.registry);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FunctionAssurancePolicyTrustedBaseImage defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String imagename;
        private @Nullable String registry;
        public Builder() {}
        public Builder(FunctionAssurancePolicyTrustedBaseImage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.imagename = defaults.imagename;
    	      this.registry = defaults.registry;
        }

        @CustomType.Setter
        public Builder imagename(@Nullable String imagename) {
            this.imagename = imagename;
            return this;
        }
        @CustomType.Setter
        public Builder registry(@Nullable String registry) {
            this.registry = registry;
            return this;
        }
        public FunctionAssurancePolicyTrustedBaseImage build() {
            final var o = new FunctionAssurancePolicyTrustedBaseImage();
            o.imagename = imagename;
            o.registry = registry;
            return o;
        }
    }
}
