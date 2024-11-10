// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package container represents the imported interface "wasi:blobstore/container@0.2.0-draft".
//
// a Container is a collection of objects
package container

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/protochron/http-password-checker-go/gen/wasi/blobstore/v0.2.0-draft/types"
	"github.com/protochron/http-password-checker-go/gen/wasi/io/v0.2.1/streams"
)

// InputStream represents the imported type alias "wasi:blobstore/container@0.2.0-draft#input-stream".
//
// See [streams.InputStream] for more information.
type InputStream = streams.InputStream

// OutputStream represents the imported type alias "wasi:blobstore/container@0.2.0-draft#output-stream".
//
// See [streams.OutputStream] for more information.
type OutputStream = streams.OutputStream

// ContainerMetadata represents the type alias "wasi:blobstore/container@0.2.0-draft#container-metadata".
//
// See [types.ContainerMetadata] for more information.
type ContainerMetadata = types.ContainerMetadata

// Error represents the type alias "wasi:blobstore/container@0.2.0-draft#error".
//
// See [types.Error] for more information.
type Error = types.Error

// IncomingValue represents the imported type alias "wasi:blobstore/container@0.2.0-draft#incoming-value".
//
// See [types.IncomingValue] for more information.
type IncomingValue = types.IncomingValue

// ObjectMetadata represents the type alias "wasi:blobstore/container@0.2.0-draft#object-metadata".
//
// See [types.ObjectMetadata] for more information.
type ObjectMetadata = types.ObjectMetadata

// ObjectName represents the type alias "wasi:blobstore/container@0.2.0-draft#object-name".
//
// See [types.ObjectName] for more information.
type ObjectName = types.ObjectName

// OutgoingValue represents the imported type alias "wasi:blobstore/container@0.2.0-draft#outgoing-value".
//
// See [types.OutgoingValue] for more information.
type OutgoingValue = types.OutgoingValue

// Container represents the imported resource "wasi:blobstore/container@0.2.0-draft#container".
//
// this defines the `container` resource
//
//	resource container
type Container cm.Resource

// ResourceDrop represents the imported resource-drop for resource "container".
//
// Drops a resource handle.
//
//go:nosplit
func (self Container) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_ContainerResourceDrop((uint32)(self0))
	return
}

