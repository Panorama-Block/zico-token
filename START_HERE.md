# 🚀 Zico Token DApp - Start Guide

## 📋 Como Rodar o Projeto

### 1️⃣ Primeiro Terminal - Blockchain Local
```bash
anvil
```

### 2️⃣ Segundo Terminal - Deploy do Contrato
```bash
# Na pasta raiz do projeto
forge script script/Deploy.s.sol --rpc-url http://127.0.0.1:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast
```

**⚠️ IMPORTANTE**: Copie o endereço do contrato e atualize `CONTRACT_ADDRESS` em `frontend/src/App.js`

### 3️⃣ Terceiro Terminal - Frontend
```bash
cd frontend
npm start
```

## 🔧 Configurar MetaMask

### Adicionar Rede Anvil
- **Network Name**: Anvil Local
- **RPC URL**: http://127.0.0.1:8545  
- **Chain ID**: 31337
- **Currency Symbol**: ETH

### Importar Conta de Teste
- **Private Key**: `0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80`

## 🎯 Funcionalidades Disponíveis

### 💰 Token Features
- ✅ Transfer tokens
- ✅ Check balance
- ✅ ERC20 standard compliance

### 🔒 Staking System  
- ✅ Stake tokens
- ✅ Unstake tokens
- ✅ Earn rewards proportionally
- ✅ Claim accumulated rewards

### 🌉 Cross-Chain (CCIP)
- ✅ Send tokens to other chains
- ✅ Avalanche, Polygon, Arbitrum support
- ⚠️ Requires CCIP setup (mock for local testing)

### 🎲 Random Lottery (VRF)
- ✅ Admin can start lottery
- ✅ Random winner selection
- ⚠️ Requires VRF subscription (mock for local testing)

### ⚡ Admin Functions (Owner Only)
- ✅ Distribute rewards to all stakers
- ✅ Start random lottery
- ✅ Set remote chain contracts

## 🎨 Interface Preview

### Dashboard Cards
- 💰 **Balance**: Current ZICO tokens
- 📊 **Staked**: Tokens currently staked  
- 🎁 **Rewards**: Available rewards to claim
- 🔄 **Total Staked**: Total tokens staked in system

### Action Sections
- 💸 **Transfer**: Send tokens to any address
- 🔒 **Staking**: Stake/unstake and claim rewards
- 🌉 **Cross-Chain**: Transfer between blockchains
- ⚡ **Admin Panel**: Owner-only functions

## 🔍 Testing Flow

1. **Connect Wallet** → MetaMask popup
2. **Check Balance** → Should show 1,000,000 ZICO
3. **Test Transfer** → Send tokens to another address
4. **Try Staking** → Stake some tokens
5. **Distribute Rewards** → (Admin) Give rewards to stakers
6. **Claim Rewards** → Claim your earned rewards
7. **Test Lottery** → (Admin) Start random lottery

## 🚨 Common Issues

- **Contract not found**: Update CONTRACT_ADDRESS in App.js
- **Wrong network**: Switch MetaMask to Anvil Local (31337)
- **No funds**: Import the test account with ETH
- **CCIP/VRF failing**: Expected in local environment

---

**Ready to go? Start with terminal 1! 🚀** 