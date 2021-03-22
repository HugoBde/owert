Owert is a Go Ray Tracer created with the help of Peter Shirley's book, "Ray Tracing in One Weekend".

To try it out:

```
git clone https://github.com/HugoBde/owert.git
cd owert
go run *.go
```


To configure your own scene, add new objects in `main()` to the `myObjs[]` slice.
To change the resolution of the image, modify the values of `const HEIGHT` and `const WIDTH`.
