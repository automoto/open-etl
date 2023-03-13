FROM jeffail/benthos
COPY benthos-config.yaml ./benthos.yaml
ENTRYPOINT ["/benthos", "--chilled", "streams", "/benthos.yaml"]
