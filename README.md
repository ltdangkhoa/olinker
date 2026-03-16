# oLinker Gateway

A modern Go-based hardware integration replacement tool for legacy Windows PMS vendor adapters.

## Features
- **Embedded Web UI** for configuration and API testing on port `:9000`
- **Dynamic Configuration:** Configure vendor-specific parameters (Ports, Hotel IDs, Mappings) directly in the browser.
- **Legacy PMS API Endpoints** compatible with endpoints like `POST /dlock/write_card`, `POST /adel/read_card`
- **Extensible `LockVendor` interface** for integrating Door Lock DLLs
- **Mock Vendor** included out-of-the-box (`/mock/write_card`) to test external PMS API boundaries
- **Single-worker job queue** to ensure vendor DLL thread-safety
- No command-line window, auto-launches local config page
- Portable executable with zero runtime dependencies (Windows 7+)

## Expected User Experience
1. Copy the `oLinker` folder to a Windows PC.
2. Ensure the proper `assets/dll/{vendor}` folder contains the `.dll` files.
3. Double-click `oLinker.exe`.
4. The default browser will automatically open the `oLinker Gateway Configuration` UI.
5. Configure your vendor and DLL settings.
6. Local PMS software can now perform encode/read API calls.

## API Gateway Endpoints
To connect to the linker from a PMS system, the API port defaults to `9000`. You can test endpoints via `curl` or Postman.
We support vendor parameters inside the URLs for dynamic pathings. Supported vendors include: `orbita`, `betech`, `adel`, `mock`.

* **Encode Card:** `POST http://localhost:9000/{vendor}/write_card`
  * Body: `{"room_name": "101", "BeginTime": "...", "EndTime": "..."}`
* **Read Card:** `POST http://localhost:9000/{vendor}/read_card`
* **Cancel Card:** `POST http://localhost:9000/{vendor}/cancel_card`
* **Extend Card:** `POST http://localhost:9000/{vendor}/extend_card`

## Windows Portable Build Instructions

To build the executable for Windows (supporting Windows 7 and above as a 32-bit graphical binary without a console window), run the following command from the project root directory:

```bash
GOOS=windows GOARCH=386 go build -ldflags="-H windowsgui" -o oLinker.exe ./cmd/olinker
```

*Note: You can cross-compile this from a macOS or Linux machine seamlessly using Go.*

## Adding a New Vendor SDK

The system is designed to dynamically load vendor SDK DLL wrappers safely.
To add the new "MyVendor" SDK:

1. **Add DLL assets**
   Place your DLLs into `assets/dll/myvendor/myvendor.dll`.

2. **Implement `LockVendor` interface**
   Create a new file `internal/vendors/myvendor_windows.go`:
   
   ```go
   //go:build windows
   package vendors
   
   import (
       "olinker/internal/core"
       "olinker/internal/platform"
   )

   type MyVendor struct {
       dll *platform.DLLLoader
   }
   
   func NewMyVendor() *MyVendor { return &MyVendor{} }
   
   func (v *MyVendor) Init(config core.VendorConfig) error {
       var err error
       v.dll, err = platform.NewDLLLoader(config.DLLPath)
       return err
   }
   
   func (v *MyVendor) EncodeCard(req core.EncodeRequest) (core.EncodeResult, error) {
       // get procedure and call
       // proc, _ := v.dll.GetProc("WriteCard")
       // proc.Call(...)
       return core.EncodeResult{}, nil
   }
   // Implement CancelCard, ExtendCard, ReadCard...
   ```

3. **Update `vendor_loader.go` Factory**
   Edit `internal/vendors/vendor_loader.go` to support the new vendor string:
   
   ```go
   case "myvendor":
       v = NewMyVendor()
   ```

4. **Update the Web UI**
   Add your new vendor to the `<select id="vendor-select">` dropdown inside `web/index.html`.
