[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse OS: Library for logical or abstracted `filesystem`
**URL** [multiverse-os.org](https://multiverse-os.org)

A generic filesystem library intended to be the foundation for abstracted or
logical filesystems. Implemented by using a minimal set of re-definable
low-level actions that can be combined and re-combined to acheive all the 
complexity needed for any type of filesystem. Built around completely logical,
memory file descriptors, or regular file descriptors, and loopback based 
mounting to provide functionality for building a wide variety of logical, and 
abstracted, and even convetional filesystem types. 

Designed specifically to build a wide variety of filesystems, with a simple API,
from memoryFS, GitFS, httpFS, or even existing implementations like ext4,
and others. 

Using a set of *re-definable* set of actions for each file type defined and 
filtered by `os.FileMode`: **`Read`**,  **`Create`**, **`Delete`**, and 
**`Write`**, along with a set of hooks before, an action, an after an action,
and a channel for handling events to avoid short-polling, and providing 
the tools for event-driven designs. Lastly, the types used are used repeatedly
throughout the design, enabling embedding and combining different types of 
filesystems together creating a simple foundation for complex development. 

For example, the same `os.FileMode` is used for Permissions, and categorizing
and filtering actions; and is compatible with existing POSIX filesystems.  

#### Middleware 
By implementing core logic as optional, before action hooks, or after action
hooks, we are able to take functionality that in some implementations of generic
logical/abstracted filesystem libraries, is implemented individually, as its 
own filesystem; for example an atomicFS, that may be manually merged with a
separate impelementation that is a memoryFS for example.  

The approach `filesystem` library uses is generic low-level actions, that can be  
redefined, including before, and after action hooks. This enables Atomicity is 
implemented as a middleware, so that it can be added as a hook to any existing 
filesystem. MemoryFS is implemented as an action, but since the action is
redefinable as anything, the MemoryFS component can be called from a CREATE
function that calls both MemoryFS and a GitFS filesystem CREATE functions. These
different abstractions are included in the form of the low-level actions that
are redifined so can be easily used individually or together, and with a wide
variety of hooks or event callbacks, for complex, repeatable and re-useable
code. 

This design enables greater sharing of code between projects, and this style of
design allows for greater complexit, with less code that is easier to read.


