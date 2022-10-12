package main

import (
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi-testpkg/sdk/go/testpkg"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// attempt to recreate the pulumi.DependsOn caused error by adding a randomId to MLC as a dependent
		rand, err := random.NewRandomId(ctx, "rand", &random.RandomIdArgs{ByteLength: pulumi.Int(4)})
		if err != nil {
			return nil
		}

		_, err = testpkg.NewStaticPage(ctx, "test", &testpkg.StaticPageArgs{
			IndexContent: pulumi.String("<body>hi</body>"),
		}, pulumi.DependsOn([]pulumi.Resource{rand}))

		return nil
	})
}
