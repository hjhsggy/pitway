**基础概念**
 - cgroup
 - namespace

Linux Namespace是linux提供的一种内核级别环境隔离的方法

  - 分类
  Mount namespace
  UTS namespace
  IPC namespace
  PID namespace
  Network namespace
  User namespace


**Docker项目的核心**

- 1启动Linux Namespace配置
- 2设置需要的cgroup参数
- 3切换进程的根目录(rootfs)

rootfs仅仅是操作系统所包含的文件, 配置和目录, 并不包括操作系统内核, 在linux操作
系统中这两部分是分开存放的, 操作系统只在开机启动时才会加在指定版本的内核镜像

*容器之间共享宿主机的内核*

**容器镜像**

- 联合文件系统: 功能将多个不同位置的目录联合挂载到同一个目录下

docker镜像的设计中, 引入层的概念, 也就是说用户制作镜像的每一步操作都会生成一个层

**Dockerfile**
  - https://images2017.cnblogs.com/blog/911490/201712/911490-20171208222222062-849020400.png

  - *Docker原语*
    - **FROM**: 指定基础镜像, 必须为第一个命令
      - FROM <image>|<image>:<tag>|<image>@<digest>
    - **MAINTAINER**: 维护者信息
    - **RUN**:构建镜像时执行的命令
      - RUN <command>
    - **ADD**: 添加文件到容器中,tar文件会自动解压(网络资源不会被解压)
      - ADD <file>
    - **COPY**:与ADD功能类似但不具备访问方资源和解压文件
    - **CMD**: 构建容器后调用, 容器启动时才进行调用
      - CMD ["param1", "param2"](设置了ENTRYPOINT, 则直接调用ENTRYPOINT添加参数)
	  - CMD command param1 param2 执行shell内部命令
    - **ENTRYPOINT**: 配置容器.使其可执行化. 配合CMD可省去"application",只使用参数
      - ENTRYPOINT ["executable", "param1", "param2"] 可执行文件优先
      - ENTRYPOINT command param1 param2 (shell内部命令)
    - **LABLE**: 用于为镜像添加元数据
    - **ENV**: 设置环境变量
	  - ENV key value(单个) | key=value ...(多组)
    - **EXPOSE**: 指定外界交互的端口, expose并不会让容器的端口访问到主机, 需要在docker run运行容器时通过 -p 来发布这些端口, 或者-P参数来发布expose导出的所有端口
	  - EXPOSE port port
    - **VOLUME**: 指定持久化目录
	  - VOLUME ["/path"]
    - **WORKDIR**: 工作目录, 类似于cd命令
    - **USER**: 指定运行容器时的用户名或者UID, 后续的run命令也会使用指定用户
    - **ARG**: 用于执行传递给构建运行时的变量
    - **ONBUILD**: 用于设置镜像触发器
   
   Docker 常用命令


 - *docker exec*如何进入容器?
 
   Linux Namespace 创建的隔离空间虽然看不见摸不着但是每个进程的Namespace的信息在
   宿主机上确确实实存在, 并且以一个文件的方式存在
   
   查看docker 创建的进程真实的进程号命令 
   `docker inspect --format '{{ .State.Pid }}' docker进程号`
   
   一个进程的每种Linux Namespace都在宿主机上 `/proc/[进程号]/ns`目录下对应一个
   虚拟文件并且链接到真实的Namespace文件上
   
   原理: 一个进程可以选择加入到某个进程已有的Namespace中, 从而达到进入这个进程所在的容器的目的

**Volume数据卷**

- 数据卷解决什么问题?
  - 容器中的文件如何让宿主机获取到
  - 宿主机上文件如何让容器访问到
  

