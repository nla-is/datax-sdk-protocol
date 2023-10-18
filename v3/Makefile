all:
	docker build -t datax-sdk $(CURDIR)
	docker run -i --rm -v $(CURDIR):/mnt datax-sdk
	cp preamble.h datax-sdk.h
	awk '/End of boilerplate cgo prologue/{p++;if(p==1){next}}p' libdatax-sdk.h | \
		sed -e 's|const_char_p|const char \*|g' -e 's|const_unsigned_char_p|const unsigned char \*|g' >> datax-sdk.h