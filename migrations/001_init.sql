-- Criação de tabelas para o CarneGestão

CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    senha_hash TEXT NOT NULL,
    tipo VARCHAR(20) NOT NULL CHECK (tipo IN ('admin', 'funcionario', 'cliente')),
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    nome_proprietario VARCHAR(100) NOT NULL,
    nome_comercial VARCHAR(100) NOT NULL,
    telefone VARCHAR(20),
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE vendas (
    id SERIAL PRIMARY KEY,
    cliente_id INTEGER REFERENCES clientes(id),
    usuario_id INTEGER REFERENCES usuarios(id),
    data DATE NOT NULL,
    numero_nota INTEGER NOT NULL,
    total_final NUMERIC(10,2) NOT NULL,
    desconto NUMERIC(10,2) DEFAULT 0,
    acrescimo NUMERIC(10,2) DEFAULT 0,
    status_pagamento VARCHAR(20) NOT NULL CHECK (status_pagamento IN ('pago', 'pago_parcialmente', 'nao_pago')),
    data_quitacao DATE,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (cliente_id, numero_nota)
);

CREATE TABLE venda_itens (
    id SERIAL PRIMARY KEY,
    venda_id INTEGER REFERENCES vendas(id),
    tipo_carne VARCHAR(20) NOT NULL CHECK (tipo_carne IN ('casada', 'dianteiro', 'traseiro', 'viscera', 'panelada')),
    animal VARCHAR(10) NOT NULL CHECK (animal IN ('boi', 'vaca')),
    peso_kg NUMERIC(10,2),
    quantidade INTEGER,
    preco_unitario NUMERIC(10,2) NOT NULL,
    preco_total NUMERIC(10,2) NOT NULL,
    lado VARCHAR(10),
    numero_animal INTEGER
);

CREATE TABLE estoque (
    id SERIAL PRIMARY KEY,
    data DATE NOT NULL,
    quantidade_bois INTEGER DEFAULT 0,
    quantidade_vacas INTEGER DEFAULT 0,
    peso_total_bois NUMERIC(10,2),
    peso_total_vacas NUMERIC(10,2),
    preco_kg_boi NUMERIC(10,2),
    preco_kg_vaca NUMERIC(10,2),
    fornecedor VARCHAR(100),
    visceras_recebidas INTEGER DEFAULT 0,
    visceras_condenadas INTEGER DEFAULT 0,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tabela_precos (
    id SERIAL PRIMARY KEY,
    tipo_carne VARCHAR(20) NOT NULL CHECK (tipo_carne IN ('casada', 'dianteiro', 'traseiro', 'viscera', 'panelada')),
    animal VARCHAR(10) NOT NULL CHECK (animal IN ('boi', 'vaca')),
    preco NUMERIC(10,2) NOT NULL,
    data_vigencia DATE NOT NULL,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE venda_pagamentos (
    id SERIAL PRIMARY KEY,
    venda_id INTEGER REFERENCES vendas(id),
    data_pagamento DATE NOT NULL,
    valor_pagamento NUMERIC(10,2) NOT NULL,
    forma_pagamento VARCHAR(20) NOT NULL CHECK (forma_pagamento IN ('DINHEIRO', 'PIX', 'TRANSFERENCIA', 'CREDITO', 'DEBITO', 'OUTROS')),
    observacao TEXT,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
