# Introdução

Olá! Este projeto é 100% Opensource e foi desenvolvido online no canal da [Spacedevs](https://twitch.tv/spacedevs) na Twitch.tv. O intuito deste projeto é oferecer as pessoas que usam a linguagem GO, um meio de pagamento rápido e fácil como o Picpay.

> Este projeto não é oficial da Picpay e não fomos patrocinados por ela.

### O que já foi feito?

- [x] Obter status de um pedido
- [x] Realizar pagamentos
- [x] Cancelar pagamento
- [ ] Notificações

## Instalando

Para instalar o pacote você deve usar o comando a seguir:

```bash
go get https://github.com/marcuxyz/golang-picpay
```

## Usando
Após instalar você deve chamar o pacote em seu código e começar a usar as APIS internas. Lembrando que antes de tentar criar um pagamento ou até mesmo obter o status do mesmo, você deve instânciar o objeto usando a função New()

```golang
p := picpay.New("MEU_TOKEN_DO_PICPAY_VAI_AQUI")
```

Agora, você pode usar para obter um status do pedido:

```golang
p.GetOrderStatus("234219")
```

Ou até mesmo você pode usar para realizar um pagamento:

```golang
type Payment struct {
	ReferenceID string  `json:"referenceId"`
	CallbackURL string  `json:"callbackUrl"`
	ReturnURL   string  `json:"returnUrl"`
	Value       float64 `json:"value"`
	ExpiresAt   string  `json:"expiresAt"`
	Buyer       `json:"buyer"`
}

type Buyer struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Document  string `json:"document"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func main() {
  p := picpay.New("MEU_TOKEN_DO_PICPAY_VAI_AQUI"))
  
  payment := Payment{
    ReferenceID: "827371918",
    CallbackURL: "https://spacedevs.com.br/callbackurl",
    ReturnURL:   "https://spacedevs.com.br/",
    Value:       0.10,
    ExpiresAt:   "2020-07-20T11:00:00-03:00",
    Buyer: Buyer{
      FirstName: "Marcus",
      LastName:  "Pereira",
      Document:  "123.456.789-10",
      Email:     "contato@marcuspereira.xyz",
      Phone:     "+55 71 12345-6789",
    },
  }
  result, err := p.PayOrder(payment)
  if err != nil {
    panic(err)
  }
  
  for _, err := range result.Errors {
    fmt.Println(err.Message)
  }
}
```