[ "$0" != "$BASH_SOURCE" ] || { echo "Error: script must be sourced, not run"; exit 1; }


if [ ! -e kubebuilder_2.3.1_darwin_amd64 ]
then
    curl -L https://github.com/kubernetes-sigs/kubebuilder/releases/download/v2.3.1/kubebuilder_2.3.1_darwin_amd64.tar.gz | tar xz
fi

export KUBEBUILDER_ASSETS=${PWD}/kubebuilder_2.3.1_darwin_amd64/bin


if [ ! -e kustomize/bin/kustomize ]
then
    mkdir -p kustomize/bin
    curl -L https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/v3.8.8/kustomize_v3.8.8_darwin_amd64.tar.gz | tar xz -C kustomize/bin
fi

# Edit kubebuilder's generated Makefile

export KUSTOMIZE=${PWD}/kustomize/kustomize


# docker run --restart=always -p 0.0.0.0:5000:5000 --name registry registry:2
#docker start registry
DOCKER_TAG=docker.io/mtinside/kubebuilder-book-cronjob:latest
