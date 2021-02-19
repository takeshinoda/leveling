
# Overview

This writer levels the flow rate per second.

# Using leveling

## Step 1: Install leveling

```shell
 $ go get github.com/takeshinoda/leveling
```

## Create leveling writer

If you want to write the byte slice each by 10KB every 10 milliseconds,
you can create leveling writer object such as the follow:

```go
    writer := leveling.New(someWriter, 10 * time.Millisecond, 10 * 1024)
```

As the another way, if you want to send the 10KB of byte slice 10 times per seconds, 
you can create leveling writer object such as the follow:

```go
    writer := leveling.NewTimesPerSecond(someWriter, 10, 10 * 1024)
````

## Author

takeshinoda (Taekshi Shinoda)

## License

MIT license
