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

# Download the specified GitHub release, unpack it and save
# it to /bin where it gets copied from in subsequent stages.
RUN curl -LO https://github.com/dominikbraun/observe/releases/download/${VERSION}/observe-linux-amd64.tar.gz && \
    tar -xzvf observe-linux-amd64.tar.gz -C /bin && \
    rm -f observe-linux-amd64.tar.gz

# Start the final stage.
FROM alpine:3.11.5 AS final

LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.name="observe"
LABEL org.label-schema.description="Observe a website and get an e-mail if something changes."
LABEL org.label-schema.url="https://github.com/dominikbraun/observe"
LABEL org.label-schema.vcs-url="https://github.com/dominikbraun/observe"
LABEL org.label-schema.version=${VERSION}
LABEL org.label-schema.docker.cmd="docker container run -v $(pwd):/settings dominikbraun/observe"

COPY --from=download ["/bin/observe", "/bin/observe"]

# The settings directory should contain the observe settings,
# i. e. a mounted settings.yml file.
RUN mkdir /settings
WORKDIR /settings

# ENTRYPOINT gets set to the observe binary so that only observe
# commands are valid. The acual command is left up to the user.
ENTRYPOINT ["/bin/observe"]
CMD [""]