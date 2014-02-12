OpenLDAP
====

This is an OpenLDAP binding in the GO language. A fork of [mqu/openldap](https://github.com/mqu/openldap)


Installation :
-----

Installation is easy:

	# install openldap library and devel packages
	sudo apt-get install libldap libldap2-dev  # debian/ubuntu.
	sudo urpmi openldap-devel # fedora, RH, ...

	# install go
	go get github.com/mqu/openldap

	# verify you've got it :
	(cd $GOPATH ; go list ./...) | grep openldap

Usage
----

- Look a this [exemple](https://github.com/mqu/openldap/blob/master/_examples/test-openldap.go).
- a more complex example making  [LDAP search](https://github.com/mqu/openldap/blob/master/_examples/ldapsearch.go) that mimics ldapsearch command, printing out result on console.

Doc:
---
- run _go doc openldap_,
- will come soon, complete documentation in this [Wiki](https://github.com/mqu/openldap/wiki).
- look at [_examples/](https://github.com/mqu/openldap/blob/master/_examples/)*.go to see how to use this library.

Todo :
----

 - thread-safe test,
 - complete LDAP:GetOption() and LDAP:SetOption() method : now, they work only for integer values,
 - avoid using deprecated function (see LDAP_DEPRECATED flag and "// DEPRECATED" comments in *.go sources),
 - write some tests,
 - verify memory leaks (Valgrind),
 - support LDIF format (in, out),
 - add support for external commands (ldapadd, ldapdelete)
 - create an LDAP CLI (command line interface), like lftp, with commands like shell,
 - a nice GUI with GTK,
 - proxy, server,
 - what else ?


Links :
----

 - goc : http://code.google.com/p/go-wiki/wiki/cgo (how to bind native libraries to GO)
 - Openldap library (and server) : http://www.openldap.org/
 - Pure Go [LDAP](https://github.com/mmitton/ldap) library, with [ASN1](https://github.com/mmitton/asn1-ber) support.
