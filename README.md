# gRPCNull

Is an empty gRPC server.

Its main use is to provide a dummy service that can replace real services.

## Usage
```
grpcnull [ADDRESS] [FLAGS]
```

Example:
```
grpcnull :4001
```

## Flags
Can be used to exit with success code if the address is already in use
```
-omit-address-in-use-error true
```


