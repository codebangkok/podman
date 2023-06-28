# podman
```
brew install podman
podman machine ls
podman machine init
podman machine start
podman info
```
podman directory
```
~/.ssh/podman-machine-default
~/.ssh/podman-machine-default.pub
~/.config/containers
```

Registries
```
podman machine ssh
vi /etc/containers/registries.conf
```

Image
```
podman image search (podman search)
podman image pull (podman pull)
podman image ls (podman images)
podman image rm (podman rmi)
podman image build (podman build)
```

Container
```
podman container ls (podman ps)
podman container run (podman run)
podman container rm (podman rm)
```

Search
```
podman search mariadb
```
