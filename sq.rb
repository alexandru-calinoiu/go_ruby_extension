require 'ffi'

module Sq
	extend FFI::Library

	ffi_lib File.expand_path('./main/libsq.so', File.dirname(__FILE__))
	attach_function :sq, [:int], :int
end

p Sq.sq(2)