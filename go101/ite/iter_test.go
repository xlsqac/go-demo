package main

import (
    "slices"
    "strconv"
    "testing"
)

func BenchmarkSliceFunctions(b *testing.B) {
    sizes := []int{10, 100, 1000, 10000}
    for _, size := range sizes {
        s := make([]int, size)
        for i := range s {
            s[i] = i
        }

        b.Run("ForLoop-"+strconv.Itoa(size), func(b *testing.B) {
            for range b.N {
                for i, v := range s {
                    _, _ = i, v
                }
            }
        })

        b.Run("All-"+strconv.Itoa(size), func(b *testing.B) {
            for range b.N {
                for i, v := range slices.All(s) {
                    _, _ = i, v
                }
            }
        })

        b.Run("BackwardForLoop-"+strconv.Itoa(size), func(b *testing.B) {
            for range b.N {
                for i := len(s) - 1; i >= 0; i-- {
                    _ = s[i]
                }
            }
        })

        b.Run("Backward-"+strconv.Itoa(size), func(b *testing.B) {
            for range b.N {
                for i, v := range slices.Backward(s) {
                    _, _ = i, v
                }
            }
        })

        b.Run("ValuesForLoop-"+strconv.Itoa(size), func(b *testing.B) {
            for range b.N {
                for _, v := range s {
                    _ = v
                }
            }
        })

        b.Run("Values-"+strconv.Itoa(size), func(b *testing.B) {
            for range b.N {
                for v := range slices.Values(s) {
                    _ = v
                }
            }
        })

        b.Run("AppendForLoop-"+strconv.Itoa(size), func(b *testing.B) {
            for range b.N {
                result := []int{}
                for _, v := range s {
                    result = append(result, v)
                }
                _ = result
            }
        })

        b.Run("AppendSeq-"+strconv.Itoa(size), func(b *testing.B) {
            for range b.N {
                _ = slices.AppendSeq([]int{}, slices.Values(s))
            }
        })

        b.Run("CollectForLoop-"+strconv.Itoa(size), func(b *testing.B) {
            for range b.N {
                result := make([]int, len(s))
                copy(result, s)
                _ = result
            }
        })

        b.Run("Collect-"+strconv.Itoa(size), func(b *testing.B) {
            for range b.N {
                _ = slices.Collect(slices.Values(s))
            }
        })

        b.Run("SortForLoop-"+strconv.Itoa(size), func(b *testing.B) {
            for range b.N {
                result := make([]int, len(s))
                copy(result, s)
                slices.Sort(result)
                _ = result
            }
        })

        b.Run("Sorted-"+strconv.Itoa(size), func(b *testing.B) {
            for range b.N {
                _ = slices.Sorted(slices.Values(s))
            }
        })

        b.Run("ChunkForLoop-"+strconv.Itoa(size), func(b *testing.B) {
            chunkSize := 10
            for range b.N {
                for i := 0; i < len(s); i += chunkSize {
                    end := i + chunkSize
                    if end > len(s) {
                        end = len(s)
                    }
                    _ = s[i:end]
                }
            }
        })

        b.Run("Chunk-"+strconv.Itoa(size), func(b *testing.B) {
            chunkSize := 10
            for range b.N {
                for chunk := range slices.Chunk(s, chunkSize) {
                    _ = chunk
                }
            }
        })
    }
}
