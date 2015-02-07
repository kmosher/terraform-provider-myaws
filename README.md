Monkey-patches the aws provider so that AMI is a computed resource

Install by including the following in your ~/.terraformrc

    providers {
        aws = "terraform-provider-myaws"
    }
