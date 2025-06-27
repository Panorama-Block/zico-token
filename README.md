# ZicoAI Token 
Um DApp completo com token ERC20, sistema de staking, rifas e transferências cross-chain usando Chainlink.

![Arquitetura Zico Token](assets/zicotoken.png)

## 🏗️ Arquitetura do Sistema

### Visão Geral da Arquitetura

O Zico Token DApp é um ecossistema Web3 completo que combina múltiplas tecnologias blockchain avançadas para criar um token utilitário com funcionalidades cross-chain, sistema de staking com recompensas aleatórias e rifas descentralizadas.

### Componentes Principais

#### 1. **ZicoToken.sol** - Contrato Principal
O contrato central que herda de múltiplas bases do OpenZeppelin e Chainlink:

```solidity
contract ZicoToken is ERC20, Ownable, CCIPReceiver, VRFConsumerBaseV2
```

**Funcionalidades:**
- **ERC20 Standard**: Token fungível com 1M de supply inicial
- **Sistema de Staking**: Usuários podem fazer stake de tokens para receber recompensas
- **Cross-Chain Bridge**: Integração com Chainlink CCIP para transferências entre blockchains
- **Recompensas Aleatórias**: Sistema de loteria usando Chainlink VRF
- **Governança**: Controle de acesso com OpenZeppelin Ownable

**Arquitetura de Staking:**
- Mapeamento de stakes por usuário: `mapping(address => uint256) public stakes`
- Lista de stakers ativa: `address[] public stakerList`
- Sistema de recompensas proporcional baseado no stake
- Recompensas aleatórias via VRF para incentivar participação

#### 2. **ZicoRaffle.sol** - Sistema de Rifas
Contrato especializado para rifas descentralizadas:

```solidity
contract ZicoRaffle is VRFConsumerBaseV2, Ownable
```

**Funcionalidades:**
- **Gestão de Participantes**: Lista dinâmica de endereços elegíveis
- **Seleção Aleatória**: Integração com Chainlink VRF para fairness
- **Prêmios Customizáveis**: Valores configuráveis pelo administrador
- **Integração Treasury**: Prêmios pagos automaticamente do cofre

#### 3. **TreasuryVault.sol** - Gestão de Fundos
Cofre centralizado para gestão de recursos do protocolo:

```solidity
contract TreasuryVault is Ownable
```

**Funcionalidades:**
- **Multi-Token Support**: Gestão de ZICO e LINK tokens
- **Sistema de Taxas**: Coleta automática de fees do protocolo
- **Reward Distribution**: Interface para distribuição de recompensas
- **Controle Administrativo**: Funções de saque e gestão limitadas ao owner

#### 4. **ZicoStaking.sol** - Sistema de Staking Separado
Contrato adicional para staking com diferentes mecânicas:

```solidity
contract ZICOStaking is ERC20, Ownable
```

**Funcionalidades:**
- **Pool de Staking**: Sistema independente de rewards
- **Cálculo de Recompensas**: Distribuição proporcional automática
- **Claim System**: Reivindicação manual de recompensas

### Integração Chainlink

#### **CCIP (Cross-Chain Interoperability Protocol)**
- **Burn & Mint**: Tokens são queimados na chain origem e mintados na chain destino
- **Fee Management**: Pagamento automático em LINK para transações cross-chain
- **Multi-Chain Support**: Suporte para Ethereum, Arbitrum, Polygon, Avalanche
- **Security**: Validação criptográfica e finality garantida

#### **VRF (Verifiable Random Function)**
- **True Randomness**: Números verdadeiramente aleatórios para rifas e recompensas
- **Provable Fairness**: Verificação on-chain da aleatoriedade
- **Callback System**: Processamento assíncrono de resultados
- **Gas Optimization**: Callbacks configuráveis para eficiência

### Arquitetura Frontend

#### **React Application Stack**
```
Frontend/
├── Components/          # Componentes reutilizáveis
├── Hooks/              # Lógica de integração Web3
├── Utils/              # Helpers e constantes
└── Services/           # Integração com contratos
```

**Tecnologias:**
- **Ethers.js**: Biblioteca principal para interação blockchain
- **TailwindCSS**: Framework CSS para UI moderna
- **React Hooks**: Gerenciamento de estado reativo
- **Toast Notifications**: Feedback visual para transações

