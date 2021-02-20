
# Overview

This writer levels the flow rate.

# Using leveling

## Step 1: Install leveling

```shell
 $ go get github.com/takeshinoda/leveling
```

## Create leveling writer

If you want to write the byte slice each by 10KB every 10 milliseconds,
you can create leveling writer object such as the follow:

```go
    writer := leveling.NewWriter(someWriter, 10 * time.Millisecond, 10 * 1024)
```

## Question
Q. How about the reader?

A. If you want, please tell me.

## Author

takeshinoda (Takeshi Shinoda)

## License

MIT license
