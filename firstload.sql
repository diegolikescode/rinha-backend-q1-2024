/*
type Cliente struct {
    ID      int32 `json:"id"`
    Nome    string `json:"nome"`
    Limite  int32 `json:"limite"`
}

type Transacao struct {
    Valor      int32 `json:"valor"`
    Tipo       string `json:"tipo"`
    Descricao  string `json:"descricao"`
    RealizadaEm string `json:"realizada_em"`
}

type Conta struct {
    Limite   int32 `json:"limite"`
    Saldo    int32 `json:"saldo"`
}

type Saldo struct {
    Total         int32 `json:"total"`
    DataExtrato   int32 `json:"data_extrato"`
    Limite	  int32 `json:"limite"`
}

type Extrato struct {
    Saldo	        Saldo `json:"saldo"`
    UltimasTransacoes   []Transacao `json:"ultimas_transacoes"`
}
*/

 /* cliente | conta | saldo */
CREATE TABLE clientes (
    id INTEGER PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    limite INT NOT NULL,
    saldo INT NOT NULL
);

CREATE TABLE transacoes (
    id SERIAL PRIMARY KEY,
    id_cliente INTEGER REFERENCES clientes(id),
    valor INTEGER,
    tipo VARCHAR(5) NOT NULL,
    descricao VARCHAR(50),
    realizada_em VARCHAR(50)
);

INSERT INTO clientes (id, nome, limite, saldo) VALUES
    (1, 'o barato sai caro', 1000 * 100, 0),
    (2, 'zan corp ltda', 800 * 100, 0),
    (3, 'les cruders', 10000 * 100, 0),
    (4, 'padaria joia de cocaia', 100000 * 100, 0),
    (5, 'kid mais', 5000 * 100, 0);
