package startup

import (
	"github.com/cjlapao/common-go/service_provider"
	"github.com/cjlapao/rediscli-go/controller"
)

var providers = service_provider.Get()

func Init() {
	// ctx := execution_context.Get()
	// ctx.WithDefaultAuthorization()
	// dbFactory, databaseName := databaseservice.GetDatabase()
	// ctx.Authorization.WithAudience("carloslapao.com")
	// kv := ctx.Authorization.KeyVault
	// kv.WithBase64RsaKey("RSA256_2048", providers.Configuration.GetString("JWT_RSA256_2048_PRIVATE_KEY"))
	// kv.WithBase64RsaKey("RSA256_4096", providers.Configuration.GetString("JWT_RSA256_4096_PRIVATE_KEY"))
	// kv.WithBase64RsaKey("RSA512_4096", providers.Configuration.GetString("JWT_RSA512_4096_PRIVATE_KEY"))
	// kv.WithBase64EcdsaKey("ECDSA256", providers.Configuration.GetString("JWT_ECDSA256_PRIVATE_KEY"))
	// kv.WithBase64HmacKey("HMAC", providers.Configuration.GetString("JWT_HMAC_PRIVATE_KEY"), encryption.Bit256)
	// kv.SetDefaultKey("RSA512_4096")
	// identity.Seed(dbFactory, databaseName)

	controller.Init()
}
