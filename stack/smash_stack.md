[Smashing the stack](https://inst.eecs.berkeley.edu/~cs161/fa08/papers/stack_smashing.pdf)

We have different layers of memory:
CPU -> RAM -> SSD -> network

- cpu have registers, which are quickly accessible locations available to a computers processor. registers usually consist of a small amount of fast storage

* RAM (random access memory) is the hardware in a computing device where the OS, app programs, and data in current user are kept to be quickly reached by the device's processor.
* RAM is the main memory in a computer. RAM can be can compared to a person's short term memory and hard disk drive to a person's long term memory.
* processes occur in RAM

Processes have 3 regions: text, data, and stack

Text region is fixed by program and includes code instructions and read-only data. It corresponds to the text section of the executable file.
Any attempt to write to it will result in a segmentation violation

Data region has initialized and uninitialized data.

New data is added between the data and stack segments

Stack is last in first out (LIFO). Two of the most important operations are push (adding an element to top of stack) and pop (removing the last element at the top of the stack)

# why do we use the stack?

High level programming languages have functions/procedures.
A procedure call alters the flow of control.  
When finished performing its task, the function returns control to the statement following the call.
This is accomplished with the help of the stack.
Also, the stack is used to:

- allocate local variables used in functions
- pass parameters to functions
- return values from functions

# pointers, prologue, and epilogue

- instruction pointer (aka program counter) : is a register that holds the address of the instruction to execute next
  - cpu is hard-wired to read the instruction pointer and execute the instruction at that address.
  - after an instruction is executed by the cpu, the instruction pointer is auto-incremented to point to the next instruction in the program. loops and other branches change control flow by changing the instruction pointer
- stack pointer: where is the top of the stack in memory
- a base pointer: the beginning of the stack frame to easily reference local variables
- instruction pointer and stack pointer hold locations (addresses in memory)

* every function has a prologue and epilogue

* a prologue sets up the stack frame of the called function
* an epilogue restores the stack frame of the calling parent function

- a prologue saves the return address and updates the base pointer
- an epilogue copies the return address (saved by prologue) and gives it to the instruction pointer

# the stack region

a stack is a contiguous block of memory containing data

- a register called the stack pointer points to the top of the stack.
- the bottom of the stack is a fixed address. its size is dynamically adjusted by the kernal at run time

- the stack is made of stack frames that are pushed when calling a function and popped when returning.
- the cpu implements the instructions to push onto and pop off of the stack.
- keeps track of where to return to - a stack frame has the data necessary to recover the previous stack frame, including the value of the instruction pointer at the time of the function call
- a stack frame also has the parameters to a function and its local variables
- a stack can grow down towards lower memory addresses or up (but it normally grows down)

# what is a buffer overflow?

- buffers are memory storage regions that temporarily hold data while it is being transferred from one region to another

- a memory buffer overflow occurs when the volume of data exceeds the storage capacity of the buffer

* the program attempting to write the data to the buffer overwrites adjacent memory locations
* Stack-based buffer overflows leverage stack memory that only exists during the execution time of a function.
