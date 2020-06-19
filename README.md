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

#### Redefining Base Actions & Middleware 
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

#### Essential Functionality Included
Every implemenetation of filesystem require access to specific functionality, 
but often this functionality is repeated, with different specific strategies. 

*Many of these common patterns are included as interdependent sub-packages as to 
not bloat the code footprint unncessarily, but available with solid 
implementation that that works seamlessly with the `filesystem` library.*

  * **bloom filter** after populating, a value can be checked if it exists in
    the collection. This can be important for very large data sets, to avoid
    unncessary searches, by checking the bloom filter before each search. Can
    even be enabled after a threshold size is hit to ensure its only being used
    in the situations its most valuable. 

  * **consistent hashing** allow allocation into different pools evenly and
    quick lookups, this can be used in a variety of ways; like for every folder
    in the root directory. Each subsequent directory could theorically be
    appended indicating the depth of the object, and allow quick lookups without
    requiring slow string comparisons.  
    
    (Multiverse OS design requirements require all string comparisons to be
    avoided, unless doing so adds unnecessary / non-preformant complexity.) 

  * **fast ID hashing** for filename, inode, and ID creation and comparisons in 
    a re-assignable function so the developer can choose their hashing function 
    of choice, but defaults to xxhash3 currently. 

  * **merkle tree** for easily building a checksum system that can confirm if
    pieces, files or any other divisible pieces that make up the filesystem 
    hiearchy are valid. 

  * **chunking** with a variety of default strategies, not just different sizes,
    but different ways of chunking-- not just even chunks. And as with any other
    critical component of the `filesystem` library, can the function can be
    re-assigned to naturally customize functionality beyond defaults. 

  * **codecs** functionality to easily add-in developer defined compression, 
    data conversion that naturally ties into the before, and after hooks. 

    Included in the codecs is encryption of data, and optionally metadata. By
    default the metadata is segregated completely from the data, so that files
    can be represented none or parts of files (for example, the beginnings of
    files for the purpose of streaming, but starting playing the media
    immediately). 

  * **moustable** filesystems via loopback and possibly in the future a fuse
    implementation. 

  * **disks: [ block / image ] + [ lvm / partition ]** optional packaging, that
    can not just interact with filesystems inside partitions, but also create
    and edit partitions, images, and working with filesystems across multiple 
    block devices. 

  * **plan9 / nfs / smb** shares for sharing data with local network, and
    combined with Multiverse OS networking tools, these shares can be made  
    easily accessible to friends, family and peers. 

#### Implemented Logical/Abstracted Filesystems 
Several filesystems are implemeneted, originally to confirm this design model
for filesystem development was reliable across wide variety of types, with a
small code footprint through itnerdependent sub-packages for various
functionality that when tied together builds more than just toy implementations
of abstracted filesystems. 

*Not all implemented yet, but this provides the roadmap of desired filesystems,
and lays out the purpose of this library.*

  * **ext{2,3,4}** the conventional filesystem of Linux, implemented to
    establish a baseline that could be benchmarked against other implementations
    and ensure all the basic functionaity expected from a filesystem is
    included. 

  * **tar** a common filesystem abstraction that is both generally useful but
    also helps establish a baseline, that can also be compared against many
    other generic filesystem libraries. 

  * **memory** a memory filesystem to replace the current Multiverse OS memory
    fs that does not currently satisfy initial design requirements because it
    lacks the ability to mount. 

  * **git** a git backed filesystem, that importantly with `filesystem` library
    can be easily combined with other filesystems; for example, memory+git can
    easily be combined for a memory read, and git write filesystem, for
    versioned persistence and quick memory reads, and combinations of two or
    more is as natural as a single storage backend with Multiverse OS's
    `filesystem` library.

  * **http** a simple http backed filesystem, that unlike *many* other
    implementations is not limited to read-only access; but supports writes
    through POST or GET requests (the chunked file can be uploaded via
    host/file-id/part/X/{data...(up to the server defined url limit)}
    simplifying concurrent uploads, and simulteanous uploads from multiple
    hosts.
  
  * **ssh** filesystem for easily attaching rented hosting, or other remote
    servers for simple access for offsite backups, publishing, or related tasks. 

  * **redis** filesystem for abstracting data stored in redis as a filesystem,
    blob, or as data within files, and heirachical structure representing
    buckets and objects. 

  * **social media stegography** backend that uses encrypted data stored in 
    profile images, comments, private messages, and many other things social 
    media companies refuse to let you throw away to backup important data like 
    cryptography keys, and other small but important data. Provides a small
    amount of storage by using crediential to several social media platforms, to
    duplicate critical data by downloading adding stegogrpahy data inline into
    images and other social media content, and re-uploading providing reliable,
    secure and off-site storage of critical data for free, and accross platorms
    unlikely to disapear. 

  * **torrent** a chunked, merkle-tree verified torrent implementation that will
    serve to demonstrate how poorly torrent compares to Multiverse's blueshift
    protocol intended to replace torrents.

  * **blueshift** a intelligently chunked, merkle-tree verified, multi-pathing,
    udp (with tcp features, like xdp or kcp) multi-casted (with tcp fallback), 
    compression supporting, streaming p2p protocol. Supporting assymetric
    encryption based, multi-signature supporting, mutability (variety of 
    algorithims supported including gpg), versioning based on diff'd binary data
    that applies updates to existing blocks, and additional features required by
    Multiverse OS key components. Intended originally to replace torrents, due
    to massive design flaws in the original torrent protocol, but resulted in a
    networking protocol that became the foundation of many Multiverse OS core
    components. 


#### Contributions
Multiverse OS projects are developed openly and transparently as possible, all
pull requests are welcomed. Design guidelines, suggestions, and further
information can be obtained in repositories within this organization, on the
Multiverse OS website, or reaching out to the active development team. 

As the project grows more community organization, discusison, and planning tools
will be made available. If you are interested in learning more or volunteering,
send an email to volunteer@multiverse-os.org or look at our projects page. 

  
