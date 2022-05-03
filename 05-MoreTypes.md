# Slices

* three values: pointer to storage array, length, capacity

![](SlicesViz.png)

* Slice tricks: https://github.com/golang/go/wiki/SliceTricks

## Relevant builtin functions

The builtin `append` can be used to append one element or a slice to an existing
slice. Builtin `copy` to duplicate a slice.

In summary:

* make, new -- [SliceMakeVsNew](x/SlicesExtra/SliceMakeVsNew/main.go)
* len
* cap
* append
* copy

> The zero value of a slice is `nil`. The `len` and `cap` functions will both return 0 for a nil slice.

Refs:

* [builtin](https://golang.org/pkg/builtin/)

## [A] From Slice to Array and back

> Q: How do you go from a slice to an array and back?

* [SliceStorage](x/SlicesExtra/SliceStorage/main.go)

## [A] Examples

> Q: What is happening?

```go
a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
a = append(a[:2], a[7:]...)
```

* [SliceCut](x/SlicesExtra/SliceCut/main.go)

### More examples

We can batch slices (also example for the reference nature of slices)

* [SliceBatching](x/SlicesExtra/SliceBatching/main.go)

A safe version of cut (deleting values).

* [SliceCutSafe](x/SlicesExtra/SliceCutSafe/main.go)

Similar idea, with deletions.

* [SliceDeleteSafe](x/SlicesExtra/SliceDeleteSafe/main.go)

### Safe slice ops in summary

1. reuse existing slice and copy values
2. set remaining values to nil
3. shrink slice

### Extend slices

Expanding a slice. The "double-appender".

* [SliceExpand](x/SlicesExtra/SliceExpand/main.go)

Extending a slice. Use "..." to extend a slice.

* [SliceExtend](x/SlicesExtra/SliceExtend/main.go)

Custom filters.

* [SliceFilter](x/SlicesExtra/SliceFilter/main.go)
* [SliceFilterNoAlloc](x/SlicesExtra/SliceFilterNoAlloc/main.go)


## Are the generic version of these functions?

> No, not yet. There might be in the future.

## [A] Slice Gotchas

> Q: What is happening in the following snippet.

* [SliceGotchaAppend](x/SlicesExtra/SliceGotchaAppend/main.go)
* [SliceGotchaAppendBytes](x/SlicesExtra/SliceGotchaAppendBytes/main.go)
