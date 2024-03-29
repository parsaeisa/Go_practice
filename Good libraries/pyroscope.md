# pyroscope

Example usage:
```go
profiler.Start(profiler.Config{
    ApplicationName: os.Getenv("PYROSCOPE_APP_NAME"),
    ServerAddress:   os.Getenv("PYROSCOPE_SERVER_ADDRESS"),
    ProfileTypes: []profiler.ProfileType{
        profiler.ProfileCPU,
        profiler.ProfileAllocObjects,
        profiler.ProfileAllocSpace,
        profiler.ProfileInuseObjects,
        profiler.ProfileInuseSpace,
    },
})
```