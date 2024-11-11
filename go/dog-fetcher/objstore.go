package main

// PART 2:
// import (
// 	"errors"
// 	"io"

// 	"github.com/bytecodealliance/wasm-tools-go/cm"
// 	"github.com/ricochet/wasmcon-na-2024/go/dog-fetcher/gen/wasi/blobstore/v0.2.0-draft/blobstore"
// 	"github.com/ricochet/wasmcon-na-2024/go/dog-fetcher/gen/wasi/blobstore/v0.2.0-draft/types"
// )

// func writeObject(reader io.ReadCloser, containerName string, filename string) error {
// 	// Get the wasi:blobstore container to store the image in
// 	container, err := ensureContainer(containerName)
// 	if err != nil {
// 		return err
// 	}

// 	// Setup an outgoing request to store
// 	outgoingValue := types.OutgoingValueNewOutgoingValue()
// 	outgoingStreamRes := outgoingValue.OutgoingValueWriteBody()
// 	if outgoingStreamRes.IsErr() {
// 		return errors.New("failed to open a new stream for storing the object")
// 	}
// 	outgoingStream := outgoingStreamRes.OK()
// 	defer outgoingStream.ResourceDrop()

// 	// Set up the write to blobstore
// 	write := container.WriteData(types.ObjectName(filename), outgoingValue)
// 	if write.IsErr() {
// 		return errors.New("failed to start writing data to blob store")
// 	}

// 	// Read 4096 bytes at a time and write it to the blobstore stream
// 	buf := make([]byte, 4096)
// 	for {
// 		written, err := reader.Read(buf)
// 		if err != nil {
// 			return errors.New("failed to flush bytes to blob store")
// 		}
// 		if written == 0 {
// 			break
// 		}
// 		writeRes := outgoingStream.BlockingWriteAndFlush(cm.ToList(buf))
// 		if writeRes.IsErr() {
// 			return errors.New("failed to flush bytes to blob store")
// 		}
// 		if written < 4096 {
// 			break
// 		}
// 	}
// 	return nil
// }

// func ensureContainer(name string) (*blobstore.Container, error) {
// 	containerName := types.ContainerName(name)
// 	exists := blobstore.ContainerExists(containerName)
// 	if ok := exists.OK(); ok != nil && *ok {
// 		containerReq := blobstore.GetContainer(containerName)
// 		if containerReq.IsErr() {
// 			return nil, errors.New("container does not exist")
// 		}
// 		return containerReq.OK(), nil
// 	}

// 	containerReq := blobstore.CreateContainer(containerName)
// 	if containerReq.IsErr() {
// 		return nil, errors.New("could not create a new container")
// 	}
// 	return containerReq.OK(), nil
// }
