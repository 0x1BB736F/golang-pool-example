# Description

developed with golang 1.16

![test coverage 100%](assets/coverage.svg)

Simple pooling implementation using a `strings.Builder`.  
In this case, the pool allows you to organize the work and reuse of several strings.Builder structs.

## Advanatages

- Allows you to set the size of the pool, while protecting against negative and zero values.
- Allows to use context for timeouts.
- It easy to replace `strings.Builder` with another one (witers/readers e.t.c)

## Disadvanatages

- Too simple example.
