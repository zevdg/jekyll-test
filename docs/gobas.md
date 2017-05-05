---
---
## gobas

BAS code generator for go

### Synopsis



gobas is the bas_codegen.pl equivalent for the go programming language.
It generates schema, client, and service packages from a bas schema xsd.
The only command meant to be run manaually is

	gobas init <xsd-file>

After that, `go generate` shoud be used to generate and update the code.

### SEE ALSO
* [gobas client](gobas_client)	 - Generate a client library package
* [gobas init](gobas_init)	 - Initialize a gobas project
* [gobas schema](gobas_schema)	 - Generate a schema library package
* [gobas service](gobas_service)	 - Generate a service main package

###### Auto generated by spf13/cobra on 10-Apr-2017