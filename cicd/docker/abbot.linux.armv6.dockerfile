ARG ARCH=armv6

FROM arhatdev/builder-go:alpine as builder

ARG ARCH=armv6
# # currently there is no pre-built armv6 cni plugins
# ARG CNI_PLUGINS_VERSION="v0.8.7"

# COPY scripts/image/download.sh /download
# RUN set -e;\
#     sh /download cni_plugins "${ARCH}" "${CNI_PLUGINS_VERSION}"

FROM arhatdev/go:debian-${ARCH}

# COPY --from=builder /opt/cni/bin /opt/cni/bin

# add required packages
RUN set -e ;\
    apt update ;\
    apt install -y iptables ;\
    rm -rf /var/lib/apt/lists/*

ENTRYPOINT [ "/abbot" ]
