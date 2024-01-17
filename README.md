# Pingdom Golang Client

This client is generated with [openapitools/openapi-generator](https://github.com/OpenAPITools/openapi-generator). It is based on the [Pingdom OpenAPI specification](https://docs.pingdom.com/api/).

## Usage

Check out the examples in the [test directory](./pkg/pingdom/test/openapi_test.go).
Further documentation can found at [generated documentation](./pkg/pingdom/openapi/README.md).
However that is not 100% accurate due to wrong openapi specification.

## DEPREACTED Generate the client

This client is generated out of the [Pingdom OpenAPI specification](https://docs.pingdom.com/api/).
We use [oapi-codegen](https://github.com/deepmap/oapi-codegen) to generate the client.

- Download the latest [Pingdom OpenAPI specification](https://docs.pingdom.com/api/).
- Execute python script to split the specification into multiple files per package

```bash
cd hack
pip install -r requirements.txt
python split_crd.py
```

- Generate the client
```bash
task generate-client
```
