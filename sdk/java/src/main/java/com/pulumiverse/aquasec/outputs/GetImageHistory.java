// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.aquasec.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetImageHistory {
    /**
     * @return The image creation comment.
     * 
     */
    private String comment;
    /**
     * @return The date and time when the image was registered.
     * 
     */
    private String created;
    private String createdBy;
    /**
     * @return The ID of this resource.
     * 
     */
    private String id;
    private Integer size;

    private GetImageHistory() {}
    /**
     * @return The image creation comment.
     * 
     */
    public String comment() {
        return this.comment;
    }
    /**
     * @return The date and time when the image was registered.
     * 
     */
    public String created() {
        return this.created;
    }
    public String createdBy() {
        return this.createdBy;
    }
    /**
     * @return The ID of this resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Integer size() {
        return this.size;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetImageHistory defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String comment;
        private String created;
        private String createdBy;
        private String id;
        private Integer size;
        public Builder() {}
        public Builder(GetImageHistory defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.comment = defaults.comment;
    	      this.created = defaults.created;
    	      this.createdBy = defaults.createdBy;
    	      this.id = defaults.id;
    	      this.size = defaults.size;
        }

        @CustomType.Setter
        public Builder comment(String comment) {
            this.comment = Objects.requireNonNull(comment);
            return this;
        }
        @CustomType.Setter
        public Builder created(String created) {
            this.created = Objects.requireNonNull(created);
            return this;
        }
        @CustomType.Setter
        public Builder createdBy(String createdBy) {
            this.createdBy = Objects.requireNonNull(createdBy);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder size(Integer size) {
            this.size = Objects.requireNonNull(size);
            return this;
        }
        public GetImageHistory build() {
            final var o = new GetImageHistory();
            o.comment = comment;
            o.created = created;
            o.createdBy = createdBy;
            o.id = id;
            o.size = size;
            return o;
        }
    }
}
