# CRI Proxy [![CircleCI](https://circleci.com/gh/Mirantis/criproxy/tree/master.svg?style=svg)](https://circleci.com/gh/Mirantis/criproxy/tree/master)

CRI Proxy makes it possible to run several CRI implementations on the
same node and run CRI implementations inside pods. CRI Proxy is
currently used by [Virtlet](https://github.com/Mirantis/virtlet)
project but it can be used by other CRI implementations, too.

It supports Kubernetes versions 1.8.x, 1.9.x, 1.10.x and 1.11.x.

## Installation on kubeadm clusters for use with Virtlet

In order to install CRI proxy on a kubeadm-deployed Kubernetes node
running Ubuntu 16.04, you can go to
[Releases](https://github.com/Mirantis/criproxy/releases) page of the
project, download latest `criproxy_X.Y.Z_amd64.deb` to the node
and install it with dpkg:
```
dpkg -i criproxy_X.Y.Z_amd64.deb
```

If you're using
[kubeadm-dind-cluster](https://github.com/Mirantis/kubeadm-dind-cluster)
to run your test cluster, you can download
`criproxy-nodeps_X.Y.Z_amd64.deb` to a DIND node using e.g. `docker
exec -it kube-node-XX ...` and install it there with `dpkg -i`.

The packages are currently tailored to support Virtlet deployments
but later they can be made more universal.

## Installation on a systemd-based system in a general case

In order to install the CRI proxy manually, first you need to download
`criproxy` binary from
[Releases](https://github.com/Mirantis/criproxy/releases) page of the
project and place it under `/usr/local/bin`.

After that you need to configure dockershim service. Dockershim can be
run using the same binary as kubelet with one catch: Kubernetes 1.8
prior to 1.8.4 don't support `kubelet --experimental-dockershim` in
`hyperkube`. If that's your case, you can either upgrade Kubernetes
binaries on your node or download `kubelet` binary for your version from
`https://storage.googleapis.com/kubernetes-release/release/vX.Y.Z/bin/linux/amd64/kubelet`
(put your version instead of vX.Y.Z) and use it for dockershim.

The idea is to start kubelet with extra flags
`--experimental-dockershim --port 11250`, other flags being the same
as those used for kubelet except for `--container-runtime`,
`--container-runtime-endpoint`, `--image-service-endpoint` and
`--port`.  More precisely, you need to pass the following flags
(others, besides container runtime and port options mentioned above,
are ignored):
* `--network-plugin`
* `--hairpin-mode`
* `--non-masquerade-cidr`
* `--cni-conf-dir`
* `--cni-bin-dir`
* `--docker-endpoint`
* `--runtime-request-timeout`
* `--image-pull-progress-deadline`
* `--streaming-connection-idle-timeout`
* `--docker-exec-handler`
* `--seccomp-profile-root`
* `--pod-infra-container-image`
* `--runtime-cgroups`
* `--cgroup-driver`
* `--network-plugin-mtu`
* `--address`

Create a file named `/etc/systemd/system/dockershim.service` with the
following content, replacing `......` with kubelet command line
arguments (a naive way to get them is just to do `ps aux|grep kubelet`
if you have `kubelet` service running):

```ini
[Unit]
Description=dockershim for criproxy

[Service]
ExecStart=/usr/bin/kubelet --experimental-dockershim --port 11250 ......
Restart=always
StartLimitInterval=0
RestartSec=10

[Install]
RequiredBy=criproxy.service
```

`--port 11250` specifies streaming port to use (it's used for things
like `kubectl attach`). If you use another port, you'll also need to
set `-streamPort XX` option for `criproxy`. If you get errors when
trying to do `kubectl exec`, `kubectl attach` or `kubectl port-forward`
on the node with CRI proxy, this means that CRI proxy fails to
determine node address properly and you need to pass `-streamUrl`
option to `criproxy`, e.g. `-streamUrl http://node-ip-address:11250/`.
This commonly happens when `--address` flag is passed to kubelet.
Note that `-streamPort` is ignored if `-streamUrl` is set.

Create a file named `/etc/systemd/system/criproxy.service` with
the following content (you can also use `systemctl --force edit criproxy.service` for it):

```ini
[Unit]
Description=CRI Proxy

[Service]
ExecStart=/usr/local/bin/criproxy -v 3 -logtostderr -connect /var/run/dockershim.sock,virtlet.cloud:/run/virtlet.sock -listen /run/criproxy.sock
Restart=always
StartLimitInterval=0
RestartSec=10

[Install]
WantedBy=kubelet.service
```

You can remove `-v 3` option to reduce verbosity level of the proxy.

Then enable and start the units after stopping kubelet:
```bash
systemctl stop kubelet
systemctl daemon-reload
systemctl enable criproxy dockershim
systemctl start criproxy dockershim
```

Then we need to reconfigure kubelet. You need to pass the following extra flags to it
to make it use CRI Proxy (you will need to do another `systemctl daemon-reload`):
```bash
--container-runtime=remote \
--container-runtime-endpoint=unix:///run/criproxy.sock \
--image-service-endpoint=unix:///run/criproxy.sock \
--enable-controller-attach-detach=false
```

## How CRI Proxy works

Below is a diagram depicting the way CRI Proxy works. The basic idea
is forwarding the requests to different runtimes based on prefixes of
image name / pod id / container id prefixes.

![CRI Request Path](criproxy.png)

Let's say CRI proxy is started as follows:
```
/usr/bin/criproxy -v 3 -logtostderr -connect /var/run/dockershim.sock,virtlet.cloud:/run/virtlet.sock -listen /run/criproxy.sock
```

`-v` option of `criproxy` controls the verbosity here. 0-1 means some
very basic logging during startup and displaying serious errors, 2 is
the same as 1 plus logging of CRI request errors and 3 causes dumping
of actual CRI requests and responses except for noisy `List*` and pod
container/status requests in addition to what's logged on level 2, so
on level 3 most of the output consists mostly of requests that change
the state of the runtime. Level 4 enables dumping of pod / container
status requests. Level 5 adds dumping `List*` requests which may cause
the log to grow fast. See
[fixing log throttling](#fixing-log-throttling) below if you're
starting CRI proxy using systemd with log level set to 3 or higher.

`-logtostderr` directs logging output to stderr (it's part of glog configuration)

`-connect /var/run/dockershim.sock,virtlet.cloud:/run/virtlet.sock` specifies the list of
runtimes that the proxy passes requests to.

`/var/run/dockershim.sock` is a primary runtime that will handle
unprefixed images and pods without `kubernetes.io/target-runtime`
annotation.

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
    affinity:
      nodeAffinity:
        requiredDuringSchedulingIgnoredDuringExecution:
          nodeSelectorTerms:
          - matchExpressions:
            - key: extraRuntime
              operator: In
              values:
              - virtlet
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

## <a name="fixing-log-throttling"></a>Fixing log throttling

If you're using log level 3 or higher, journald may throttle CRI Proxy
logs. If this is the case, you'll see some delays when checking CRI
Proxy logs with commands like `journalctl -xef -u criproxy`. In order
to work around this issue, you need to disable journald's rate
limiter. To do so, add the following line to
`/etc/systemd/journald.conf` (if it already has `RateLimitInterval=...`
line, you just need to change the value to 0)

```
RateLimitInterval=0
```

And then restart journald:
```
systemctl restart systemd-journald
```

See
[this post](https://www.rootusers.com/how-to-change-log-rate-limiting-in-linux/)
for more info on the issue.
