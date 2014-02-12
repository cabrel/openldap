/**
 *
 * Original work Copyright (C) 2012 [Marc Quinton]
 * Modified work Copyright 2014 Robin Harper
 *
 * Use of this source code is governed by the MIT Licence :
 *  http://opensource.org/licenses/mit-license.php
 *
 * Permission is hereby granted, free of charge, to any person obtaining
 * a copy of this software and associated documentation files (the
 * "Software"), to deal in the Software without restriction, including
 * without limitation the rights to use, copy, modify, merge, publish,
 * distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to
 * the following conditions:
 *
 * The above copyright notice and this permission notice shall be
 * included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY
 * CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
 * TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
 * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 *  LDIF Spec: http://www.ietf.org/rfc/rfc2849.txt
 *
 *  Example 1: An simple LDAP file with two entries
 *
 *  version: 1
 *  dn: cn=Barbara Jensen, ou=Product Development, dc=airius, dc=com
 *  objectclass: top
 *  objectclass: person
 *  objectclass: organizationalPerson
 *  cn: Barbara Jensen
 *  cn: Barbara J Jensen
 *  cn: Babs Jensen
 *  sn: Jensen
 *  uid: bjensen
 *  telephonenumber: +1 408 555 1212
 *  description: A big sailing fan.
 *
 *  dn: cn=Bjorn Jensen, ou=Accounting, dc=airius, dc=com
 *  objectclass: top
 *  objectclass: person
 *  objectclass: organizationalPerson
 *  cn: Bjorn Jensen
 *  sn: Jensen
 *  telephonenumber: +1 408 555 1212
 *
 */

package openldap

import (
	"fmt"
)

func (self *LdapEntry) Ldif() string {
	return fmt.Sprintf("dn: %s\r\n", self.dn)
}
