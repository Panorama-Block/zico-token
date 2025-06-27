# 🪙 Zico Token DApp

Uma DApp completa para interagir com o ZicoToken, featuring:
- 💰 Transferências ERC20
- 🔒 Sistema de Staking com recompensas
- 🌉 Cross-Chain transfers via Chainlink CCIP
- 🎲 Loteria aleatória usando Chainlink VRF

## 🚀 Quick Start

### 1. Setup Inicial
```bash
# Tornar o script executável e rodar
chmod +x setup.sh
./setup.sh
```

### 2. Iniciar Blockchain Local
```bash
# Terminal 1
anvil
```

### 3. Deploy do Contrato
```bash
# Terminal 2
forge script script/Deploy.s.sol --rpc-url http://127.0.0.1:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast
```

**⚠️ IMPORTANTE**: Copie o endereço do contrato deployado e atualize `CONTRACT_ADDRESS` em `frontend/src/App.js`

### 4. Configurar MetaMask
1. Adicionar rede Anvil:
   - Network Name: `Anvil Local`
   - RPC URL: `http://127.0.0.1:8545`
   - Chain ID: `31337`
   - Currency Symbol: `ETH`

2. Importar conta de teste:
   - Private Key: `0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80`

### 5. Iniciar Frontend
```bash
# Terminal 3
cd frontend
npm start
```

## 🎯 Funcionalidades

### 💸 Transfer Tokens
- Transferir ZICO para qualquer endereço
- Validação de saldo em tempo real

### 🔒 Staking System
- **Stake**: Depositar tokens para ganhar recompensas
- **Unstake**: Retirar tokens do staking
- **Claim Rewards**: Reivindicar recompensas acumuladas

### 🌉 Cross-Chain Bridge
- Enviar tokens entre diferentes blockchains
- Suporte para Avalanche, Polygon, Arbitrum
- Powered by Chainlink CCIP

### 🎲 Random Lottery (Admin Only)
- Sistema de loteria para stakers
- Seleção aleatória via Chainlink VRF
- Apenas o owner pode iniciar

### ⚡ Admin Functions
- **Distribute Rewards**: Distribuir recompensas proporcionais para todos os stakers
- **Start Lottery**: Iniciar sorteio com prêmio customizado

## 📊 Interface

### Dashboard Principal
- **Balance**: Saldo atual de ZICO tokens
- **Staked**: Quantidade de tokens em staking
- **Rewards**: Recompensas disponíveis para claim
- **Total Staked**: Total de tokens em staking no sistema

### Seções Interativas
1. **Transfer Tokens**: Interface para envio de tokens
2. **Staking**: Gerenciamento de stake/unstake
3. **Cross-Chain**: Transferências entre blockchains
4. **Admin Panel**: Funções administrativas (visível apenas para owner)

## 🛠️ Tecnologias

### Smart Contract
- **Solidity ^0.8.24**
- **Foundry** para desenvolvimento e testes
- **OpenZeppelin** para padrões de segurança
- **Chainlink CCIP** para cross-chain
- **Chainlink VRF** para aleatoriedade

### Frontend
- **React 18**
- **Ethers.js** para interação blockchain
- **TailwindCSS** para estilização
- **React Hot Toast** para notificações

## 🔧 Troubleshooting

### Problemas Comuns

**1. "Insufficient funds" no MetaMask**
- Certifique-se de estar usando a conta com ETH da Anvil
- A primeira conta da Anvil vem com 10,000 ETH

**2. "Contract not deployed"**
- Verifique se o CONTRACT_ADDRESS está correto no App.js
- Certifique-se de que o deploy foi bem-sucedido

**3. "Network mismatch"**
- Confirme que o MetaMask está conectado na rede Anvil (Chain ID: 31337)

**4. "CCIP/VRF functions failing"**
- No ambiente local, as funções Chainlink podem falhar
- Isso é esperado, pois são serviços externos

### Testing Local

Para testar sem Chainlink (apenas funcionalidades básicas):
1. Use apenas Transfer e Staking
2. As funções Cross-Chain e Lottery precisam de mocks para teste local

## 📝 Contratos

### ZicoToken.sol
```solidity
// Contrato principal que implementa:
- ERC20 (OpenZeppelin)
- Ownable (OpenZeppelin) 
- CCIPReceiver (Chainlink)
- VRFConsumerBaseV2 (Chainlink)
```

### Deploy.s.sol
```solidity
// Script de deployment com endereços mock para desenvolvimento local
```

## 🌟 Features Avançadas

### Cross-Chain Mapping
O contrato suporta mapeamento de contratos remotos:
```solidity
mapping(uint64 => address) public remotes;
```

### Sistema de Staking Proporcional
Recompensas calculadas proporcionalmente ao stake:
```solidity
uint256 reward = (stakes[stakerAddr] * rewardPool) / totalStaked;
```

### VRF para Aleatoriedade Verificável
Sistema de loteria com aleatoriedade provável:
```solidity
mapping(uint256 => uint256) public requestIdToReward;
```

## 🤝 Contribuindo

1. Fork o projeto
2. Crie sua feature branch
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

---

**Desenvolvido com ❤️ para demonstrar o poder dos smart contracts multi-chain** 