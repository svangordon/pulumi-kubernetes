// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";

export class Eip extends lumi.NamedResource implements EipArgs {
    public readonly allocationId?: string;
    public readonly associateWithPrivateIp?: string;
    public readonly associationId?: string;
    public readonly domain?: string;
    public readonly instance?: string;
    public readonly networkInterface?: string;
    public readonly privateIp?: string;
    public readonly publicIp?: string;
    public readonly vpc?: boolean;

    constructor(name: string, args: EipArgs) {
        super(name);
        this.allocationId = args.allocationId;
        this.associateWithPrivateIp = args.associateWithPrivateIp;
        this.associationId = args.associationId;
        this.domain = args.domain;
        this.instance = args.instance;
        this.networkInterface = args.networkInterface;
        this.privateIp = args.privateIp;
        this.publicIp = args.publicIp;
        this.vpc = args.vpc;
    }
}

export interface EipArgs {
    readonly allocationId?: string;
    readonly associateWithPrivateIp?: string;
    readonly associationId?: string;
    readonly domain?: string;
    readonly instance?: string;
    readonly networkInterface?: string;
    readonly privateIp?: string;
    readonly publicIp?: string;
    readonly vpc?: boolean;
}
