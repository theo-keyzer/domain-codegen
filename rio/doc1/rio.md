# Rio Format Specification (v1.0)

**Rio** — a lightweight, human-writable configuration and code-generation language for describing hardware, models, and domains in a single file.

- **File extension**: `.rio`
- **MIME type**: `text/rio`
- **Pronunciation**: /ˈriː.oʊ/ (like "Rio de Janeiro")
- **Design goals**: readable · editable · auto-parses JSON where possible · safe multi-line code blocks · domain-agnostic

## Basic Structure

A Rio file is a sequence of typed blocks:

```rio
Domain {
  name = tsu
  desc = "Thermodynamic Sampling Unit domain"
}

Project {
  project = MaxCutPBit
  domain = tsu
}
```

Each block has:
- A type (`Domain`, `Project`, `Tensor`, etc.)
- A unique name (quoted if it contains spaces)
- A body in `{ … }` containing key-value fields

## Field Syntax

| Example                              | Result Type              | Notes |
|--------------------------------------|--------------------------|-------|
| `count = 1024`                       | `int`                    | auto-detected |
| `temp = 300.0`                       | `float64`                | |
| `active = true`                      | `bool`                   | |
| `name = "MaxCutPBit"`                | `string`                 | |
| `tags = [gpu, cuda, ampere]`         | `[]any`                  | parsed as JSON array |
| `meta = {a=1, b=true}`               | `map[string]any`         | parsed as JSON object |

The parser attempts JSON unmarshaling for every value. If it succeeds → native Go type; otherwise → raw string.

## Multi-line Values

### Structured multi-line (JSON-first)

Use `""" … """` when you want structured data that may span many lines:

```rio
metrics = """
{
  "cut_value": 2847,
  "energy": -2847.0,
  "steps": 8500
}
"""
```

→ automatically parsed into `map[string]any`  
→ fallback to raw string if invalid JSON

### Verbatim multi-line (code / templates)

Use `~~~ … ~~~` when you need exact text preservation (CUDA kernels, Python, shaders, etc.):

```rio
code = ~~~
__global__ void qubo_gibbs(...) {
    printf("He said \"\"\"triple quotes are fine here\"\"\"");
}
~~~
```

→ always stored as raw string  
→ can contain `"""`, `'''`, `~~~`, anything  
→ indentation and newlines are preserved exactly

**Closing delimiter rules** (both types):
- Can be indented arbitrarily
- Must be alone on its line (after trimming whitespace)
- No content after the closing delimiter on the same line

## Comments

```rio
// This is a comment
desc = "hello"  // trailing comments are stripped
```

`//` comments are ignored unless inside a `~~~` block.

## Example File (complete)

```rio
// maxcut_pbit.rio
Domain {
  name = tsu
  desc = "Thermodynamic Sampling Unit"
}

Project {
  project = MaxCutPBit
  domain = tsu
  desc = "Max-Cut on P-bit array"
}

Hardware {
  hardware = tsu_extropic_v1
  pbit_count = 1024
  operating_temp_k = 300
}

Model {
  model = maxcut_qubo_model
  num_variables = 100
}

Tensor {
  tensor = qubo_matrix
  parent = maxcut_qubo_model
  shape = [100, 100]
  dtype = float32
}

Op {
  op = qubo_gibbs_update
  code = ~~~
__global__ void qubo_gibbs(float* Q, int8_t* state) {
    // real CUDA code here
    printf("\"\"\"allowed\"\"\"");
}
~~~
}

Checkpoint {
  checkpoint_id = best_20251122
  metrics = """
{
  "cut_value": 2847,
  "energy": -2847.0
}
"""
}
```


**Rio — one file, all domains. Let it flow.**
