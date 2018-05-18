# grpc-math-service

This is a totally useless service that I created in order to refresh some
Go/Protobuf/GRPC skills that had been atropyhing.

It serves as an example of about the simpliest server and client you can have
that has more than one RPC available, and which doesn't inline the business
logic, allowing it to be ignorant of anything protobuf and generated-code
related. This makes the business logic easier to test, and potentially, able to
be more easily reused, if you dare to invoke the "r" word.
