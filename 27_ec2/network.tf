resource "aws_vpc" "prjctr" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "prjctr_vpc"
  }
}

resource "aws_internet_gateway" "prjctr" {
  vpc_id = aws_vpc.prjctr.id

  tags = {
    Name = "prjctr_igw"
  }
}

resource "aws_subnet" "prjctr_public_a" {
  vpc_id                  = aws_vpc.prjctr.id
  cidr_block              = "10.0.64.0/19"
  availability_zone       = "${var.aws_region}a"
  map_public_ip_on_launch = true

  tags = {
    "Name" = "prjctr_public-${var.aws_region}a"
  }
}

resource "aws_subnet" "prjctr_public_b" {
  vpc_id                  = aws_vpc.prjctr.id
  cidr_block              = "10.0.96.0/19"
  availability_zone       = "${var.aws_region}b"
  map_public_ip_on_launch = true

  tags = {
    "Name" = "prjctr_public-${var.aws_region}b"
  }
}


# Route table

resource "aws_route_table" "prjctr_public" {
  vpc_id = aws_vpc.prjctr.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.prjctr.id
  }

  tags = {
    Name = "prjctr_public"
  }
}

resource "aws_route_table_association" "prjctr_public_a" {
  subnet_id      = aws_subnet.prjctr_public_a.id
  route_table_id = aws_route_table.prjctr_public.id
}

resource "aws_route_table_association" "prjctr_public_b" {
  subnet_id      = aws_subnet.prjctr_public_b.id
  route_table_id = aws_route_table.prjctr_public.id
}


