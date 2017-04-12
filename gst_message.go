package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

func MessageNewEos(o *GstObj) *Message {
	return (*Message)(C.gst_message_new_eos(o.g()))
}
