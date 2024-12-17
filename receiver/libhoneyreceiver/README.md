# Libhoney Receiver
<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [development]: traces, logs   |
| Distributions | [] |
| Issues        | [![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Areceiver%2Flibhoney%20&label=open&color=orange&logo=opentelemetry)](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aopen+is%3Aissue+label%3Areceiver%2Flibhoney) [![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Areceiver%2Flibhoney%20&label=closed&color=blue&logo=opentelemetry)](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aclosed+is%3Aissue+label%3Areceiver%2Flibhoney) |
| [Code Owners](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/CONTRIBUTING.md#becoming-a-code-owner)    | [@TylerHelmuth](https://www.github.com/TylerHelmuth) |

[development]: https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/component-stability.md#development
<!-- end autogenerated section -->

### The purpose and use-cases of the new component

The Libhoney receiver will accept data for either Trace or Logs signals that are emitted from applications that were
instrumented using [Libhoney](https://docs.honeycomb.io/send-data/logs/structured/libhoney/) libraries.

## Configuration

The configuration has 2 parts, One is the HTTP receiver configuration and the rest is about mapping attributes from the
freeform libhoney format into the more structured OpenTelemetry objects.

### Example configuration for the component

The following settings are required:

- `http`
  - `endpoint` must set an endpoint. Defaults to `127.0.0.1:8080`
- `resources`: if the `service.name` field is different, map it here.
- `scopes`: to get the `library.name` and `library.version` set in the scope section, set them here.
- `attributes`: if the other trace-related data have different keys, map them here, defaults are otlp-like field names.

The following setting is required for refinery traffic since:

- `auth_api`: should be set to `https://api.honeycomb.io` or a proxy that forwards to that host.
  Some libhoney software checks `/1/auth` to get environment names so it needs to be passed through.


```yaml
  libhoney:
    http:
      endpoint: 0.0.0.0:8088
      traces_url_paths:
        - "/1/events"
        - "/1/batch"
      include_metadata: true
    auth_api: https://api.honeycomb.io
    resources:
      service_name: service_name
    scopes:
      library_name: library.name
      library_version: library.version
    attributes:
      trace_id: trace_id
      parent_id: parent_id
      span_id: span_id
      name: name
      error: error
      spankind: span.kind
      durationFields:
        - duration_ms
```

### Telemetry data types supported

It will subscribe to the Traces and Logs signals but accept traffic destined for either pipeline using one http receiver
component. Libhoney doesnot differentiate between the two so the receiver will identify which pipeline to deliver the 
spans or log records to.

No support for metrics since they'd look just like logs.