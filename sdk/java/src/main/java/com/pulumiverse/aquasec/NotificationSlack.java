// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.aquasec;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumiverse.aquasec.NotificationSlackArgs;
import com.pulumiverse.aquasec.Utilities;
import com.pulumiverse.aquasec.inputs.NotificationSlackState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Aquasec Notification Slack resource
 * 
 * &gt; **Note about resource deprecation**
 * Resource aquasec.NotificationSlack is deprecated, please use aquasec.Notification instead
 * 
 */
@ResourceType(type="aquasec:index/notificationSlack:NotificationSlack")
public class NotificationSlack extends com.pulumi.resources.CustomResource {
    @Export(name="channel", refs={String.class}, tree="[0]")
    private Output<String> channel;

    public Output<String> channel() {
        return this.channel;
    }
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    public Output<Boolean> enabled() {
        return this.enabled;
    }
    @Export(name="icon", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> icon;

    public Output<Optional<String>> icon() {
        return Codegen.optional(this.icon);
    }
    @Export(name="mainText", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mainText;

    public Output<Optional<String>> mainText() {
        return Codegen.optional(this.mainText);
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="serviceKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serviceKey;

    public Output<Optional<String>> serviceKey() {
        return Codegen.optional(this.serviceKey);
    }
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    public Output<String> type() {
        return this.type;
    }
    @Export(name="userName", refs={String.class}, tree="[0]")
    private Output<String> userName;

    public Output<String> userName() {
        return this.userName;
    }
    @Export(name="webhookUrl", refs={String.class}, tree="[0]")
    private Output<String> webhookUrl;

    public Output<String> webhookUrl() {
        return this.webhookUrl;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NotificationSlack(String name) {
        this(name, NotificationSlackArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NotificationSlack(String name, NotificationSlackArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NotificationSlack(String name, NotificationSlackArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aquasec:index/notificationSlack:NotificationSlack", name, args == null ? NotificationSlackArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NotificationSlack(String name, Output<String> id, @Nullable NotificationSlackState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aquasec:index/notificationSlack:NotificationSlack", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static NotificationSlack get(String name, Output<String> id, @Nullable NotificationSlackState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NotificationSlack(name, id, state, options);
    }
}
