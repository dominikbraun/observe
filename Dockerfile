#
# observe Dockerfile (Light)
#
FROM alpine:3.11.5 AS download

# The VERSION build argument specifies the observe release
# version to be downloaded from GitHub.
ARG VERSION

RUN apk add --no-cache \
    curl \
    tar

RUN curl -LO https://github.com/dominikbraun/observe/releases/download/${VERSION}/observe-linux-amd64.tar.gz && \
    tar -xzvf observe-linux-amd64.tar.gz -C /bin && \
    rm -f observe-linux-amd64.tar.gz

FROM alpine:3.11.5 AS final

LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.name="observe"
LABEL org.label-schema.description="Observe a website and get an e-mail if something changes."
LABEL org.label-schema.url="https://github.com/dominikbraun/observe"
LABEL org.label-schema.vcs-url="https://github.com/dominikbraun/observe"
LABEL org.label-schema.version=${VERSION}
LABEL org.label-schema.docker.cmd="docker container run -v $(pwd):/settings dominikbraun/observe"

COPY --from=download ["/bin/observe", "/bin/observe"]

RUN mkdir /settings
WORKDIR /settings

ENTRYPOINT ["/bin/observe"]
CMD [""]