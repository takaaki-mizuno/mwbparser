# mwbparser - MWB Parser for golang

### Features

* Parse [MySQL Workbench](https://www.mysql.com/jp/products/workbench/) file and extract information about tables, columns, indexes and foreign keys.


### Examples

```go
import (
	"github.com/takaaki-mizuno/mwbparser"
)

tables, err := Parse(MWB_FILE_PATH)
```
