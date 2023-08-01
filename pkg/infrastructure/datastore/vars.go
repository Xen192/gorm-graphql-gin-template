package datastore

import (
	"sync"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)

var Mutex = sync.Mutex{}
var ClerkClient clerk.Client
