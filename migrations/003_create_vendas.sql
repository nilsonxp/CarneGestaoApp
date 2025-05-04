-- Criação da tabela de produtos
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    unidade VARCHAR(10) NOT NULL,         -- Ex: kg, unid
    ativo BOOLEAN DEFAULT TRUE,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Inserção de produtos padrão
INSERT INTO produtos (nome, unidade) VALUES
('casada_boi', 'kg'),
('casada_vaca', 'kg'),
('dianteiro', 'kg'),
('traseiro', 'kg'),
('viscera', 'unid'),
('panelada', 'kg');

-- Criação da tabela de preços
CREATE TABLE tabela_precos (
    id SERIAL PRIMARY KEY,
    produto_id INT REFERENCES produtos(id),
    tipo_animal VARCHAR(10),              -- 'boi', 'vaca' ou NULL
    preco FLOAT NOT NULL,
    data_inicio DATE NOT NULL,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Seed de preços iniciais (com data fictícia, ajuste se necessário)
INSERT INTO tabela_precos (produto_id, tipo_animal, preco, data_inicio) VALUES
(1, 'boi', 10.00, '2025-04-01'),
(2, 'vaca', 9.00, '2025-04-01'),
(3, NULL, 6.00, '2025-04-01'),
(4, NULL, 8.00, '2025-04-01'),
(5, NULL, 5.00, '2025-04-01'),
(6, NULL, 7.00, '2025-04-01');

-- Criação da tabela de vendas
CREATE TABLE vendas (
    id SERIAL PRIMARY KEY,
    cliente_id INT REFERENCES clientes(id),
    usuario_id INT REFERENCES usuarios(id),
    data DATE NOT NULL,
    numero_nota INT NOT NULL,
    total_final NUMERIC(10, 2) NOT NULL,
    desconto NUMERIC(10, 2) DEFAULT 0,
    acrescimo NUMERIC(10, 2) DEFAULT 0,
    status_pagamento VARCHAR(20) DEFAULT 'nao_pago',
    data_quitacao DATE,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    criado_por INT NOT NULL REFERENCES usuarios(id)
);

-- Alteração na tabela já existente venda_itens
ALTER TABLE venda_itens
    ADD COLUMN preco_padrao NUMERIC(10, 2),
    ADD COLUMN preco_alterado BOOLEAN DEFAULT FALSE;

-- Criação da tabela de pagamentos
CREATE TABLE venda_pagamentos (
    id SERIAL PRIMARY KEY,
    venda_id INT NOT NULL REFERENCES vendas(id) ON DELETE CASCADE,
    forma_pagamento VARCHAR(20) NOT NULL,
    valor NUMERIC(10, 2) NOT NULL,
    data_pagamento TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
