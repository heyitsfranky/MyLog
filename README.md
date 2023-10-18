# MyLog

**MyLog** is a Go package designed to enable fast logging via Kafka.

## Installation

To use **MyLog** in your Go project, you can simply run:

```bash
go get github.com/heyitsfranky/MyLog@latest
```

## Usage

Here's a basic example of how to use **MyLog**:
```go
package main

import (
	"fmt"
	"github.com/heyitsfranky/MyLog"
)

func main() {
	configFilePath := "path/to/your/config.yaml"

	// Initialize MyLog with the configuration file
	if err := myLog.Init(configFilePath); err != nil {
		fmt.Printf("Error initializing MyLog: %s\n", err)
		return
	}
    // Create a intricate log event asynchronously
	body := map[string]interface{}{"msg": "hello world!", "id": 42}
	myLog.CreateEvent(body, "main()", 0, true)
	// (or synchronously)
	myLog.CreateEvent(body, "main()", 0, false)
    
    // Create predefined log events asynchronously
	myLog.CreateInfoEvent("This is an informational message", "MyPackage::main()")
	myLog.CreateWarningEvent("This is a warning message", "somefn()")
	myLog.CreateErrorEvent("This is an error message", "foobar()")
	myLog.CreateCriticalEvent("This is a critical error message", "MyClass.hello()")
}
```

## Configuration

To load the initial settings for event communication via kafka, a configuration file must be given that contains:

```yaml
client-origin: your-app-name,
kafka-broker: localhost:9092,
```

An up-to-date template can *always* be found under 'configs/config.yaml'.

## License

This package is distributed under the MIT License.
Feel free to contribute or report issues on GitHub.

Happy coding!