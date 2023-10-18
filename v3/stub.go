package main

// #include <stdint.h>
// typedef const char * const_char_p;
// typedef const unsigned char * const_unsigned_char_p;
// typedef uintptr_t datax_sdk_v3_message;
import "C"

func main() {
}

// datax_sdk_v3_initialize Initialize DataX SDK
//
//export datax_sdk_v3_initialize
func datax_sdk_v3_initialize() {
	panic("datax_sdk_v3_begin_initialize not implemented")
}

// datax_sdk_v3_next Receives an input message.
//
// The input message must be freed using the function datax_sdk_v3_message_close()
//
//export datax_sdk_v3_next
func datax_sdk_v3_next() C.datax_sdk_v3_message {
	panic("datax_sdk_v3_next not implemented")
}

// datax_sdk_v3_emit Publishes a message
//
// The message data should be encoded in msgpack format
//
//export datax_sdk_v3_emit
func datax_sdk_v3_emit(data C.const_unsigned_char_p, data_size C.int32_t, reference C.const_char_p) {
	panic("datax_sdk_v3_emit not implemented")
}

// datax_sdk_v3_message_close Frees the resources of the message
//
//export datax_sdk_v3_message_close
func datax_sdk_v3_message_close(message C.datax_sdk_v3_message) {
	panic("datax_sdk_v3_message_close not implemented")
}

// datax_sdk_v3_message_reference Obtains the reference of the message
//
//export datax_sdk_v3_message_reference
func datax_sdk_v3_message_reference(message C.datax_sdk_v3_message) C.const_char_p {
	panic("datax_sdk_v3_message_reference not implemented")
}

// datax_sdk_v3_message_stream Obtains the name of the stream that generated the message
//
//export datax_sdk_v3_message_stream
func datax_sdk_v3_message_stream(message C.datax_sdk_v3_message) C.const_char_p {
	panic("datax_sdk_v3_message_stream not implemented")
}

// datax_sdk_v3_message_data Obtains the data of the message
//
//export datax_sdk_v3_message_data
func datax_sdk_v3_message_data(message C.datax_sdk_v3_message) C.const_unsigned_char_p {
	panic("datax_sdk_v3_message_data not implemented")
}

// datax_sdk_v3_message_data_size Obtains the size of the data in the message
//
//export datax_sdk_v3_message_data_size
func datax_sdk_v3_message_data_size(message C.datax_sdk_v3_message) C.int32_t {
	panic("datax_sdk_v3_message_data_size not implemented")
}
