FROM alpine

EXPOSE 8081

COPY ./Gitkins_linux_amd64 /
COPY ./Gitkins-config.yaml /
ENTRYPOINT ["/Gitkins_linux_amd64"]
