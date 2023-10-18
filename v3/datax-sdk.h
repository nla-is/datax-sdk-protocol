//
// Created by Giuseppe Coviello on 9/18/23.
//

#pragma once

#include <stdint.h>

typedef uintptr_t datax_sdk_v3_message;

#ifdef __cplusplus
extern "C" {
#endif


// datax_sdk_v3_initialize Initialize DataX SDK
//
extern void datax_sdk_v3_initialize();

// datax_sdk_v3_next Receives an input message.
//
// The input message must be freed using the function datax_sdk_v3_message_close()
//
extern datax_sdk_v3_message datax_sdk_v3_next();

// datax_sdk_v3_emit Publishes a message
//
// The message data should be encoded in msgpack format
//
extern void datax_sdk_v3_emit(const unsigned char * data, int32_t data_size, const char * reference);

// datax_sdk_v3_message_close Frees the resources of the message
//
extern void datax_sdk_v3_message_close(datax_sdk_v3_message message);

// datax_sdk_v3_message_reference Obtains the reference of the message
//
extern const char * datax_sdk_v3_message_reference(datax_sdk_v3_message message);

// datax_sdk_v3_message_stream Obtains the name of the stream that generated the message
//
extern const char * datax_sdk_v3_message_stream(datax_sdk_v3_message message);

// datax_sdk_v3_message_data Obtains the data of the message
//
extern const unsigned char * datax_sdk_v3_message_data(datax_sdk_v3_message message);

// datax_sdk_v3_message_data_size Obtains the size of the data in the message
//
extern int32_t datax_sdk_v3_message_data_size(datax_sdk_v3_message message);

#ifdef __cplusplus
}
#endif
