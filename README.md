## Ghostinthepdf

This is a small tool that helps to embed a PostScript file into a PDF in a way that GhostScript will run the PostScript code during the PDF processing.

The goal of the embedding is to run exploits against GhostScript. The output file is a semi-correct PDF that should bypass most signature checks. Thus, if you have an exploit that bypasses `-dSAFER`, you can "embed" it into a PDF using this tool and upload the resulting file to the target.

The resulting file will execute the PostScript code only if processed using GhostScript. Other tools/frameworks will not run the code.

### Installation

```
$ go version
go version go1.17 linux/amd64
$ go install github.com/neex/ghostinthepdf@latest
```

### Usage

```bash
ghostinthepdf input.ps output.pdf
```

### Examples

One can use the file [print_version.ps](print_version.ps) to detect if a target uses GhostScript for PDF processing. The PDF obtained by `ghostinthepdf print_version.ps output.pdf` will print the GhostScript version into the preview if processed by GhostScript. The preview will be blank or contain the words "GhostScript not detected" if the PDF is processed by another tool.

### License

This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or distribute this software, either in source code form or as a compiled binary, for any purpose, commercial or non-commercial, and by any means.

In jurisdictions that recognize copyright laws, the author or authors of this software dedicate any and all copyright interest in the software to the public domain. We make this dedication for the benefit of the public at large and to the detriment of our heirs and successors. We intend this dedication to be an overt act of relinquishment in perpetuity of all present and future rights to this software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to https://unlicense.org
