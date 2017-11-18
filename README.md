# CRI Proxy

CRI Proxy makes it possible to run several CRI implementations on the
same node. It also includes `dockershim` (CRI->docker interface) which
is currently imported from Kubernetes code. CRI Proxy isn't tied to
Virtlet in the sense that it can be used with other runtimes,
too. `dockershim` usage is also optional.

## How CRI Proxy works

Below is a diagram depicting the way CRI Proxy works. The basic idea
is forwarding the requests to different runtimes based on prefixes of
image name / pod id / container id prefixes.

![CRI Request Path](criproxy.png)

Let's say CRI proxy is started as follows:
```
/usr/local/bin/criproxy -v 3 -alsologtostderr -connect docker,virtlet:/run/virtlet.sock
```

`-v 3 -alsologtostderr` options here may be quite useful for
debugging, because they make CRI proxy log detailed info about every
CRI request going through it, including any errors and the result.

`-connect docker,virtlet:/run/virtlet.sock` specifies the list of
runtimes that the proxy passes requests to.

The `docker` part is a special case, meaning that `criproxy` must
start in-process `dockershim` and use it as the primary (prefixless)
runtime (*TODO:* should not use this). It's also possible to specify
other primary runtime instead, e.g. `/run/some-other-runtime.sock`.

`virtlet.cloud:/run/virtlet.sock` denotes an alternative runtime
socket. This means that image service requests that include image
names starting with `virtlet.cloud/` must be directed to the CRI
implementation listening on a Unix domain socket at
`/run/virtlet.sock`. Pods that need to run on `virtlet.cloud` runtime must
have `virtlet.cloud` as the value of `kubernetes.io/target-runtime`
annotation.

There can be any number of runtimes, although probably using more than
a couple of runtimes is a rare use case.

Here's an example of a pod that needs to run on `virtlet.cloud` runtime:
```
apiVersion: v1
kind: Pod
metadata:
  name: cirros-vm
  annotations:
    kubernetes.io/target-runtime: virtlet.cloud
    scheduler.alpha.kubernetes.io/affinity: >
      {
        "nodeAffinity": {
          "requiredDuringSchedulingIgnoredDuringExecution": {
            "nodeSelectorTerms": [
              {
                "matchExpressions": [
                  {
                    "key": "extraRuntime",
                    "operator": "In",
                    "values": ["virtlet"]
                  }
                ]
              }
            ]
          }
        }
      }
spec:
  containers:
    - name: cirros-vm
      image: virtlet.cloud/image-service/cirros
```

First of all, there's `kubernetes.io/target-runtime: virtlet.cloud`
annotation that directs `RunPodSandbox` requests to `virtlet.cloud` runtime.

There's also `nodeAffinity` spec that makes the pod run only on the
nodes that have `extraRuntime=virtlet` label. This is not required
by CRI proxy mechanism itself and is related to deployment mechanism
being used.

Another important part is `virtlet.cloud/image-service/cirros` image name.
It means that the image is handled by `virtlet.cloud` runtime and actual
image name passed to the runtime is `image-service/cirros`. In case of
virtlet this means downloading QCOW2 image from
`http://image-service/cirros`.

In order to distinguish between runtimes during requests that don't
include image name or pod annotations such as `RemovePodSandbox`, CRI
proxy adds prefixes to pod and container ids returned by the runtimes.
