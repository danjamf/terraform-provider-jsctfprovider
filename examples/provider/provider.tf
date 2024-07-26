terraform {
  required_providers {
    jsc = {
      source  = "jsctf"
      version = "1.0.0"
    }

  }
}


provider "jsc" {
  # Configure provider-specific settings if needed
  # Only local email accounts are supported. No SSO or SAML
  # CustomerID is optional
  username   = "wanderauth@jsc.com"
  password   = "passwordhere"
  customerid = "993ae0ee-4bd8-4325-bc5d-1db0ea45b4f6"
}
