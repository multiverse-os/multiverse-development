##
##  Multiverse OS: Database [Development Notes]
===========================================================
In addition to being development notes, Multiverse OS will
provide a three layer database that will be easily accessible
to the user, and tie in to alot of the foundational Multiverse OS
functionality.


===========================================================
## Multiverse OS: SkylabDB (K/V DB, Graph DB, Cache DB)
-----------------------------------------------------------
A three layer database that provides rich functionality 
that can be accessed at whatever level of complexity is
needed.
#### Layer 0

#### Layer 1

#### Layer 2

===========================================================
## Database Utilities
-----------------------------------------------------------
#### Compression
[pgzip](https://github.com/klauspost/pgzip)
Go parallel gzip (de)compression

[readahead](https://github.com/klauspost/readahead)
Asynchronous read-ahead for Go readers

[dedup](https://github.com/klauspost/dedup)
Streaming Deduplication Package for Go

[reedsolomon](https://github.com/klauspost/reedsolomon)
Reed-Solomon Erasure Coding in Go

[compress](https://github.com/klauspost/compress)
Optimized compression packages



===========================================================
## Data Types / Data Containers 
-----------------------------------------------------------
[vecstring](https://github.com/hillbig/vecstring)
vecstring is a Go library for a vector representation of strings.  Strings are concatenated in a byte array, and each offset and its length are encoded by unary coding, and stored using rsdic package.

[fixvec](https://github.com/hillbig/fixvec)
fixvec is a Go library for a vector representation of values using fixed bits.
fixvec provides a vector representation of value using fixed bits. Conceptually, fixvec represents a vector V[0...num), and each value V[i] can represent in [0...2^(blen)). The total working space is num * blen bits (+ some small overhead)

[succinct trie](https://github.com/hillbig/wx)
**This switched to be a suffix instead of prefix, could be really good at matching trades in an orderbook**

[wavelet tree](https://github.com/hillbig/waveletTree)
waveletTree is a Go package for myriad array operations using wavelet trees.

[sparse dense array](https://github.com/hillbig/sdarray)
In numerical analysis and computer science, a sparse matrix or sparse array is a matrix in which most of the elements are zero. By contrast, if most of the elements are nonzero, then the matrix is considered dense. The number of zero-valued elements divided by the total number of elements (e.g., m × n for an m × n matrix) is called the sparsity of the matrix (which is equal to 1 minus the density of the matrix).

[rdic](https://github.com/hillbig/rsdic)
rsdic is a Go package for rank/select dictionary supporting rank/select operations efficiently, and space efficient for both sparse and dense bit arrays. (Such data structures are also called as fully indexable dictionary in CS literatures (FID)).

===========================================================
## General Database Options [ Development Notes ] 
-----------------------------------------------------------

