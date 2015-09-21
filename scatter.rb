require 'ffi'
require 'json'

module Request
	extend FFI::Library

	ffi_lib File.expand_path('./main/libscatter.so', File.dirname(__FILE__))
	attach_function :scatter_request, [:string], :string
end

params = { uris: ['http://www.google.com', 'http://www.omnipasteapp.com'] }
p Request.scatter_request(params.to_json)
