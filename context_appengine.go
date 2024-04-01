// Copyright 2017 Manu Martinez-Almeida. All rights reserved.

//go:build appengine

package gin

func init() {
	defaultPlatform = PlatformGoogleAppEngine
}
