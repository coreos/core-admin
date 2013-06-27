package main

import (
	"fmt"
)

func ExampleNewVersion() {
	*versionD = true
	*versionA = "app"
	*versionV = "ver"
	*versionT = "track"
	*versionP = "/asdf/"
	*versionM = "fixtures/update.metadata"

	args := []string{"README.md"}

	fmt.Printf("%s", newVersionRequestBody(args))

	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <version>
	//  <app id="app" version="ver" track="track"></app>
	//  <package name="README.md" size="419" path="/asdf/" sha1sum="ROSqaSsyVOcKZXgmf4BfOZ1ZqBo=" sha256sum="yXbomCsHeoku0Xvh1NPkdW+jNxyP77Y5VAsvzslYVGk=" required="false" MetadataSignatureRsa="VQ7HB6VA9hO4SZqUy6CLHF5LFZ2+qjL3cQDA5AqxsrHbWuAJuvib0uqIjQFjFle0Oo2D5GcbtSFnv6WA3tZSvnKDUF9g2x6diSveJtTqVZA2AGBD8hqOj6xr42aWxcfMsaA7QWY8sLxoALsSn5ZGukz1DpHRBTuSjDAMhUyH78+dOtReVRRDeLu4tSLjhEz0b/5+ZfR+bL5PGxTJJOyU031ybOf0/TMVrCtVKpam0LwA4h74Z++yZA+t+4l19dQGffWs+DwTNXC8FZueePB29LSxMqlNL9JLFfvXvpSGbJJI9gws5bSes6dj6sKMuE5rc/Cl5zMTZeWePLs4Xr9JXg==" MetadataSize="55961"></package>
	// </version>
}
