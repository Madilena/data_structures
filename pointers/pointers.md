These are my notes from [Golang Tutorial 3 by Junmin Lee](https://www.youtube.com/watch?v=sTFJtxJXkaY)

A goroutine is an independent path of execution
Each goroutine is given a stack of memory - a very lightweight frame
every time a goroutine makes a function call, a part of the stack is allocated and we call that a frame

Each frame can only work within itself - this guarantees immutabilitiy

when we pass in a value to a function, we are copying that value from one frame to another
this is safe but not efficient because we are copying the arguments each time we make a function call

the stack is self cleaning - when a function call is finished, it just discards the frame of that function and everything inside it. when another function is called, this space will be used by another frame.

the heap needs a garbage collector - giving too much work to garbage collector can affect performance.
