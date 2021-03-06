// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1
{

    [OutputType]
    public sealed class HostPortRange
    {
        /// <summary>
        /// max is the end of the range, inclusive.
        /// </summary>
        public readonly int Max;
        /// <summary>
        /// min is the start of the range, inclusive.
        /// </summary>
        public readonly int Min;

        [OutputConstructor]
        private HostPortRange(
            int max,

            int min)
        {
            Max = max;
            Min = min;
        }
    }
}
