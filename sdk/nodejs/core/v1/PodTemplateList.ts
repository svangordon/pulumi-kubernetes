// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";
import { getVersion } from "../../version";

    /**
     * PodTemplateList is a list of PodTemplates.
     */
    export class PodTemplateList extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"v1">;

      /**
       * List of pod templates
       */
      public readonly items: pulumi.Output<outputApi.core.v1.PodTemplate[]>;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"PodTemplateList">;

      /**
       * Standard list metadata. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ListMeta>;

      /**
       * Get the state of an existing `PodTemplateList` resource, as identified by `id`.
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
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PodTemplateList {
          return new PodTemplateList(name, undefined, { ...opts, id: id });
      }

      /** @internal */
      private static readonly __pulumiType = "kubernetes:core/v1:PodTemplateList";

      /**
       * Returns true if the given object is an instance of PodTemplateList.  This is designed to work even
       * when multiple copies of the Pulumi SDK have been loaded into the same process.
       */
      public static isInstance(obj: any): obj is PodTemplateList {
          if (obj === undefined || obj === null) {
              return false;
          }

          return obj["__pulumiType"] === PodTemplateList.__pulumiType;
      }

      /**
       * Create a core.v1.PodTemplateList resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.core.v1.PodTemplateList, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "v1";
          inputs["items"] = args && args.items || undefined;
          inputs["kind"] = "PodTemplateList";
          inputs["metadata"] = args && args.metadata || undefined;

          if (!opts) {
              opts = {};
          }

          if (!opts.version) {
              opts.version = getVersion();
          }
          super(PodTemplateList.__pulumiType, name, inputs, opts);
      }
    }
