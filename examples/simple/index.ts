import * as testpkg from "@pulumi/testpkg";

const page = new testpkg.StaticPage("page", {
    indexContent: "<html><body><p>Hello world!</p></body></html>",
    myEnum: testpkg.Access.Public,
});

export const bucket = page.bucket;
