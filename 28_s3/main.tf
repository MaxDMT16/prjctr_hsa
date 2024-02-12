resource "aws_s3_bucket" "main" {
  bucket = "prjctr-aws-s3-main-with-logs-n-object-lock"

  force_destroy = true

  tags = {
    "Name" = "main"
  }
}

resource "aws_s3_bucket_versioning" "main" {
  bucket = aws_s3_bucket.main.id

  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_object_lock_configuration" "main" {
  depends_on = [ aws_s3_bucket_versioning.main ]

  bucket = aws_s3_bucket.main.id

  rule {
    default_retention {
      mode = "GOVERNANCE"   # COMPLIANCE, GOVERNANCE. With COMPLIANCE mode you can't change or delete the object till the retention period is over. With GOVERNANCE mode only users with s3:BypassGovernanceRetention permission can change or delete the object.
      days = 5
    }
  }
}

resource "aws_s3_bucket_ownership_controls" "main" {
  bucket = aws_s3_bucket.main.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_acl" "main" {
  depends_on = [ aws_s3_bucket_ownership_controls.main ]

  bucket = aws_s3_bucket.main.id
  acl    = "private"
}

resource "aws_s3_bucket" "log_bucket" {
  bucket = "prjctr-aws-s3-logs-for-main"

  force_destroy = true
}

resource "aws_s3_bucket_ownership_controls" "log_bucket" {
  bucket = aws_s3_bucket.log_bucket.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_acl" "log_bucket_acl" {
  depends_on = [ aws_s3_bucket_ownership_controls.log_bucket ]

  bucket = aws_s3_bucket.log_bucket.id
  acl    = "log-delivery-write"
}

resource "aws_s3_bucket_logging" "main" {
  bucket = aws_s3_bucket.main.id

  target_bucket = aws_s3_bucket.log_bucket.id
  target_prefix = "log/"
}