// Clear represents the imported method "clear".
//
// removes all objects within the container, leaving the container empty.
//
//	clear: func() -> result<_, error>
//
//go:nosplit
func (self Container) Clear() (result cm.Result[Error, struct{}, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_ContainerClear((uint32)(self0), &result)
	return
}

// DeleteObject represents the imported method "delete-object".
//
// deletes object.
// does not return error if object did not exist.
//
//	delete-object: func(name: object-name) -> result<_, error>
//
//go:nosplit
func (self Container) DeleteObject(name ObjectName) (result cm.Result[Error, struct{}, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	name0, name1 := cm.LowerString(name)
	wasmimport_ContainerDeleteObject((uint32)(self0), (*uint8)(name0), (uint32)(name1), &result)
	return
}

// DeleteObjects represents the imported method "delete-objects".
//
// deletes multiple objects in the container
//
//	delete-objects: func(names: list<object-name>) -> result<_, error>
//
//go:nosplit
func (self Container) DeleteObjects(names cm.List[ObjectName]) (result cm.Result[Error, struct{}, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	names0, names1 := cm.LowerList(names)
	wasmimport_ContainerDeleteObjects((uint32)(self0), (*ObjectName)(names0), (uint32)(names1), &result)
	return
}

// GetData represents the imported method "get-data".
//
// retrieves an object or portion of an object, as a resource.
// Start and end offsets are inclusive.
// Once a data-blob resource has been created, the underlying bytes are held by the
// blobstore service for the lifetime
// of the data-blob resource, even if the object they came from is later deleted.
//
//	get-data: func(name: object-name, start: u64, end: u64) -> result<incoming-value,
//	error>
//
//go:nosplit
func (self Container) GetData(name ObjectName, start uint64, end uint64) (result cm.Result[string, IncomingValue, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	name0, name1 := cm.LowerString(name)
	start0 := (uint64)(start)
	end0 := (uint64)(end)
	wasmimport_ContainerGetData((uint32)(self0), (*uint8)(name0), (uint32)(name1), (uint64)(start0), (uint64)(end0), &result)
	return
}

// HasObject represents the imported method "has-object".
//
// returns true if the object exists in this container
//
//	has-object: func(name: object-name) -> result<bool, error>
//
//go:nosplit
func (self Container) HasObject(name ObjectName) (result cm.Result[string, bool, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	name0, name1 := cm.LowerString(name)
	wasmimport_ContainerHasObject((uint32)(self0), (*uint8)(name0), (uint32)(name1), &result)
	return
}

// Info represents the imported method "info".
//
// returns container metadata
//
//	info: func() -> result<container-metadata, error>
//
//go:nosplit
func (self Container) Info() (result cm.Result[ContainerMetadataShape, ContainerMetadata, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_ContainerInfo((uint32)(self0), &result)
	return
}

// ListObjects represents the imported method "list-objects".
//
// returns list of objects in the container. Order is undefined.
//
//	list-objects: func() -> result<stream-object-names, error>
//
//go:nosplit
func (self Container) ListObjects() (result cm.Result[string, StreamObjectNames, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_ContainerListObjects((uint32)(self0), &result)
	return
}

// Name represents the imported method "name".
//
// returns container name
//
//	name: func() -> result<string, error>
//
//go:nosplit
func (self Container) Name() (result cm.Result[string, string, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_ContainerName((uint32)(self0), &result)
	return
}

// ObjectInfo represents the imported method "object-info".
//
// returns metadata for the object
//
//	object-info: func(name: object-name) -> result<object-metadata, error>
//
//go:nosplit
func (self Container) ObjectInfo(name ObjectName) (result cm.Result[ObjectMetadataShape, ObjectMetadata, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	name0, name1 := cm.LowerString(name)
	wasmimport_ContainerObjectInfo((uint32)(self0), (*uint8)(name0), (uint32)(name1), &result)
	return
}

// WriteData represents the imported method "write-data".
//
// creates or replaces an object with the data blob.
//
//	write-data: func(name: object-name, data: borrow<outgoing-value>) -> result<_,
//	error>
//
//go:nosplit
func (self Container) WriteData(name ObjectName, data OutgoingValue) (result cm.Result[Error, struct{}, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	name0, name1 := cm.LowerString(name)
	data0 := cm.Reinterpret[uint32](data)
	wasmimport_ContainerWriteData((uint32)(self0), (*uint8)(name0), (uint32)(name1), (uint32)(data0), &result)
	return
}

// StreamObjectNames represents the imported resource "wasi:blobstore/container@0.2.0-draft#stream-object-names".
//
// this defines the `stream-object-names` resource which is a representation of stream<object-name>
//
//	resource stream-object-names
type StreamObjectNames cm.Resource

// ResourceDrop represents the imported resource-drop for resource "stream-object-names".
//
// Drops a resource handle.
//
//go:nosplit
func (self StreamObjectNames) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_StreamObjectNamesResourceDrop((uint32)(self0))
	return
}

// ReadStreamObjectNames represents the imported method "read-stream-object-names".
//
// reads the next number of objects from the stream
//
// This function returns the list of objects read, and a boolean indicating if the
// end of the stream was reached.
//
//	read-stream-object-names: func(len: u64) -> result<tuple<list<object-name>, bool>,
//	error>
//
//go:nosplit
func (self StreamObjectNames) ReadStreamObjectNames(len_ uint64) (result cm.Result[TupleListObjectNameBoolShape, cm.Tuple[cm.List[ObjectName], bool], Error]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_StreamObjectNamesReadStreamObjectNames((uint32)(self0), (uint64)(len0), &result)
	return
}

// SkipStreamObjectNames represents the imported method "skip-stream-object-names".
//
// skip the next number of objects in the stream
//
// This function returns the number of objects skipped, and a boolean indicating if
// the end of the stream was reached.
//
//	skip-stream-object-names: func(num: u64) -> result<tuple<u64, bool>, error>
//
//go:nosplit
func (self StreamObjectNames) SkipStreamObjectNames(num uint64) (result cm.Result[TupleU64BoolShape, cm.Tuple[uint64, bool], Error]) {
	self0 := cm.Reinterpret[uint32](self)
	num0 := (uint64)(num)
	wasmimport_StreamObjectNamesSkipStreamObjectNames((uint32)(self0), (uint64)(num0), &result)
	return
}
