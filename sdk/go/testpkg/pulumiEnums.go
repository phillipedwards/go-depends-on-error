// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package testpkg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Access string

const (
	AccessPublic  = Access("public")
	AccessPrivate = Access("private")
)

func (Access) ElementType() reflect.Type {
	return reflect.TypeOf((*Access)(nil)).Elem()
}

func (e Access) ToAccessOutput() AccessOutput {
	return pulumi.ToOutput(e).(AccessOutput)
}

func (e Access) ToAccessOutputWithContext(ctx context.Context) AccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessOutput)
}

func (e Access) ToAccessPtrOutput() AccessPtrOutput {
	return e.ToAccessPtrOutputWithContext(context.Background())
}

func (e Access) ToAccessPtrOutputWithContext(ctx context.Context) AccessPtrOutput {
	return Access(e).ToAccessOutputWithContext(ctx).ToAccessPtrOutputWithContext(ctx)
}

func (e Access) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Access) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Access) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Access) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessOutput struct{ *pulumi.OutputState }

func (AccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Access)(nil)).Elem()
}

func (o AccessOutput) ToAccessOutput() AccessOutput {
	return o
}

func (o AccessOutput) ToAccessOutputWithContext(ctx context.Context) AccessOutput {
	return o
}

func (o AccessOutput) ToAccessPtrOutput() AccessPtrOutput {
	return o.ToAccessPtrOutputWithContext(context.Background())
}

func (o AccessOutput) ToAccessPtrOutputWithContext(ctx context.Context) AccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Access) *Access {
		return &v
	}).(AccessPtrOutput)
}

func (o AccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Access) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Access) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessPtrOutput struct{ *pulumi.OutputState }

func (AccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Access)(nil)).Elem()
}

func (o AccessPtrOutput) ToAccessPtrOutput() AccessPtrOutput {
	return o
}

func (o AccessPtrOutput) ToAccessPtrOutputWithContext(ctx context.Context) AccessPtrOutput {
	return o
}

func (o AccessPtrOutput) Elem() AccessOutput {
	return o.ApplyT(func(v *Access) Access {
		if v != nil {
			return *v
		}
		var ret Access
		return ret
	}).(AccessOutput)
}

func (o AccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Access) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AccessInput is an input type that accepts AccessArgs and AccessOutput values.
// You can construct a concrete instance of `AccessInput` via:
//
//          AccessArgs{...}
type AccessInput interface {
	pulumi.Input

	ToAccessOutput() AccessOutput
	ToAccessOutputWithContext(context.Context) AccessOutput
}

var accessPtrType = reflect.TypeOf((**Access)(nil)).Elem()

type AccessPtrInput interface {
	pulumi.Input

	ToAccessPtrOutput() AccessPtrOutput
	ToAccessPtrOutputWithContext(context.Context) AccessPtrOutput
}

type accessPtr string

func AccessPtr(v string) AccessPtrInput {
	return (*accessPtr)(&v)
}

func (*accessPtr) ElementType() reflect.Type {
	return accessPtrType
}

func (in *accessPtr) ToAccessPtrOutput() AccessPtrOutput {
	return pulumi.ToOutput(in).(AccessPtrOutput)
}

func (in *accessPtr) ToAccessPtrOutputWithContext(ctx context.Context) AccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessInput)(nil)).Elem(), Access("public"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPtrInput)(nil)).Elem(), Access("public"))
	pulumi.RegisterOutputType(AccessOutput{})
	pulumi.RegisterOutputType(AccessPtrOutput{})
}