#!/bin/bash

echo "🚀 Starting Zico Token DApp..."

# Function to kill background processes on script exit
cleanup() {
    echo "🛑 Stopping all processes..."
    pkill -f anvil
    pkill -f "npm start"
    exit 0
}

# Set up cleanup on script exit
trap cleanup EXIT

# Start Anvil in background
echo "🔗 Starting Anvil blockchain..."
anvil &
ANVIL_PID=$!

# Wait for Anvil to start
sleep 5

# Deploy contract
echo "🚀 Deploying ZicoToken contract..."
cd contracts
forge script script/Deploy.s.sol:DeployScript --rpc-url http://127.0.0.1:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast
cd ..

# Check if deployment was successful
if [ $? -eq 0 ]; then
    echo "✅ Contract deployed successfully!"
    
    # Start frontend
    echo "🌐 Starting React frontend..."
    cd frontend
    npm start &
    FRONTEND_PID=$!
    
    echo ""
    echo "🎉 DApp is now running!"
    echo ""
    echo "📋 Setup MetaMask:"
    echo "   - Network: Anvil Local"
    echo "   - RPC URL: http://127.0.0.1:8545"
    echo "   - Chain ID: 31337"
    echo "   - Currency: ETH"
    echo ""
    echo "🔑 Import Account:"
    echo "   - Private Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
    echo ""
    echo "🌐 Frontend: http://localhost:3000"
    echo ""
    echo "Press Ctrl+C to stop all services"
    
    # Wait for frontend to exit
    wait $FRONTEND_PID
else
    echo "❌ Contract deployment failed!"
    exit 1
fi 