#### **Fluxo de Interação Web3**
1. **Wallet Connection**: Detecção e conexão automática com MetaMask
2. **Contract Instances**: Inicialização de contratos com providers
3. **Transaction Handling**: Gestão de estados de transação (pending, success, error)
4. **Real-time Updates**: Polling e eventos para atualizações em tempo real

### Segurança e Governança

#### **Controles de Acesso**
- **Ownable Pattern**: Funções administrativas protegidas
- **Role-Based Access**: Diferentes níveis de permissão
- **Emergency Functions**: Pausas e withdrawals de emergência

#### **Validações On-Chain**
- **Input Validation**: Verificação de parâmetros de entrada
- **Balance Checks**: Validação de saldos antes de operações
- **Reentrancy Protection**: Uso de padrões seguros do OpenZeppelin

#### **Audit Trail**
- **Event Logging**: Todos os eventos importantes são emitidos
- **Transparent Operations**: Operações verificáveis on-chain
- **Immutable History**: Histórico permanente de transações

### Casos de Uso do Protocolo

#### **Para Usuários Regulares**
1. **Staking Passivo**: Depositar tokens para recompensas regulares
2. **Cross-Chain Trading**: Mover tokens entre diferentes blockchains
3. **Participação em Rifas**: Chances de ganhar prêmios especiais
4. **Yield Farming**: Maximizar retornos através do staking

#### **Para Administradores**
1. **Liquidity Management**: Gestão de liquidez cross-chain
2. **Reward Distribution**: Distribuição estratégica de incentivos
3. **Protocol Governance**: Decisões sobre parâmetros do protocolo
4. **Event Management**: Criação e gestão de rifas especiais

### Escalabilidade e Performance

#### **Gas Optimization**
- **Batch Operations**: Múltiplas operações em uma transação
- **State Packing**: Otimização de storage slots
- **View Functions**: Operações read-only para economia de gas

#### **Multi-Chain Architecture**
- **Layer 2 Support**: Compatibilidade com rollups (Arbitrum, Polygon)
- **Cross-Chain Composability**: Interação entre diferentes protocolos
- **Unified Experience**: Interface única para múltiplas chains

## Estrutura do Projeto

```
zico-token/
├── contracts/              # Smart contracts (Foundry)
│   ├── src/               # Contratos Solidity
│   ├── script/            # Scripts de deploy
│   ├── test/              # Testes unitários
│   └── foundry.toml       # Configuração Foundry
├── frontend/              # Interface React
│   ├── src/               # Código fonte React
│   ├── public/            # Arquivos públicos
│   └── package.json       # Dependências NPM
├── assets/                # Recursos visuais e documentação
├── .github/               # GitHub Actions CI/CD
├── setup.sh               # Script de configuração inicial
├── run_all.sh             # Script para executar tudo
└── start_frontend.sh      # Script para iniciar frontend
```

## Funcionalidades

- **Token ERC20**: Token ZICO com funcionalidades avançadas
- **Sistema de Staking**: Stake tokens e receba recompensas aleatórias
- **Rifas**: Sistema de rifas usando Chainlink VRF
- **Cross-Chain**: Transferências entre blockchains via Chainlink CCIP
- **Interface Web**: Frontend React moderno e responsivo

## Início Rápido

### 1. Configuração Inicial
```bash
./setup.sh
```

### 2. Executar Tudo (Recomendado)
```bash
./run_all.sh
```

### 3. Ou Executar Manualmente

#### Terminal 1 - Blockchain Local
```bash
anvil
```

#### Terminal 2 - Deploy dos Contratos
```bash
cd contracts
forge script script/Deploy.s.sol --rpc-url http://127.0.0.1:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast
```

#### Terminal 3 - Frontend
```bash
./start_frontend.sh
```

## Desenvolvimento

### Contratos
```bash
cd contracts
forge build                # Compilar
forge test                 # Testes
forge fmt                  # Formatação
```

### Frontend
```bash
cd frontend
npm install                # Instalar dependências
npm start                  # Servidor de desenvolvimento
npm run build              # Build de produção
```

## Configuração MetaMask

- **Network**: Anvil Local
- **RPC URL**: http://127.0.0.1:8545
- **Chain ID**: 31337
- **Currency**: ETH
- **Private Key**: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

## Tecnologias

- **Solidity**: Smart contracts
- **Foundry**: Framework de desenvolvimento
- **React**: Interface do usuário
- **Ethers.js**: Interação com blockchain
- **Chainlink VRF**: Números aleatórios
- **Chainlink CCIP**: Transferências cross-chain
- **TailwindCSS**: Estilização
