-- Adicionar coluna criado_por em usuarios
ALTER TABLE usuarios ADD COLUMN criado_por INTEGER;
-- Não faz UPDATE, deixa NULL para o primeiro usuário
ALTER TABLE usuarios ADD CONSTRAINT fk_usuarios_criado_por FOREIGN KEY (criado_por) REFERENCES usuarios(id);

-- Adicionar coluna criado_por em clientes
ALTER TABLE clientes ADD COLUMN criado_por INTEGER;
UPDATE clientes SET criado_por = 1; -- Atribuir ID 1 (admin) para registros já existentes
ALTER TABLE clientes ALTER COLUMN criado_por SET NOT NULL;
ALTER TABLE clientes ADD CONSTRAINT fk_clientes_criado_por FOREIGN KEY (criado_por) REFERENCES usuarios(id);

-- Adicionar coluna criado_por em estoque
ALTER TABLE estoque ADD COLUMN criado_por INTEGER;
UPDATE estoque SET criado_por = 1;
ALTER TABLE estoque ALTER COLUMN criado_por SET NOT NULL;
ALTER TABLE estoque ADD CONSTRAINT fk_estoque_criado_por FOREIGN KEY (criado_por) REFERENCES usuarios(id);

-- Adicionar coluna criado_por em precos
ALTER TABLE precos ADD COLUMN criado_por INTEGER;
UPDATE precos SET criado_por = 1;
ALTER TABLE precos ALTER COLUMN criado_por SET NOT NULL;
ALTER TABLE precos ADD CONSTRAINT fk_precos_criado_por FOREIGN KEY (criado_por) REFERENCES usuarios(id);

-- Adicionar coluna criado_por em vendas
ALTER TABLE vendas ADD COLUMN criado_por INTEGER;
UPDATE vendas SET criado_por = 1;
ALTER TABLE vendas ALTER COLUMN criado_por SET NOT NULL;
ALTER TABLE vendas ADD CONSTRAINT fk_vendas_criado_por FOREIGN KEY (criado_por) REFERENCES usuarios(id);

-- Adicionar coluna criado_por em pagamentos
ALTER TABLE pagamentos ADD COLUMN criado_por INTEGER;
UPDATE pagamentos SET criado_por = 1;
ALTER TABLE pagamentos ALTER COLUMN criado_por SET NOT NULL;
ALTER TABLE pagamentos ADD CONSTRAINT fk_pagamentos_criado_por FOREIGN KEY (criado_por) REFERENCES usuarios(id);
