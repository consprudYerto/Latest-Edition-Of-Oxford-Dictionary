// Code generated by command: go run gen_amd64_compress_asm.go -out ../compress/blocks_amd64.s -stubs ../compress/blocks_amd64.go -pkg compress. DO NOT EDIT.

//go:build !purego

package compress

// blocksSSE2 performs BLAKE-224 and BLAKE-256 block compression
// using SSE2 extensions.  See [Blocks] in blocksisa_amd64.go for
// parameter details.
//
//go:noescape
func blocksSSE2(state *State, msg []byte, counter uint64)

// blocksSSE41 performs BLAKE-224 and BLAKE-256 block compression
// using SSE41 extensions.  See [Blocks] in blocksisa_amd64.go for
// parameter details.  The scratch parameter is not used.
//
//go:noescape
func blocksSSE41(state *State, msg []byte, counter uint64)