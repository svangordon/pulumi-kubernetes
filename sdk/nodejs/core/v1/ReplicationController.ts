// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";

    /**
     * ReplicationController represents the configuration of a replication controller.
     */
    export class ReplicationController extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"v1">;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"ReplicationController">;

      /**
       * If the Labels of a ReplicationController are empty, they are defaulted to be the same as
       * the Pod(s) that the replication controller manages. Standard object's metadata. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ObjectMeta>;

      /**
       * Spec defines the specification of the desired behavior of the replication controller. More
       * info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
       */
      public readonly spec: pulumi.Output<outputApi.core.v1.ReplicationControllerSpec>;

      /**
       * Status is the most recently observed status of the replication controller. This data may be
       * out of date by some window of time. Populated by the system. Read-only. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
       */
      public readonly status: pulumi.Output<outputApi.core.v1.ReplicationControllerStatus>;

      /**
       * Get the state of an existing `ReplicationController` resource, as identified by `id`.
       * Typically this ID  is of the form <namespace>/<name>; if <namespace> is omitted, then (per
       * Kubernetes convention) the ID becomes default/<name>.
       *
       * Pulumi will keep track of this resource using `name` as the Pulumi ID.
       *
       * @param name _Unique_ name used to register this resource with Pulumi.
       * @param id An ID for the Kubernetes resource to retrieve. Takes the form
       *  <namespace>/<name> or <name>.
       * @param opts Uniquely specifies a CustomResource to select.
       */
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReplicationController {
          return new ReplicationController(name, undefined, { ...opts, id: id });
      }

      /**
       * Create a core.v1.ReplicationController resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.core.v1.ReplicationController, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "v1";
          inputs["kind"] = "ReplicationController";
          inputs["metadata"] = args && args.metadata || undefined;
          inputs["spec"] = args && args.spec || undefined;
          inputs["status"] = args && args.status || undefined;
          super("kubernetes:core/v1:ReplicationController", name, inputs, opts);
      }
    }
