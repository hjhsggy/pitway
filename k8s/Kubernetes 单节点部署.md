`Linux`环境下` Kubernetes`单节点部署踩坑

> 说明: 以下所有的操作都在root用户下执行, 非root用户对应命令前加sudo

- 硬件环境

```
系统: ubuntu 18.04TLS
CPU: 2核 (kubernetes需要至少两核配置)
内存: 4G  
硬盘: 50G
```

#### docker配置

- 准备工作

  - 卸载旧的docker版本

    ```shell
    $sudo apt-get remove docker docker-engine docker.io containerd runc
    ```

- 安装

  - 添加依赖

    ```shell
    $sudo apt-get install apt-transport-https ca-certificates software-properties-common
    ```

  - 使用阿里云镜像地址

    ```shell
    $sudo apt-get -y install apt-transport-https ca-certificates curl software-properties-common
    $curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
    $add-apt-repository "deb [arch=amd64] http://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
    ```

  - 更新软件源

    ```shell
    $sudo apt-get -y update
    ```

  - 安装docker(默认安装最新版, 也可以指定版本)

    ```shell
    $sudo apt-get install docker-ce #默认安装最新版
    $apt-cache madison docker-ce #查看Docker-CE版本
    $sudo apt-get install docker-ce=<VERSION> #选择上面中的版本号
    ```

  - 配置阿里云镜像加速(需要使用个人账号, 也可以使用其他加速站点)

    ```shell
    $touch /etc/docker/daemon.json
    $cat <<EOF > /etc/docker/daemon.json
    {
    "registry-mirrors": [
       "https://dockerhub.azk8s.cn",
       "https://hub-mirror.c.163.com",
       "https://ktrsh7na.mirror.aliyuncs.com"
    ]
    }
    EOF
    # 重启docker
    $systemctl daemon-reload
    $service docker restart
    $docker info #查看docker镜像源验证配置是否生效
    ```

#### Kubernetes配置

- 准备工作

  - 禁用swap (kubernetes不支持swap, 如果未关闭, 执行kubeadm init的时候会报错)

    ```shell
    $swapoff -a
    ```

  - 安装必要工具

    ```shell
    $sudo apt update && sudo apt install -y apt-transport-https curl
    ```

  - 添加阿里云源

    ```shell
    $echo "deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main" >>/etc/apt/sources.list.d/kubernetes.list
    ```

  - 添加签名

    ```shell
    $curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -
    ```

  - 更新软件源

    ```shell
    $sudo apt update
    ```

  - 安装`kubeadm`, `kubectl` `kubelet` (`node`节点仅需要安装`kubeadm`即可)

    ```shell
    $sudo apt install kubeadm kubelet kubectl
    ```

  - 修改`docker`的`driver`为`systemd`

    ```shell
    # 添加 "exec-opts": ["native.cgroupdriver=systemd"] 到 /etc/docker/daemon.json中
    $systemctl daemon-reload
    $service docker restart # 重启docker
    $docker info | grep -i cgroup  #查看docker driver信息
    $Cgroup Driver: systemd
    ```

- 安装`kubernetes`

  - 部署master节点

    - 建议修改下hostname为master `hostname -b master`

    - 执行kubeadm init

      ```shell
      kubeadm init --image-repository registry.aliyuncs.com/google_containers --kubernetes-version v1.20.4 --pod-network-cidr=192.169.0.0/16 | tee /root/k8slog
      ```

      参数说明

      --image-repository 指定镜像源, 此处为阿里云源

      --kubernetes-version 系统安装的kubernetes版本, 通过`kubectl version`可查看当前安装的版本

      --pod-network-cidr 指定pod网络地址

      k8slog 安装日志

    - 安装成功后会生成 token, node节点加入master的时候会用到, 也可以使用下面的命令重新生成

      ```shell
      kubeadm token create --print-join-command
      ```

    - 配置kubectl配置文件

      ```shell
      mkdir -p $HOME/.kube
      sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
      sudo chown $(id -u):$(id -g) $HOME/.kube/config
      ```

    - 查看master节点是否部署成功

      ```shell
      $kubectl get nodes
      NAME     STATUS      ROLES                  AGE   VERSION
      master   NotReady    control-plane,master   17h   v1.20.4
      # 可以看到节点状态NotReady,这是因为网络插件还未安装
      $kubectl get pods -n kube-system
      NAME                             READY   STATUS             RESTARTS   AGE
      coredns-6955765f44-j7lvd         0/1     CrashLoopBackOff   14         51m
      coredns-6955765f44-kmhfc         0/1     CrashLoopBackOff   14         51m
      etcd-ubuntu                      1/1     Running            0          52m
      kube-apiserver-ubuntu            1/1     Running            0          52m
      kube-controller-manager-ubuntu   1/1     Running            0          52m
      kube-proxy-qlhfs                 1/1     Running            0          51m
      kube-scheduler-ubuntu            1/1     Running            0          52m
      # 使用kubectl 命令查看节点信息
      $kubectl describe node master
      ```

    - 安装网络插件[Calico]

      ```shell
      $kubectl apply -f https://docs.projectcalico.org/v3.11/manifests/calico.yaml
      ```

    - 安装dashboard

      ```shell
      $kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-rc7/aio/deploy/recommended.yaml
      #网络不好可以手动拷贝recommended.yaml文件内容到本地, 执行kubectl apply -f 文件名
      ```

    - master节点部署pod(出于安全考虑pod不会调动到master node, master node不参与工作负载)

      - 个人测试使用允许master节点参与工作负载

        ```shell
        kubectl taint nodes --all node-role.kubernetes.io/master-
        ```

      - 禁止master节点参与工作负载

        ```shell
        kubectl taint nodes k8s node-role.kubernetes.io/master=true:NoSchedule
        ```

  - 部署node节点

- [参考链接](https://blog.csdn.net/subfate/article/details/103774087)

