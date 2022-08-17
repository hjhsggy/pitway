## **Kubernetes**

- 解决什么问题?
- 编排? 调度? 容器云? 集群管理?

## 架构

![img](https://static001.geekbang.org/resource/image/8e/67/8ee9f2fa987eccb490cfaa91c6484f67.png)

- `Kubernetes`项目的架构模型跟`Borg`非常类似, 都由`Master`和`Node`两种节点组成. 其中控制节点(Master), 由三个紧密协作的独立组件构成, 分别是负责API服务的`kuber-apiserver`, 负责调度的`kube-schedule`, 一以及负责容器编排的`kube-controller-manager`.整个集群的持久化数据, 则由`kube-apiserver`处理后保存在`Etcd`中
- 计算节点上最核心的部分`kubelet`, 主要负责同容器运行时交互, 这个交互依赖的是CRI(Container Runtime Interface)的远程调用接口, 该接口定义容器运行时的各项核心操作. 除此之外`kubelet`还通过grpc协议同`Device Plugin`的插件交互, `Device Plugin`插件是`Kubernetes`项目管理宿主机物理设备的主要组件, 也是机器学习,高性能作业支持必须关注的功能
- 容器运行时一般通过OCI这个容器运行时规范同底层的`Linux`操作系统交互, 即:把CRI请求翻译成对Linux操作系统的调用(操作Linux Namespace和Cgroups)

**Kubernetes的本质**:运行在大规模集群中的各种任务之间, 实际上存在着各种各样的关系. 这些关系的处理才是作业编排和管理系统最困难的地方

**Kubernetes设计思想**:从更宏观的角度，以统一的方式来定义任务之间的各种关系，并且为将来支持更多种类的关系留有余地

## Pod

![img](https://static001.geekbang.org/resource/image/16/06/16c095d6efb8d8c226ad9b098689f306.png)

- Pod是`Kubernetes`的原子调户单位

### Pod的实现原理

- `Pod`只是一个逻辑概念, `Kubernetes`真正处理的是还是宿主机操作系统上`Linux`容器的`Namespace`和`Cgroups`, 并且不存在一个所谓Pod的边界或者隔离环境, **Pod的实质是一组共享某些资源的容器**
- Pod的实现: 使用中间容器, 该容器叫作`Infra`容器(`infra`容器镜像: `k8s.gcr.io/pause`), 在这个`Pod`中`Infra`容器永远都是第一个被创建的容器, 其他容器通过`Join Network Namespace`的方式与`Infra`容器关联在一起

### Kubeadm 部署

- 使用配置文件部署`master`节点

```sh
# 使用v12
apiVersion: kubeadm.k8s.io/v1beta2
kind: ClusterConfiguration
controllerManager:
  extraArgs:
    horizontal-pod-autoscaler-use-rest-clients: "true"
    horizontal-pod-autoscaler-sync-period: "10s"
    node-monitor-grace-period: "10s"
apiServer:
  extraArgs:
    runtime-config: "api/all=true"
kubernetesVersion: "v1.20.4"
```


```sh
$ kubeadm init --image-repository registry.aliyuncs.com/google_containers --kubernetes-version v1.20.4 --pod-network-cidr=192.169.0.0/16 | tee /root/k8slog
# 执行完kubeadm init 部署完master节点后, 生成的token信息, 用于给master节点添加子节点
kubeadm join 172.17.0.10:6443 --token 27vba1.0wzdiqft99watbpw \
    --discovery-token-ca-cert-hash sha256:aa3fe34f31677a6a4678d351c26223cb28e5668d7aee01066880ef3c4d8af9c9
```

- 查看当前节点的状态

```shell
$ kubectl get nodes
NAME             STATUS     ROLES                  AGE   VERSION
vm-0-10-ubuntu   NotReady   control-plane,master   37m   v1.20.4
```

- 踩坑问题

  - 执行`kubeadm init`时卡住后报错 `failed to pull image k8s.gcr.io/kube-apiserver:v1.20.4`

    - 原因: ks8.gcr.io的Docker镜像国内访问不了

    - 解决办法: `kubeadm init --image-repository=registry.aliyuncs.com/google_containers --pod-network-cidr=10.244.0.0/16 --kubernetes-version=v1.20.4`

  - 执行`kubectl get nodes` 报错 `The connection to the server localhost:8080 was refused - did you specify the right host or port?`

    - 原因: `kubectl`命令需要使用`kubernetes-admin`来运行
    - 解决办法: `echo “export KUBECONFIG=/etc/kubernetes/admin.conf” >> ~/.bash_profile`

### Pod

#### 基本概念

> **Pod扮演传统部署中虚拟机的角色**, 凡是调度, 网络, 存储,以及安全相关的属性, 基本上都是Pod级别的

- **NodeSelector**: 是一个供用户将Pod与Node进行绑定的字段

- **NodeName**:

- **HostAliases**: 定义Pod的hosts文件(比如/etc/hosts)中的内容

  - 用法

  ```yaml
  apiVersion: v1
  kind: Pod
  ...
  spec:
    hostAliases:
    - ip: "10.1.2.3"
      hostnames:
      - "foo.remote"
      - "bar.remote"
  ...
  ```

  - 在这个`Pod`的`yaml`文件中设置了一组`ip`和`hostname`

  ```shell
  $cat /etc/hosts
  # Kubernetes-managed hosts file.
  127.0.0.1 localhost
  ...
  10.244.135.10 hostaliases-pod
  10.1.2.3 foo.remote
  10.1.2.3 bar.remote
  ```
  

- **ImagePullPolicy**: 镜像拉取策略

  镜像是容器定义的一部分, `ImagePullPolicy`的值默认为`always`, 这样每次创建`Pod`都需要重新拉取一次镜像. 除此之外当容器镜像类似于 `nginx`或者`nginx:latest`这样的名字时, `ImagePullPolicy`也会被认为是`always`

- **Lifecycle**: 定义`Container Lifecycle Hooks`, 即当容器状态发生变化时触发一系列'钩子'

#### Pod 的状态

- **Pending**:  这个状态意味着，`Pod` 的`YAML` 文件已经提交给了 `Kubernetes`，`API` 对象已经被创建并保存在 `Etcd` 当中。但是，这个 `Pod` 里有些容器因为某种原因而不能被顺利创建。比如，调度不成功。
- **Running**: 这个状态下，`Pod `已经调度成功，跟一个具体的节点绑定。它包含的容器都已经创建成功，并且至少有一个正在运行中。
- **Succeeded**: 这个状态意味着，`Pod` 里的所有容器都正常运行完毕，并且已经退出了。这种情况在运行一次性任务时最为常见。
- **Failed**: 这个状态下，`Pod` 里至少有一个容器以不正常的状态（非 0 的返回码）退出。这个状态的出现，意味着你得想办法 `Debug `这个容器的应用，比如查看 `Pod` 的 `Events `和日志。
- **Unknown**: 这是一个异常状态，意味着 `Pod `的状态不能持续地被 `kubelet` 汇报给 `kube-apiserver`，这很有可能是主从节点（`Master` 和 `Kubelet`）间的通信出现了问题.
- `Pod` 对象的 `Status `字段，还可以再细分出一组 `Conditions`。这些细分状态的值包括：`PodScheduled`、`Ready`、`Initialized`，以及 `Unschedulable`。它们主要用于描述造成当前 `Status` 的具体原因是什么。

#### Projected Volume

`Kubernetes`支持的`Projected Volume`类型

- `Secret`
- `ConfigMap`
- `Downward API`
- `ServiceAccountToken`

##### Secret

- 作用，是帮你把 Pod 想要访问的加密数据，存放到 Etcd 中。然后，你就可以通过在 Pod 的容器里挂载 Volume 的方式，访问到这些 Secret 里保存的信息了

- 典型的使用场景: 存放数据库的credential信息

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: test-projected-volume 
spec:
  containers:
  - name: test-secret-volume
    image: busybox
    args:
    - sleep
    - "86400"
    volumeMounts:
    - name: mysql-cred
      mountPath: "/projected-volume"
      readOnly: true
  volumes:
  - name: mysql-cred
    projected:
      sources:
      - secret:
          name: user
      - secret:
          name: pass
```

##### configMap

##### Downward API

##### ServiceAccountToken