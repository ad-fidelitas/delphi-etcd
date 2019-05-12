# delphi-etcd

Dockerized instance to run an etcd server cluster for RPi clients to connect to.


## Running Dockerized ETCD

Use [this link](https://coreos.com/etcd/docs/latest/v2/docker_guide.html) to host a server cluster on the cloud using the etcd pre-made docker

In the build process, this Go script will run on startup of the etcd server to populate it with a key-value pair for each Node currently within the database.