# ================ S3 ================

resource "aws_s3_bucket" "prjctr_bucket" {
  bucket = var.bucket_name

}

resource "aws_s3_object" "folder_jpeg_images" {
  bucket = aws_s3_bucket.prjctr_bucket.id
  acl    = "private"
  key    = "jpeg_images/"
}

resource "aws_s3_object" "folder_bmp_images" {
  bucket = aws_s3_bucket.prjctr_bucket.id
  acl    = "private"
  key    = "bmp_images/"
}

resource "aws_s3_object" "folder_gif_images" {
  bucket = aws_s3_bucket.prjctr_bucket.id
  acl    = "private"
  key    = "gif_images/"
}

resource "aws_s3_object" "folder_png_images" {
  bucket = aws_s3_bucket.prjctr_bucket.id
  acl    = "private"
  key    = "png_images/"
}


# ================ Lambda ================

data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "iam_for_lambda" {
  name               = "iam_for_lambda"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

locals {
  function_name = "image-converter"
  src_path      = "${path.module}/../cmd/main.go"

  binary_name  = local.function_name
  binary_path  = "${path.module}/tf_generated/${local.binary_name}"
  archive_path = "${path.module}/tf_generated/${local.function_name}.zip"
}

output "binary_path" {
  value = local.binary_path
}

resource "null_resource" "function_binary" {
  triggers = {
    always_run = timestamp()
  }

  provisioner "local-exec" {
    command = "GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOFLAGS=-trimpath go build -mod=readonly -ldflags='-s -w' -o ${local.binary_path} ${local.src_path}"
  }
}

data "archive_file" "image_converter_binary" {
  depends_on = [null_resource.function_binary]

  type        = "zip"
  source_file = local.binary_path
  output_path = local.archive_path
}

resource "aws_lambda_function" "img_converter_lambda" {
  filename      = data.archive_file.image_converter_binary.output_path
  function_name = local.function_name
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = local.binary_name
  memory_size   = 128
  timeout       = 120

  source_code_hash = data.archive_file.image_converter_binary.output_base64sha256

  runtime = "go1.x"

  environment {
    variables = {
      REGION = var.aws_region
      BUCKET = aws_s3_bucket.prjctr_bucket.id
    }
  }
}

# ================ Trigger Lambda on S3 upload ================

resource "aws_s3_bucket_notification" "my-trigger" {
  bucket = aws_s3_bucket.prjctr_bucket.id

  lambda_function {
    lambda_function_arn = aws_lambda_function.img_converter_lambda.arn
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = "jpeg_images/"
    filter_suffix       = ".jpeg"
  }
}

resource "aws_lambda_permission" "prjctr_bucket_invoke_lambda" {
  statement_id  = "AllowS3Invoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.img_converter_lambda.arn
  principal     = "s3.amazonaws.com"
  source_arn    = aws_s3_bucket.prjctr_bucket.arn
}

# ================ Policy for lambda to write logs ================

data "aws_iam_policy" "AWSLambdaBasicExecutionRole" {
  arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "lambda_logging_policy_attachment" {
  role       = aws_iam_role.iam_for_lambda.id
  policy_arn = data.aws_iam_policy.AWSLambdaBasicExecutionRole.arn
}

data "aws_iam_policy" "AmazonS3FullAccess" {
  arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
}

resource "aws_iam_role_policy_attachment" "lambda_s3_policy_attachment" {
  role       = aws_iam_role.iam_for_lambda.id
  policy_arn = data.aws_iam_policy.AmazonS3FullAccess.arn
}
