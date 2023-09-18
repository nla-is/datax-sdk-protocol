//
// Created by Giuseppe Coviello on 9/18/23.
//

#pragma once

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

extern void datax_sdk_v3_begin_initialize();
extern void datax_sdk_v3_end_initialize();
extern void datax_sdk_v3_next();
extern const char *datax_sdk_v3_get_message_reference();
extern const char *datax_sdk_v3_get_message_stream();
extern const unsigned char *datax_sdk_v3_get_message_data();
extern int32_t datax_sdk_v3_get_message_data_size();
extern void datax_sdk_v3_set_message_reference(const char *reference);
extern void datax_sdk_v3_set_message_data(const unsigned char *data, int32_t data_size);
extern void datax_sdk_v3_emit();

#ifdef __cplusplus
}
#endif
