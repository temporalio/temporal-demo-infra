import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cfg = new pulumi.Config()
const awsCfg = new pulumi.Config('aws')

const awsProfile = awsCfg.require('profile')
const region = awsCfg.require('region')

let provider = new aws.Provider(region, {
    region: region as aws.Region,
    profile: awsProfile,
})

const applicationId = cfg.require('applicationId')

// Create an AWS resource (S3 Bucket)
const bucket = new aws.s3.Bucket(applicationId,{},{provider});

const indexContent = `<html><head>
<title>Hello S3</title><meta charset="UTF-8">
</head>
<body><p>Hello, ${applicationId}!</p><p>Made with ❤️ with <a href="https://pulumi.com">Pulumi</a></p>
</body></html>
`

// write our index.html into the site bucket
let object = new aws.s3.BucketObject("index", {
    bucket: bucket,
    content: indexContent,
    contentType: "text/html; charset=utf-8",
    key: "index.html"
}, { provider});


export const bucketId = bucket.id;
export const bucketArn = bucket.arn
export const objectId = object.id