// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";

export class ResourceGroup extends lumi.NamedResource implements ResourceGroupArgs {
    public readonly arn?: string;
    public readonly tags: {[key: string]: any};

    constructor(name: string, args: ResourceGroupArgs) {
        super(name);
        this.arn = args.arn;
        this.tags = args.tags;
    }
}

export interface ResourceGroupArgs {
    readonly arn?: string;
    readonly tags: {[key: string]: any};
}
