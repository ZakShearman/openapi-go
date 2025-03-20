# Fork specific information

When using this fork, add the following to your go.mod:
```
replace github.com/mworzala/openapi-go => github.com/ZakShearman/openapi-go VERSION_HERE
```
e.g.
```
replace github.com/mworzala/openapi-go => github.com/ZakShearman/openapi-go v0.0.0-20250320093544-f83875d717af
```

# openapi-go
An opinionated openapi server code generator for go

todo format me + add more

Below are the additional spec properties added
* `info.x-base-path` - path for the server excl. version (default: `/{specName}`, example: `/abc` -> `/v1/abc`)
* `paths.{path}.{method}.responses.{code}.x-type` - set to `empty` to use that code as the empty response. 
  That code may not have a response body.

Object properties are treated as required unless the individual field has `required: false` or 
the object root has `required: false`.
