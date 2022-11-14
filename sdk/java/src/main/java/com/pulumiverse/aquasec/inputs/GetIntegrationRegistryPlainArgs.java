// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.aquasec.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIntegrationRegistryPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIntegrationRegistryPlainArgs Empty = new GetIntegrationRegistryPlainArgs();

    /**
     * Additional condition for pulling and rescanning images, Defaults to &#39;none&#39;
     * 
     */
    @Import(name="imageCreationDateCondition")
    private @Nullable String imageCreationDateCondition;

    /**
     * @return Additional condition for pulling and rescanning images, Defaults to &#39;none&#39;
     * 
     */
    public Optional<String> imageCreationDateCondition() {
        return Optional.ofNullable(this.imageCreationDateCondition);
    }

    /**
     * The name of the registry; string, required - this will be treated as the registry&#39;s ID, so choose a simple alphanumerical name without special signs and spaces
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The name of the registry; string, required - this will be treated as the registry&#39;s ID, so choose a simple alphanumerical name without special signs and spaces
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * When auto pull image enabled, sets maximum age of auto pulled images
     * 
     */
    @Import(name="pullImageAge")
    private @Nullable String pullImageAge;

    /**
     * @return When auto pull image enabled, sets maximum age of auto pulled images
     * 
     */
    public Optional<String> pullImageAge() {
        return Optional.ofNullable(this.pullImageAge);
    }

    /**
     * When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
     * 
     */
    @Import(name="pullImageCount")
    private @Nullable Integer pullImageCount;

    /**
     * @return When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
     * 
     */
    public Optional<Integer> pullImageCount() {
        return Optional.ofNullable(this.pullImageCount);
    }

    /**
     * List of scanner names
     * 
     */
    @Import(name="scannerNames")
    private @Nullable List<String> scannerNames;

    /**
     * @return List of scanner names
     * 
     */
    public Optional<List<String>> scannerNames() {
        return Optional.ofNullable(this.scannerNames);
    }

    /**
     * Scanner type
     * 
     */
    @Import(name="scannerType")
    private @Nullable String scannerType;

    /**
     * @return Scanner type
     * 
     */
    public Optional<String> scannerType() {
        return Optional.ofNullable(this.scannerType);
    }

    private GetIntegrationRegistryPlainArgs() {}

    private GetIntegrationRegistryPlainArgs(GetIntegrationRegistryPlainArgs $) {
        this.imageCreationDateCondition = $.imageCreationDateCondition;
        this.name = $.name;
        this.pullImageAge = $.pullImageAge;
        this.pullImageCount = $.pullImageCount;
        this.scannerNames = $.scannerNames;
        this.scannerType = $.scannerType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIntegrationRegistryPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIntegrationRegistryPlainArgs $;

        public Builder() {
            $ = new GetIntegrationRegistryPlainArgs();
        }

        public Builder(GetIntegrationRegistryPlainArgs defaults) {
            $ = new GetIntegrationRegistryPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param imageCreationDateCondition Additional condition for pulling and rescanning images, Defaults to &#39;none&#39;
         * 
         * @return builder
         * 
         */
        public Builder imageCreationDateCondition(@Nullable String imageCreationDateCondition) {
            $.imageCreationDateCondition = imageCreationDateCondition;
            return this;
        }

        /**
         * @param name The name of the registry; string, required - this will be treated as the registry&#39;s ID, so choose a simple alphanumerical name without special signs and spaces
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param pullImageAge When auto pull image enabled, sets maximum age of auto pulled images
         * 
         * @return builder
         * 
         */
        public Builder pullImageAge(@Nullable String pullImageAge) {
            $.pullImageAge = pullImageAge;
            return this;
        }

        /**
         * @param pullImageCount When auto pull image enabled, sets maximum age of auto pulled images tags from each repository.
         * 
         * @return builder
         * 
         */
        public Builder pullImageCount(@Nullable Integer pullImageCount) {
            $.pullImageCount = pullImageCount;
            return this;
        }

        /**
         * @param scannerNames List of scanner names
         * 
         * @return builder
         * 
         */
        public Builder scannerNames(@Nullable List<String> scannerNames) {
            $.scannerNames = scannerNames;
            return this;
        }

        /**
         * @param scannerNames List of scanner names
         * 
         * @return builder
         * 
         */
        public Builder scannerNames(String... scannerNames) {
            return scannerNames(List.of(scannerNames));
        }

        /**
         * @param scannerType Scanner type
         * 
         * @return builder
         * 
         */
        public Builder scannerType(@Nullable String scannerType) {
            $.scannerType = scannerType;
            return this;
        }

        public GetIntegrationRegistryPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
