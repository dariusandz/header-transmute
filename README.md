# Header transmutation plugin for traefik (DEPRECATED AND MOVED TO https://github.com/inventi/traefik-header-transmute)

Plugin replaces header in request with value loaded from environment variable.

## Plugin Configuration

```yaml
FromHeader: 'Header to be replaced' 
ToHeader: 'Header to replace FromHeader with'
MappingEnvKey: 'Environment variable key in which mapping configuration is stored'
```

Mapping configuration syntax is a simple key value pairs multiline string where each line's mapping pair is split by `:`, e.g.
```text
key1:value1
key2:value2
key3:value3
```

## Example

Provided with plugin configuration
```yaml
FromHeader: X-Old-Header
ToHeader: X-New-Header
MappingEnvKey: HeaderMapping
```
and header value mapping configuration stored in environment variable with key `HeaderMapping`
```text
value:newValue
otherValue:newOtherValue
```

will transmute request
```text
curl <url> -H "X-Old-Header: value"
```
deleting `X-Old-Header` header from request and appending `X-New-Header` header with value `newValue` to request.

If header values mapping does not exist or there is no such mapping for header's value, 
plugin will always transmute header from `FromHeader` to `ToHeader`.

### Improvements to be made
* Support multiple headers to be transmuted
* Load configuration from k8s secret (when https://github.com/traefik/traefik/pull/9103 is released)
