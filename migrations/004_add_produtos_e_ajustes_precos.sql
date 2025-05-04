-- DROP das tabelas antigas para recriação
DROP TABLE IF EXISTS venda_pagamentos, venda_itens, vendas, tabela_precos, produtos CASCADE;

-- Tabela de produtos
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    unidade VARCHAR(20) NOT NULL,
    ativo BOOLEAN DEFAULT TRUE,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Inserir produtos padrão
INSERT INTO produtos (nome, unidade) VALUES
('Casada', 'kg'),
('Dianteiro', 'kg'),
('Traseiro', 'kg'),
('Víscera', 'un'),
('Panelada', 'kg'),
('Outros', 'kg');

-- Tabela de preços
CREATE TABLE tabela_precos (
    id SERIAL PRIMARY KEY,
    produto_id INT REFERENCES produtos(id),
    preco NUMERIC(10,2) NOT NULL,
    tipo_animal VARCHAR(10) NOT NULL, -- 'boi' ou 'vaca'
    inicio_vigencia DATE NOT NULL,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Preços iniciais
INSERT INTO tabela_precos (produto_id, tipo_animal, preco, inicio_vigencia) VALUES
(1, 'boi', 10.00, '2025-04-10'),
(1, 'vaca', 9.00, '2025-04-10'),
(2, 'boi', 7.00, '2025-04-10'),
(2, 'vaca', 6.00, '2025-04-10'),
(3, 'boi', 15.00, '2025-04-10'),
(3, 'vaca', 14.00, '2025-04-10');

-- Tabela de vendas
CREATE TABLE vendas (
    id SERIAL PRIMARY KEY,
    cliente_id INT REFERENCES clientes(id),
    usuario_id INT REFERENCES usuarios(id),
    data DATE NOT NULL,
    numero_nota INT,
    total_final NUMERIC(10,2) NOT NULL,
    desconto NUMERIC(10,2) DEFAULT 0,
    acrescimo NUMERIC(10,2) DEFAULT 0,
    status_pagamento VARCHAR(20) DEFAULT 'nao_pago',
    data_quitacao DATE,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    criado_por INT NOT NULL REFERENCES usuarios(id)
);

-- Itens da venda
CREATE TABLE venda_itens (
    id SERIAL PRIMARY KEY,
    venda_id INT REFERENCES vendas(id),
    produto_id INT REFERENCES produtos(id),
    tipo_animal VARCHAR(10) NOT NULL,
    peso_kg NUMERIC(10,2),
    quantidade INT,
    preco_unitario NUMERIC(10,2),
    preco_total NUMERIC(10,2),
    lado VARCHAR(10),
    numero_animal INT
);

-- Pagamentos de venda
CREATE TABLE venda_pagamentos (
    id SERIAL PRIMARY KEY,
    venda_id INT REFERENCES vendas(id),
    valor_pagamento NUMERIC(10,2) NOT NULL,
    forma_pagamento VARCHAR(20) NOT NULL,
    data_pagamento DATE NOT NULL,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    criado_por INT NOT NULL REFERENCES usuarios(id)
    observacao TEXT,
);
