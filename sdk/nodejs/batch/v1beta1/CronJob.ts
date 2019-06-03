// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";

    /**
     * CronJob represents the configuration of a single cron job.
     */
    export class CronJob extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"batch/v1beta1">;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"CronJob">;

      /**
       * Standard object's metadata. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ObjectMeta>;

      /**
       * Specification of the desired behavior of a cron job, including the schedule. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
       */
      public readonly spec: pulumi.Output<outputApi.batch.v1beta1.CronJobSpec>;

      /**
       * Current status of a cron job. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
       */
      public readonly status: pulumi.Output<outputApi.batch.v1beta1.CronJobStatus>;

      /**
       * Get the state of an existing `CronJob` resource, as identified by `id`.
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
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CronJob {
          return new CronJob(name, undefined, { ...opts, id: id });
      }

      /**
       * Create a batch.v1beta1.CronJob resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.batch.v1beta1.CronJob, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "batch/v1beta1";
          inputs["kind"] = "CronJob";
          inputs["metadata"] = args && args.metadata || undefined;
          inputs["spec"] = args && args.spec || undefined;
          inputs["status"] = args && args.status || undefined;
          super("kubernetes:batch/v1beta1:CronJob", name, inputs, opts);
      }
    }
