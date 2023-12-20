import yaml

def split_openapi_by_tags(file_path):
    # Read the OpenAPI spec file
    with open(file_path, 'r') as file:
        openapi_spec = yaml.safe_load(file)

    # Extract the paths and components from the original spec
    paths = openapi_spec.get('paths', {})
    components = openapi_spec.get('components', {})

    # Dictionary to store the new specs divided by tags
    specs_by_tag = {}

    # Iterate over each path and its operations
    for path, operations in paths.items():
        for operation_type, operation_details in operations.items():
            # Extract the tags for each operation
            tags = operation_details.get('tags', [])
            for tag in tags:
                # Initialize spec for this tag if not already done
                if tag not in specs_by_tag:
                    specs_by_tag[tag] = {
                        'openapi': openapi_spec['openapi'],
                        'info': openapi_spec['info'],
                        'paths': {},
                        'components': components
                    }
                # Add the path and operation to the spec of the respective tag
                if path not in specs_by_tag[tag]['paths']:
                    specs_by_tag[tag]['paths'][path] = {}
                specs_by_tag[tag]['paths'][path][operation_type] = operation_details

    # Save each new spec into a separate YAML file
    output_files = []
    for tag, spec in specs_by_tag.items():
        output_file = f'specs/{tag.replace(" ", "_")}_API_Spec.yaml'
        with open(output_file, 'w') as file:
            yaml.safe_dump(spec, file)
        output_files.append(output_file)

    return output_files

# Example usage
file_path = 'API_3.1.yaml' # Replace with your file path
split_openapi_by_tags(file_path)