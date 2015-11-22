docker-storm-cluster
====================

This was created to get an Apache Storm cluster migrated over to Kubernetes

Kubernetes and Zookeeper
------------------------

This was tested with a local docker install v1.1.1. See
[thegillis/docker-zookeeper-cluster](https://github.com/thegillis/docker-zookeeper-cluster)
to first setup Kubernetes and then Zookeeper locally.

Apache Storm
------------

This currently uses environment variables for server configurations:

* STORM_CMD - The storm command to run. Currently can be "nimbus", "supervisor", or "ui"
* CONFIGURE_ZOOKEEPER - Set this to signal that we want to automatically set up the servers in the storm.yaml file. If this is set, the following variables are checked.
* ZK_SERVER_1_SERVICE_HOST - Zookeeper Server 1
* ZK_SERVER_2_SERVICE_HOST - Zookeeper Server 2
* ZK_SERVER_${N}_SERVICE_HOST - Will keep searching for Zookeeper servers
* APACHE_STORM_NIMBUS_SERVICE_HOST - Nimbus server IP

To run this in Kubernetes:

```
kubectl create -f kubernetes-nimbus-service.yaml
kubectl create -f kubernetes-ui-service.yaml
```

Create the replication controllers:

```
kubectl create -f kubernetes-nimbus-rc.yaml
kubectl create -f kubernetes-supervisor-rc.yaml
kubectl create -f kubernetes-ui-rc.yaml
```


