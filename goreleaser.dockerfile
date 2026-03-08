FROM scratch
ARG TARGETPLATFORM
COPY $TARGETPLATFORM/asciigraph /asciigraph
ENTRYPOINT ["/asciigraph"]
