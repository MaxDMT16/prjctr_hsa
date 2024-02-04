# Serverless

This repo folder contains the code for lamda func that is triggered on upload to S3 object to `jpeg_images` folder with `.jpeg` extension. The function will convert it to other formats and save to appropriate folders.

|format|folder|
|---|---|
|png|png_images|
|bmp|bmp_images|
|gif|gif_images|

## Infrastructure

### Prerequisites
It's expected that you have `aws` cli installed and configured with the necessary permissions.
To configure the aws cli, run the following command:

```bash
aws configure
```

### Setup

To setup the infrastructure, move to `terraform` folder and run the following command:

```bash
terraform init
terraform apply
```
