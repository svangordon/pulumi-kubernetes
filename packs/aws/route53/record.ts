// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";

export class Record extends lumi.NamedResource implements RecordArgs {
    public readonly alias?: { evaluateTargetHealth: boolean, _name: string, zoneId: string }[];
    public readonly failoverRoutingPolicy?: { type: string }[];
    public readonly fqdn?: string;
    public readonly geolocationRoutingPolicy?: { continent?: string, country?: string, subdivision?: string }[];
    public readonly healthCheckId?: string;
    public readonly latencyRoutingPolicy?: { region: string }[];
    public readonly _name: string;
    public readonly records?: string[];
    public readonly setIdentifier?: string;
    public readonly ttl?: number;
    public readonly type: string;
    public readonly weightedRoutingPolicy?: { weight: number }[];
    public readonly zoneId: string;

    constructor(name: string, args: RecordArgs) {
        super(name);
        this.alias = args.alias;
        this.failoverRoutingPolicy = args.failoverRoutingPolicy;
        this.fqdn = args.fqdn;
        this.geolocationRoutingPolicy = args.geolocationRoutingPolicy;
        this.healthCheckId = args.healthCheckId;
        this.latencyRoutingPolicy = args.latencyRoutingPolicy;
        this._name = args._name;
        this.records = args.records;
        this.setIdentifier = args.setIdentifier;
        this.ttl = args.ttl;
        this.type = args.type;
        this.weightedRoutingPolicy = args.weightedRoutingPolicy;
        this.zoneId = args.zoneId;
    }
}

export interface RecordArgs {
    readonly alias?: { evaluateTargetHealth: boolean, _name: string, zoneId: string }[];
    readonly failoverRoutingPolicy?: { type: string }[];
    readonly fqdn?: string;
    readonly geolocationRoutingPolicy?: { continent?: string, country?: string, subdivision?: string }[];
    readonly healthCheckId?: string;
    readonly latencyRoutingPolicy?: { region: string }[];
    readonly _name: string;
    readonly records?: string[];
    readonly setIdentifier?: string;
    readonly ttl?: number;
    readonly type: string;
    readonly weightedRoutingPolicy?: { weight: number }[];
    readonly zoneId: string;
}
