#!/bin/bash

echo "🚀 Setting up Zico Token DApp..."

# Criar arquivo .env
echo "📝 Creating .env file..."
cat > .env << EOF
PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
RPC_URL=http://127.0.0.1:8545
EOF

echo "✅ Environment file created"

# Build do contrato
echo "🔨 Building smart contract..."
cd contracts
forge build
cd ..

# Setup do frontend
echo "🌐 Setting up frontend..."
cd frontend
npm install
cd ..

echo "📋 Setup complete! Now run the following commands in separate terminals:"
echo ""
echo "Terminal 1 - Start local blockchain:"
echo "anvil"
echo ""
echo "Terminal 2 - Deploy contract:"
echo "cd contracts && forge script script/Deploy.s.sol --rpc-url http://127.0.0.1:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast"
echo ""
echo "Terminal 3 - Start frontend:"
echo "cd frontend && npm start"
echo ""
echo "🔥 Don't forget to:"
echo "1. Copy the deployed contract address from the deployment log"
echo "2. Update CONTRACT_ADDRESS in frontend/src/App.js"
echo "3. Add the Anvil network to MetaMask (Chain ID: 31337, RPC: http://127.0.0.1:8545)"
echo "4. Import the test account private key to MetaMask" 