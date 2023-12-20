# Pingdom Golang Client

This client is generated out of the [Pingdom OpenAPI specification](https://docs.pingdom.com/api/).
We use [oapi-codegen](https://github.com/deepmap/oapi-codegen) to generate the client.

## Generate the client

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
