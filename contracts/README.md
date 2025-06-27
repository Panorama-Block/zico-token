# Zico Token Smart Contracts

Esta pasta contém todos os contratos inteligentes e arquivos relacionados ao desenvolvimento com Foundry.

## Estrutura

```
contracts/
├── src/                    # Contratos inteligentes
│   ├── ZicoToken.sol      # Contrato principal do token
│   ├── ZicoRaffle.sol     # Sistema de rifas
│   ├── TreasuryVault.sol  # Cofre do tesouro
│   └── ZicoStaking.sol    # Sistema de staking
├── script/                # Scripts de deploy
│   └── Deploy.s.sol       # Script de deploy principal
├── test/                  # Testes unitários
│   └── ZicoToken.t.sol    # Testes do token
├── lib/                   # Dependências (forge-std)
├── foundry.toml           # Configuração do Foundry
└── remappings.txt         # Mapeamentos de importação
```

## Comandos

### Build
```bash
cd contracts
forge build
```

### Testes
```bash
cd contracts
forge test
```

### Deploy Local (Anvil)
```bash
cd contracts
forge script script/Deploy.s.sol --rpc-url http://127.0.0.1:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast
```

### Formatação
```bash
cd contracts
forge fmt
```

## Contratos Principais

### ZicoToken
- Token ERC20 com funcionalidades de staking
- Integração com Chainlink VRF para recompensas aleatórias
- Suporte a CCIP para transferências cross-chain

### ZicoRaffle
- Sistema de rifas usando Chainlink VRF
- Integração com o token ZICO

### TreasuryVault
- Cofre para gerenciar fundos do tesouro
- Controle de acesso para operações financeiras 