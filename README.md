# sqldot

sqldot converts SQL to DOT file. The file is intended to be used to render diagram using `graphviz` and may help illustrating complexity of CTEs inside a query.

## Usage

Generate DOT file in specified path:

```shell
sqldot --output=path/to/output/file.dot
```

## Development

This tool uses LLM instead of regex to parse SQL. The LLM module is under development and will return mock response.
