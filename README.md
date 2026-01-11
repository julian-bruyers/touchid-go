<!--  README badges  -->
<p>
    <a href="#platform-support"><img src="https://img.shields.io/badge/macOS-333333?logo=apple&logoColor=F0F0F0" align="right"></a>
    <a href="#installation"><img src="https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&labelColor=2D3748&color=2D3748" align="right" height="20" hspace="5"></a>
    <a href="https://github.com/julian-bruyers/touchid-go/releases"><img src="https://img.shields.io/github/v/release/julian-bruyers/touchid-go?label=Latest&labelColor=2D3748&color=003087" height="20"></a>
    <a href="https://github.com/julian-bruyers/touchid-go/blob/main/LICENSE"><img src="https://img.shields.io/github/license/julian-bruyers/touchid-go?&label=License&logo=opensourceinitiative&logoColor=ffffff&labelColor=2D3748&color=2D3748" height="20" hspace="5"></a>
    <a href="https://goreportcard.com/report/github.com/julian-bruyers/touchid-go"><img src="https://goreportcard.com/badge/github.com/julian-bruyers/touchid-go" height="20"></a>
</p>

# touchid-go: macOS Touch ID authentication for Go

A **lightweight**, **zero external dependencies** and **simple to use** Go library providing Touch ID and Face ID biometric authentication (fingerprint, facial recognition) for macOS.

## Overview
**touchid-go** bridges Go with the native macOS API by using standard CGo calls. This let's you use Touch ID or Face ID directly without any external dependencies or complex CGo calls, just one simple Go function.


### Features
- Supports Intel and Apple Silicon Mac's
- Compatible with all macOS versions since macOS 10.12.1 (Sierra)
- No external dependencies
- Extremely easy to use API

---

> [!IMPORTANT]
> **App Store & Info.plist Requirements**
>
> If you plan to distribute your application via the **Mac App Store** or use a **hardened runtime**, you **must** add the `NSFaceIDUsageDescription` key to your application's `Info.plist` file.
>
> Even if you only intend to use Touch ID, Apple requires this key because the library links against the `LocalAuthentication` framework, which supports both Face ID and Touch ID.
>
> **Not including this key might result in App Store rejection or cause your application to crash immediately when `Authenticate()` is called.**
>
> **Example `Info.plist` entry:**
> ```xml
> <key>NSFaceIDUsageDescription</key>
> <string>This app uses Touch ID / Face ID to verify your identity.</string>
> ```

---

## Installation
Add touchid-go to your Go project:

```bash
go get github.com/julian-bruyers/touchid-go@latest
```

> [!NOTE]
> **C Compiler Required**
> 
> This library relies on CGo, so you must have the macOS development tools installed to build your project.
> Simply run `xcode-select --install` in your terminal to install them.

## Usage Example
```go
package main

import (
    "fmt"

    "github.com/julian-bruyers/touchid-go"
)

func main() {
	// Check if Touch ID is available
	if !touchid.Available() {
		log.Fatalln("Touch ID is unavailable")
	}
	
	// Authenticate the user
    if isAuthenticated, _ := touchid.Authenticate("Verify your identity for touchid-go test"); isAuthenticated {
        fmt.Println("Authentication successful!")
    } else {
        fmt.Println("Authentication failed!")
    }
}
```

**With error handling:**
```go
package main

import (
    "fmt"
    "log"

    "github.com/julian-bruyers/touchid-go"
)

func main() {
	// Check if Touch ID is available
	if !touchid.Available() {
		log.Fatalln("Touch ID is unavailable")
	}
	
	// Authenticate the user
	isAuthenticated, err := touchid.Authenticate("Verify your identity for touchid-go test")

	// Handle the auth error
	if err != nil {
		log.Fatal(err)
	}

	if isAuthenticated {
		fmt.Println("Authentication successful!")
	} else {
		fmt.Println("Authentication failed!")
	}
}
```

## API / Usage
> [!TIP]
> **touchid.Avialable() (bool)**
> Checks if Touch ID is available on the current system.
> 
> _Parameters:_
> - None
> 
> _Return:_
> - `bool:` `true` if Touch ID is available, `false` otherwise

---

> [!TIP]
> **touchid.Authenticate(promptMsg string) (bool, error)**
> Prompts the user to authenticate using Touch ID.
> 
> _Parameters:_
> - `promptMsg`: The message displayed to the user during authentication
> 
> _Returns:_
> - `bool`: `true` if authentication was successful, `false` otherwise
> - `error`: An error if something went wrong, `nil` on success
>
> | Error | Description |
> |-------|-------------|
> | `ErrOsNotSupported` | Called on non-macOS systems |
> | `ErrArchNotSupported` | Unsupported CPU architecture (only AMD64 and ARM64 supported) |
> | `ErrNotAvailable` | Touch ID not configured or available on the system |
> | `ErrUserCanceled` | User canceled the authentication prompt |
> | `ErrInternal` | Internal macOS API error |


## System Requirements
- macOS 10.12.1 (Sierra)
- Go 1.25 or later (earlier versions likely work aswell)
- Apple C compiler (standard with `xcode-select --install`)
- AMD64 or ARM64 processor
- MacBook with Touch ID sensor or Touch ID compatible keyboard

## Development and Building
1. **Go 1.25+**
   - Download: https://golang.org/dl/

2. **Clone the Repository:**
   ```bash
   git clone https://github.com/julian-bruyers/touchid-go.git
   cd touchid-go
   ```
   
3. **Build Example Application:**
   ```bash
   go run examples/main.go
   ```

## Project Structure
```
touchid-go/ 
├── examples/ 
│   ├── go.mod              # Module definition for examples 
│   └── main.go             # Example application 
├── native/ 
│   └── touchid.c           # Native C implementation 
├── auth.stub.go            # Stub for non-macOS systems 
├── auth_darwin.go          # macOS implementation 
├── errors.go               # Error type definitions 
├── go.mod                  # Go module definition 
├── LICENSE                 # MIT License 
└── README.md               # This file
```

## License
This project is licensed under the MIT License. See [LICENSE](LICENSE) file for details.

Copyright (c) 2026 Julian Bruyers
