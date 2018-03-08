package common

import "encoding/base64"

const Base64Encoder = "zxcvbnmqwertyuiopasdfghjklASDFGHJKLZXCVBNMQWERTYUIOP-+1234567890"

var Encoding64 = base64.NewEncoding(Base64Encoder)
