## Go

These settings apply only when `--go` is specified on the command line.

``` yaml $(go)
go:
  license-header: MICROSOFT_MIT_NO_VERSION
  namespace: sdk
  go-version: 1.19
  clear-output-folder: false
  module-name: sdk
  module: github.com/microsoft/commercial-marketplace-offer-deploy/sdk
  output-folder: $(go-sdk-folder) 
```

### Tag: preview-2023-03-01 and go

These settings apply only when `--tag=preview-2023-03-01 --go` is specified on the command line.
Please also specify `--go-sdk-folder=../sdk`.

``` yaml $(tag) == 'preview-2023-03-01' && $(go)
output-folder: $(go-sdk-folder)
```
