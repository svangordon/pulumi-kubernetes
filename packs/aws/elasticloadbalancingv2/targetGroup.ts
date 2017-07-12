// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";

export class TargetGroup extends lumi.NamedResource implements TargetGroupArgs {
    public readonly arn?: string;
    public readonly arnSuffix?: string;
    public readonly deregistrationDelay?: number;
    public readonly healthCheck?: { healthyThreshold?: number, interval?: number, matcher?: string, path?: string, port?: string, protocol?: string, timeout?: number, unhealthyThreshold?: number }[];
    public readonly _name?: string;
    public readonly namePrefix?: string;
    public readonly port: number;
    public readonly protocol: string;
    public readonly stickiness?: { cookieDuration?: number, enabled?: boolean, type: string }[];
    public readonly tags?: {[key: string]: any};
    public readonly vpcId: string;

    constructor(name: string, args: TargetGroupArgs) {
        super(name);
        this.arn = args.arn;
        this.arnSuffix = args.arnSuffix;
        this.deregistrationDelay = args.deregistrationDelay;
        this.healthCheck = args.healthCheck;
        this._name = args._name;
        this.namePrefix = args.namePrefix;
        this.port = args.port;
        this.protocol = args.protocol;
        this.stickiness = args.stickiness;
        this.tags = args.tags;
        this.vpcId = args.vpcId;
    }
}

export interface TargetGroupArgs {
    readonly arn?: string;
    readonly arnSuffix?: string;
    readonly deregistrationDelay?: number;
    readonly healthCheck?: { healthyThreshold?: number, interval?: number, matcher?: string, path?: string, port?: string, protocol?: string, timeout?: number, unhealthyThreshold?: number }[];
    readonly _name?: string;
    readonly namePrefix?: string;
    readonly port: number;
    readonly protocol: string;
    readonly stickiness?: { cookieDuration?: number, enabled?: boolean, type: string }[];
    readonly tags?: {[key: string]: any};
    readonly vpcId: string;
}
