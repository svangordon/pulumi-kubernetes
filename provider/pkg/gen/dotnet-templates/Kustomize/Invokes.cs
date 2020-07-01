// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Kubernetes.Kustomize
{
    internal static class Invokes
    {
        /// <summary>
        /// Invoke the resource provider to process a kustomization.
        /// </summary>
        internal static Output<ImmutableArray<ImmutableDictionary<string, object>>> KustomizeDirectory(KustomizeDirectoryArgs args,
            InvokeOptions? options = null)
            => Output.Create(Deployment.Instance.InvokeAsync<KustomizeDirectoryResult>("kubernetes:kustomize:directory", args,
                options.WithVersion())).Apply(r => r.Result.ToImmutableArray());
    }

    internal class KustomizeDirectoryArgs : InvokeArgs
    {
        [Input("directory")]
        public string? Directory { get; set; }
    }

    [OutputType]
    internal class KustomizeDirectoryResult
    {
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Result;

        [OutputConstructor]
        private KustomizeDirectoryResult(
            ImmutableArray<ImmutableDictionary<string, object>> result)
        {
            Result = result;
        }
    }
}
