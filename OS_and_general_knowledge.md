# OS and General Knowledge

## Terminal Usage

Went through this:

https://launchschool.com/books/command_line/read/introduction

## How OSs work in general

Major components:

- file system
- scheduler
- device driver

Three Key Elements:

1. **Abstractions** (process, thread, file, socket, memory)
2. **Mechanisms** (create, schedule, open, write, allocate)
3. **Policies** (Least Recently Used (LRU), Earliest deadline first (EDF))

Three common OS Types

1. **Monolithic** - entire OS working in kernel space and is alone in supervisor mode
2. **Modular** - some parts of the system core are located in independent files called modules that can be added to the system at runtime
3. **Micro** - kernel is broken down into separate processes, known as servers

## Process Management

A process is basically a program in execution.

A process can be divided into four sections

1. **stack** - contains temp data like method/function params, return address, local vars
2. **heap** - dynamically allocated memory to a process during its run time
3. **text** - includes current activity contents
4. **data** - contains global and static vars

## Threads and Concurrency

A thread is a flow of execution through the process code

Implemented in two ways:

1. **User Level Threads**
   - Disadvantages
     - in a typical OS, most system calls are blocking
     - multi-threaded application cant take advantage of multiprocessing
2. **Kernel Level Threads**
   - Disadvantages
     - generally slower to create and manage than user threads
     - transfer of control from on thread to another within the same process requites a mode switch to the kernel

## Memory Management

Memory management is the functionality of an operating system which handles or manages primary memory

keeps track of:

- every memory location
- how much memory is to be allocated to process
- decides which process will get memory at what time
- tracks whenever memory gets freed up or unallocated, and updates the status

Three types of addresses used in a program before and after memory is allocated:

- **Symbolic address** - addresses use in source code (var names, constants, instruction labels)
- **Relative address** - at time of compilation, compiler converts symbolic addresses into relative addresses
- **Physical address** - generates these when program is loaded into main memory

## Inter-process Communication (IPCS)

Mechanism which allows processes to communicate with each other and sync their actions

Two types of processes

- **Independent** - not affected by execution of other processes
- **Cooperating** - can be affected by other executing processes

Can communicate with each other in two ways

- **Shared Memory**
- **Message Parsing**

## (I)nput/(O)utput Management

Manage I/O devices, such as, mouse, keyboard, touch pad, disk drives, display adapters, etc.

An I/O system takes in a request (from a device) and sends it to the physical device

Three approaches to communicate with the CPU and Device

- **Special Instruction I/O**
- **Memory-mapped I/O**
- **Direct memory access (DMA)**

## POSIX Basics

Skimmed this page: https://www.howtogeek.com/435903/what-are-stdin-stdout-and-stderr-on-linux/

## Basic Networking Concepts

Skimmed through this: https://www.ece.uvic.ca/~itraore/elec567-13/notes/dist-03-4.pdf

## References

- https://launchschool.com/books/command_line/read/introduction
- http://markburgess.org/os/os.pdf
- https://medium.com/cracking-the-data-science-interview/how-operating-systems-work-10-concepts-you-should-know-as-a-developer-8d63bb38331f
- https://www.howtogeek.com/435903/what-are-stdin-stdout-and-stderr-on-linux/
- https://www.ece.uvic.ca/~itraore/elec567-13/notes/dist-03-4.pdf