FROM alpine:latest as base

RUN adduser -u 10000 -H -D discord-command-cleaner

FROM scratch

COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=base /etc/passwd /etc/passwd

COPY discord-command-cleaner /bin/discord-command-cleaner

USER discord-command-cleaner

ENTRYPOINT [ "/bin/discord-command-cleaner" ]
