Example TF for JSC

## Building

`go build -o terraform-provider-jsctf`

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

resource "jsc_uemc" "my_uemc_config" {
  domain = "https://terraform.jamfcloud.com/"
  clientid = "1b752ccb-eaee-4250-a202-a5d1d091053c"
  clientsecret = "bvaKmX7voLbvk7uEWm9ET3-GcST8-rPjpVxAjhniNNBCHRKSdx9EvRGKZmHp66jB"
}

resource "jsc_blockpage" "myblockpage" {
  description = "I am the new description for the block page"
  title = "I am a new title here"
}



```
