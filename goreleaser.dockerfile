FROM scratch
ARG TARGETPLATFORM
ARG VERSION
ARG COMMIT
ARG DATE
COPY $TARGETPLATFORM/asciigraph /asciigraph
LABEL org.opencontainers.image.title="asciigraph" \
      org.opencontainers.image.name="asciigraph" \
      org.opencontainers.image.description="Go package to make lightweight line graphs ╭┈╯ in CLI" \
      org.opencontainers.image.url="https://github.com/guptarohit/asciigraph" \
      org.opencontainers.image.source="https://github.com/guptarohit/asciigraph" \
      org.opencontainers.image.version="${VERSION}" \
      org.opencontainers.image.created="${DATE}" \
      org.opencontainers.image.revision="${COMMIT}" \
      org.opencontainers.image.licenses="BSD-3-Clause"
ENTRYPOINT ["/asciigraph"]
