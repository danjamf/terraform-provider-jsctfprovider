Example TF for JSC 

## Building

```go build -o terraform-provider-jsctf```

## Configuring

Add `~/.terraformrc` for dev overrides. Example config

```provider_installation {

  dev_overrides {
        "jsctf" = "/Users/userpath/golang/jsctfprovider"
}
```

## Running
Create a `main.tf`. You will not need to `terraform init` when using the above overrides.

``` 
terraform {
  required_providers {
    jsc = {
      source  = "jsctf"
      version = "1.0.0"
    }
  }
}

provider "jsc" {
  
  username = "email@company.com"
password = "password"
customerid = "customerid"
}

resource "jsc_oktaidp" "my_okta_config" {
  name = "abc"
  orgdomain = "abc.test.com"
  clientid = "0oaal7sr2ZeAQVEji5d6"
  
}


