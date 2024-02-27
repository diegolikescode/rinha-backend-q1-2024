ALTER DATABASE rinha SET TIMEZONE TO 'America/Sao_Paulo';

CREATE TABLE public.clientes (
    id INTEGER PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    limite INT NOT NULL,
    saldo INT NOT NULL
);

CREATE TABLE public.transacoes (
    id SERIAL PRIMARY KEY,
    id_cliente INTEGER REFERENCES public.clientes(id),
    valor INTEGER,
    tipo CHAR(1) NOT NULL,
    descricao VARCHAR(10),
    realizada_em TIMESTAMP DEFAULT now()
);

INSERT INTO public.clientes (id, nome, limite, saldo) VALUES
    (1, 'o barato sai caro', 1000 * 100, 0),
    (2, 'zan corp ltda', 800 * 100, 0),
    (3, 'les cruders', 10000 * 100, 0),
    (4, 'padaria joia de cocaia', 100000 * 100, 0),
    (5, 'kid mais', 5000 * 100, 0);

CREATE OR REPLACE FUNCTION public.inserir_credito(id_cliente INT, valor INT, descricao VARCHAR)
RETURNS TABLE(novo_saldo INT, cliente_limite INT) AS $$
DECLARE
    novo_saldo INT;
    cliente_limite INT;
BEGIN
    IF NOT EXISTS (SELECT 1 FROM public.clientes WHERE id = id_cliente) THEN
        RAISE EXCEPTION 'NOUSER';
    END IF;
    
    -- lock exclusivo do cliente
    PERFORM pg_advisory_xact_lock(id_cliente);
    
    INSERT INTO public.transacoes (id_cliente, valor, descricao, tipo)
    VALUES (id_cliente, valor, descricao, 'c');
    
    UPDATE public.clientes
    SET saldo = saldo + valor
    WHERE id = id_cliente
    RETURNING saldo, limite INTO novo_saldo, cliente_limite;

    RETURN QUERY SELECT novo_saldo, cliente_limite;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION inserir_debito(id_cliente INT, valor INT, descricao VARCHAR)
RETURNS TABLE(novo_saldo INT, cliente_limite INT) AS $$
DECLARE
    novo_saldo INT;
    cliente_limite INT;
BEGIN
    IF NOT EXISTS (SELECT 1 FROM public.clientes WHERE id = id_cliente) THEN
        RAISE EXCEPTION 'NOUSER';
    END IF;
    
    PERFORM pg_advisory_xact_lock(id_cliente);

    IF NOT EXISTS (
        SELECT 1 FROM public.clientes
        WHERE id = id_cliente AND saldo - valor >= -limite
    ) THEN
        RAISE EXCEPTION 'NOLIMIT';
    END IF;

    INSERT INTO public.transacoes (id_cliente, valor, descricao, tipo)
    VALUES (id_cliente, valor, descricao, 'd');

    UPDATE public.clientes
    SET saldo = saldo - valor
    WHERE id = id_cliente
    RETURNING saldo, limite INTO novo_saldo, cliente_limite;

    RETURN QUERY SELECT novo_saldo, cliente_limite;
END;
$$ LANGUAGE plpgsql;

CREATE TYPE transacao_info AS (
    valor INT, 
    tipo CHAR, 
    descricao VARCHAR, 
    realizada_em TIMESTAMP, 
    saldo INT, 
    limite INT
);

CREATE OR REPLACE FUNCTION obter_ultimas_transacoes(var_id_cliente INT)
RETURNS SETOF transacao_info AS $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM public.clientes WHERE id = var_id_cliente) THEN
        RAISE EXCEPTION 'NOUSER';
    END IF;

    PERFORM pg_advisory_xact_lock(var_id_cliente);

    RETURN QUERY
    SELECT t.valor as valor, 
        t.tipo as tipo,
        t.descricao as descricao, 
        t.realizada_em as realizada_em, 
        c.saldo as saldo, 
        c.limite as limite
    FROM public.transacoes t
    JOIN public.clientes c ON t.id_cliente = c.id
    WHERE t.id_cliente = var_id_cliente
    ORDER BY t.id DESC
    LIMIT 10;
END;
$$ LANGUAGE plpgsql;

CREATE INDEX idx_transacoes_id_cliente ON transacoes(id_cliente);
CREATE INDEX idx_transacoes_realizadas_em ON transacoes(realizada_em);

