

Where the CNI configurations are stored
/etc/cni/net.d

List the network namespaces
ip netns list

ctr c create docker.io/bmcrae/counter-app:latest app


# gVisor start
ctr i pull docker.io/bmcrae/counter-app:latest
ctr c create --runtime io.containerd.runsc.v1 docker.io/bmcrae/counter-app:latest app
ctr t start -d app




ctr c create --runtime runsc docker.io/bmcrae/counter-app:latest app
