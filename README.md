# Header transmutation plugin for traefik

This plugin replaces header in request with value loaded from environment variable configuration

## Example

Provided with plugin configuration
```yaml
FromHeader: X-Old-Header
ToHeader: X-New-Header
MappingEnvKey: HeaderMapping
```
and header value mapping configuration in environment variable named `HeaderMapping`
```text
value:newValue
otherValue:newOtherValue
```

Will transmute request
```text
curl -X GET <url> -H "X-Old-Header: value"
```
deleting `X-Old-Header` header from request and appending `X-New-Header` header with value `newValue` to request.

If value of configured header `FromHeader` is not mapped in environment variable `HeaderMapping`, 
plugin will not modify request.

### Note
If multiple header values exist, plugin will ignore them and lose these values retaining only newly mapped value.

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

### Improvements to be made
* Support multiple headers to be transmuted
* Retain multiple unmapped header values
* Load configuration straight from k8s secret (when https://github.com/traefik/traefik/pull/9103 is released)