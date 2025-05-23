// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

type Oid int

func NewOid(name string) Oid {
	switch name {
	case NAME:
		return OID_NAME
	case UPDATED:
		return OID_UPDATED
	case CREATED:
		return OID_CREATED
	}
	return OID_IGNORE
}

func (me Oid) String() string {
	switch me {
	case OID_NAME:
		return NAME
	case OID_UPDATED:
		return UPDATED
	case OID_CREATED:
		return CREATED
	}
	return "Unordered"
}

func (me Oid) Sql() string {
	switch me {
	case OID_NAME:
		return "ORDER BY LOWER(Name)"
	case OID_UPDATED:
		return "ORDER BY updated DESC"
	case OID_CREATED:
		return "ORDER BY created"
	}
	return "" // unordered
}
