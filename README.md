# Reading Files From Tests

In the Go function: `os.Open()` relative paths are read in relation to the
current working directory and not in relation to the working directory from
where the program was executed.

Usually these two paths are one and the same, but when running tests the
current working directory is set to the directory that contains the test file,
which can cause unexpected results.

One way for working around this issue is by using the embed package, which has the
drawback of increasing the size of your binary (usually not by much if you are not using very big files in your tests).

This example repository illustrates how you can do this. 
