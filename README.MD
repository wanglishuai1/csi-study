# 原理层：

### 我们需要干的是：
>IdentityService、NodeService、ControllerService   ---- 称为 自定义组件

### 官方提供的有：
> node-driver-registrar、csi-attacher、csi-provisioner  --- 这些都是外部组件 是以sidecar形式 和我们的 自定义组件  配合

# 详解原理层二开：
#### Identity Service ： 
>暴露插件的名称  和能力

#### Controller Service 
>(NFS)   mount –t xxxxxx--  NodeService
用于实现创建/删除 volume等，不需要在特定宿主机完成的操作、譬如和云商的API进行交互 以及attach操作等

#### NodeService
>真正的执行 mount、unmount 。 所以它必须在每台机器上都存在  (daemonset)

### livenessprobe
>监控CSI 驱动程序的运行状况

### node-driver-registrar
>获取驱动程序信息（ NodeGetInfo ），并 在节点上的 kubelet 中进行注册

### csi-provisioner
>举例：监听 PVC 创建，调用 CreateVolume 方法创建对存储对象 ，进一步创建 PV资源。当监听PVC被删，则调用 DeleteVolume 方法删除底层存储，进一步删除PV

### csi-attacher
监听 Kubernetes VolumeAttachment 对象并触发 CSI  Controller[Publish|Unpublish]Volume 操作。

    

# yaml代码大概使用方式    
1、执行rbac.yaml 给驱动添加权限

    1、创建ServiceAccount，创建一个账户
    2、创建ClusterRole，创建一个角色
    3、创建ClusterRoleBinding，将角色绑定到账户上

2、执行driver.yaml 运行驱动

    1、创建驱动的Deployment，里面有5个容器，
    2、其中一个是驱动容器，其他的是csi的容器，
    3、csi的容器是csi的插件，用来和k8s交互的，
    4、驱动容器是驱动的实现，用来和存储交互的，驱动容器和csi的容器通过unix socket通信，csi的容器和k8s通过grpc通信
    


3、执行testcsi.yaml 

    1、创建存储类，存储类的驱动器就是上面的driver
    2、使用存储类创建pvc
    3、使用Deployment创建pod，pod使用pvc



