import React, { useState, useEffect } from 'react';
import { ethers } from 'ethers';
import toast, { Toaster } from 'react-hot-toast';

// ABI do contrato ZicoToken (simplificado para as funções principais)
const ZICO_TOKEN_ABI = [
  "function name() view returns (string)",
  "function symbol() view returns (string)",
  "function totalSupply() view returns (uint256)",
  "function balanceOf(address) view returns (uint256)",
  "function transfer(address to, uint256 amount) returns (bool)",
  "function approve(address spender, uint256 amount) returns (bool)",
  "function allowance(address owner, address spender) view returns (uint256)",
  "function stake(uint256 amount)",
  "function unstake(uint256 amount)",
  "function stakes(address) view returns (uint256)",
  "function rewards(address) view returns (uint256)",
  "function claimReward()",
  "function distributeRewards()",
  "function sendCrossChain(uint64 destChain, uint256 amount)",
  "function requestRandomReward(uint256 rewardAmount)",
  "function totalStaked() view returns (uint256)",
  "function owner() view returns (address)",
  "event Transfer(address indexed from, address indexed to, uint256 value)",
  "event Staked(address indexed user, uint256 amount)",
  "event Unstaked(address indexed user, uint256 amount)",
  "event RewardDistributed(address indexed user, uint256 reward)",
  "event CrossChainSend(uint64 indexed toChain, address indexed to, uint256 amount)"
];

function App() {
  const [provider, setProvider] = useState(null);
  const [signer, setSigner] = useState(null);
  const [contract, setContract] = useState(null);
  const [account, setAccount] = useState('');
  const [balance, setBalance] = useState('0');
  const [stakedAmount, setStakedAmount] = useState('0');
  const [rewards, setRewards] = useState('0');
  const [totalStaked, setTotalStaked] = useState('0');
  const [isOwner, setIsOwner] = useState(false);
  
  // Form states
  const [transferTo, setTransferTo] = useState('');
  const [transferAmount, setTransferAmount] = useState('');
  const [stakeAmount, setStakeAmount] = useState('');
  const [unstakeAmount, setUnstakeAmount] = useState('');
  const [crossChainDestination, setCrossChainDestination] = useState('43114'); // Avalanche
  const [crossChainAmount, setCrossChainAmount] = useState('');
  const [lotteryAmount, setLotteryAmount] = useState('');

  // Contract address deployado
  const CONTRACT_ADDRESS = "0x5FbDB2315678afecb367f032d93F642f64180aa3";

  useEffect(() => {
    connectWallet();
  }, []);

  useEffect(() => {
    if (contract && account) {
      loadData();
    }
  }, [contract, account]);

  const connectWallet = async () => {
    try {
      if (window.ethereum) {
        const provider = new ethers.providers.Web3Provider(window.ethereum);
        await provider.send("eth_requestAccounts", []);
        const signer = provider.getSigner();
        const account = await signer.getAddress();
        
        setProvider(provider);
        setSigner(signer);
        setAccount(account);
        
        // Connect to contract
        const contract = new ethers.Contract(CONTRACT_ADDRESS, ZICO_TOKEN_ABI, signer);
        setContract(contract);
        
        toast.success('Wallet connected!');
      } else {
        toast.error('Please install MetaMask!');
      }
    } catch (error) {
      toast.error('Error connecting wallet: ' + error.message);
    }
  };

  const loadData = async () => {
    try {
      const balance = await contract.balanceOf(account);
      const staked = await contract.stakes(account);
      const userRewards = await contract.rewards(account);
      const totalStaked = await contract.totalStaked();
      const owner = await contract.owner();
      
      setBalance(ethers.utils.formatEther(balance));
      setStakedAmount(ethers.utils.formatEther(staked));
      setRewards(ethers.utils.formatEther(userRewards));
      setTotalStaked(ethers.utils.formatEther(totalStaked));
      setIsOwner(account.toLowerCase() === owner.toLowerCase());
    } catch (error) {
      console.error('Error loading data:', error);
    }
  };

  const handleTransfer = async () => {
    try {
      const tx = await contract.transfer(transferTo, ethers.utils.parseEther(transferAmount));
      toast.promise(tx.wait(), {
        loading: 'Transferring tokens...',
        success: 'Transfer successful!',
        error: 'Transfer failed!'
      });
      await tx.wait();
      loadData();
      setTransferTo('');
      setTransferAmount('');
    } catch (error) {
      toast.error('Error: ' + error.message);
    }
  };

  const handleStake = async () => {
    try {
      const tx = await contract.stake(ethers.utils.parseEther(stakeAmount));
      toast.promise(tx.wait(), {
        loading: 'Staking tokens...',
        success: 'Stake successful!',
        error: 'Stake failed!'
      });
      await tx.wait();
      loadData();
      setStakeAmount('');
    } catch (error) {
      toast.error('Error: ' + error.message);
    }
  };

  const handleUnstake = async () => {
    try {
      const tx = await contract.unstake(ethers.utils.parseEther(unstakeAmount));
      toast.promise(tx.wait(), {
        loading: 'Unstaking tokens...',
        success: 'Unstake successful!',
        error: 'Unstake failed!'
      });
      await tx.wait();
      loadData();
      setUnstakeAmount('');
    } catch (error) {
      toast.error('Error: ' + error.message);
    }
  };

  const handleClaimRewards = async () => {
    try {
      const tx = await contract.claimReward();
      toast.promise(tx.wait(), {
        loading: 'Claiming rewards...',
        success: 'Rewards claimed!',
        error: 'Claim failed!'
      });
      await tx.wait();
      loadData();
    } catch (error) {
      toast.error('Error: ' + error.message);
    }
  };

  const handleDistributeRewards = async () => {
    try {
      const tx = await contract.distributeRewards();
      toast.promise(tx.wait(), {
        loading: 'Distributing rewards...',
        success: 'Rewards distributed!',
        error: 'Distribution failed!'
      });
      await tx.wait();
      loadData();
    } catch (error) {
      toast.error('Error: ' + error.message);
    }
  };

  const handleCrossChain = async () => {
    try {
      const tx = await contract.sendCrossChain(
        crossChainDestination, 
        ethers.utils.parseEther(crossChainAmount)
      );
      toast.promise(tx.wait(), {
        loading: 'Sending cross-chain...',
        success: 'Cross-chain transfer initiated!',
        error: 'Cross-chain failed!'
      });
      await tx.wait();
      loadData();
      setCrossChainAmount('');
    } catch (error) {
      toast.error('Error: ' + error.message);
    }
  };

  const handleLottery = async () => {
    try {
      const tx = await contract.requestRandomReward(ethers.utils.parseEther(lotteryAmount));
      toast.promise(tx.wait(), {
        loading: 'Starting lottery...',
        success: 'Lottery started!',
        error: 'Lottery failed!'
      });
      await tx.wait();
      setLotteryAmount('');
    } catch (error) {
      toast.error('Error: ' + error.message);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <Toaster position="top-right" />
      
      {/* Header */}
      <div className="bg-white shadow-lg">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
          <div className="flex justify-between items-center">
            <h1 className="text-3xl font-bold text-gray-900">🪙 Zico Token DApp</h1>
            {account ? (
              <div className="flex items-center space-x-4">
                <span className="text-sm text-gray-600">
                  {account.slice(0, 6)}...{account.slice(-4)}
                </span>
                <span className="h-6 w-6 text-green-500">💼</span>
              </div>
            ) : (
              <button
                onClick={connectWallet}
                className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
              >
                Connect Wallet
              </button>
            )}
          </div>
        </div>
      </div>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Stats Cards */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center">
              <span className="text-3xl">💰</span>
              <div className="ml-4">
                <p className="text-sm font-medium text-gray-600">Balance</p>
                <p className="text-2xl font-bold text-gray-900">{parseFloat(balance).toFixed(2)} ZICO</p>
              </div>
            </div>
          </div>
          
          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center">
              <span className="text-3xl">📊</span>
              <div className="ml-4">
                <p className="text-sm font-medium text-gray-600">Staked</p>
                <p className="text-2xl font-bold text-gray-900">{parseFloat(stakedAmount).toFixed(2)} ZICO</p>
              </div>
            </div>
          </div>
          
          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center">
              <span className="text-3xl">🎁</span>
              <div className="ml-4">
                <p className="text-sm font-medium text-gray-600">Rewards</p>
                <p className="text-2xl font-bold text-gray-900">{parseFloat(rewards).toFixed(2)} ZICO</p>
              </div>
            </div>
          </div>
          
          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center">
              <span className="text-3xl">🔄</span>
              <div className="ml-4">
                <p className="text-sm font-medium text-gray-600">Total Staked</p>
                <p className="text-2xl font-bold text-gray-900">{parseFloat(totalStaked).toFixed(2)} ZICO</p>
              </div>
            </div>
          </div>
        </div>

        {/* Main Content */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          {/* Transfer Section */}
          <div className="bg-white rounded-lg shadow-md p-6">
            <h2 className="text-xl font-bold text-gray-900 mb-4">💸 Transfer Tokens</h2>
            <div className="space-y-4">
              <input
                type="text"
                placeholder="Recipient Address"
                value={transferTo}
                onChange={(e) => setTransferTo(e.target.value)}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
              <input
                type="number"
                placeholder="Amount"
                value={transferAmount}
                onChange={(e) => setTransferAmount(e.target.value)}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
              <button
                onClick={handleTransfer}
                disabled={!transferTo || !transferAmount}
                className="w-full bg-blue-600 text-white py-2 rounded-md hover:bg-blue-700 disabled:bg-gray-400"
              >
                Transfer
              </button>
            </div>
          </div>

          {/* Staking Section */}
          <div className="bg-white rounded-lg shadow-md p-6">
            <h2 className="text-xl font-bold text-gray-900 mb-4">🔒 Staking</h2>
            <div className="space-y-4">
              <div className="grid grid-cols-2 gap-2">
                <input
                  type="number"
                  placeholder="Stake Amount"
                  value={stakeAmount}
                  onChange={(e) => setStakeAmount(e.target.value)}
                  className="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
                />
                <button
                  onClick={handleStake}
                  disabled={!stakeAmount}
                  className="bg-green-600 text-white py-2 rounded-md hover:bg-green-700 disabled:bg-gray-400"
                >
                  Stake
                </button>
              </div>
              <div className="grid grid-cols-2 gap-2">
                <input
                  type="number"
                  placeholder="Unstake Amount"
                  value={unstakeAmount}
                  onChange={(e) => setUnstakeAmount(e.target.value)}
                  className="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-red-500"
                />
                <button
                  onClick={handleUnstake}
                  disabled={!unstakeAmount}
                  className="bg-red-600 text-white py-2 rounded-md hover:bg-red-700 disabled:bg-gray-400"
                >
                  Unstake
                </button>
              </div>
              <button
                onClick={handleClaimRewards}
                disabled={parseFloat(rewards) === 0}
                className="w-full bg-purple-600 text-white py-2 rounded-md hover:bg-purple-700 disabled:bg-gray-400"
              >
                Claim Rewards ({parseFloat(rewards).toFixed(2)} ZICO)
              </button>
            </div>
          </div>

          {/* Cross-Chain Section */}
          <div className="bg-white rounded-lg shadow-md p-6">
            <h2 className="text-xl font-bold text-gray-900 mb-4">🌉 Cross-Chain Transfer</h2>
            <div className="space-y-4">
              <select
                value={crossChainDestination}
                onChange={(e) => setCrossChainDestination(e.target.value)}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-orange-500"
              >
                <option value="43114">🔺 Avalanche</option>
                <option value="137">🟣 Polygon</option>
                <option value="42161">🔷 Arbitrum</option>
              </select>
              <input
                type="number"
                placeholder="Amount to send"
                value={crossChainAmount}
                onChange={(e) => setCrossChainAmount(e.target.value)}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-orange-500"
              />
              <button
                onClick={handleCrossChain}
                disabled={!crossChainAmount}
                className="w-full bg-orange-600 text-white py-2 rounded-md hover:bg-orange-700 disabled:bg-gray-400"
              >
                Send Cross-Chain
              </button>
            </div>
          </div>

          {/* Admin Section */}
          {isOwner && (
            <div className="bg-white rounded-lg shadow-md p-6">
              <h2 className="text-xl font-bold text-gray-900 mb-4">⚡ Admin Functions</h2>
              <div className="space-y-4">
                <button
                  onClick={handleDistributeRewards}
                  className="w-full bg-indigo-600 text-white py-2 rounded-md hover:bg-indigo-700"
                >
                  📊 Distribute Rewards to All Stakers
                </button>
                <div className="grid grid-cols-2 gap-2">
                  <input
                    type="number"
                    placeholder="Lottery Reward Amount"
                    value={lotteryAmount}
                    onChange={(e) => setLotteryAmount(e.target.value)}
                    className="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-yellow-500"
                  />
                  <button
                    onClick={handleLottery}
                    disabled={!lotteryAmount}
                    className="bg-yellow-600 text-white py-2 rounded-md hover:bg-yellow-700 disabled:bg-gray-400"
                  >
                    🎲 Start Lottery
                  </button>
                </div>
              </div>
            </div>
          )}
        </div>

        {/* Instructions */}
        {!account && (
          <div className="mt-8 bg-yellow-50 border border-yellow-200 rounded-lg p-6">
            <h3 className="text-lg font-bold text-yellow-800 mb-2">🚀 Getting Started</h3>
            <ul className="text-yellow-700 space-y-1">
              <li>1. Connect your MetaMask wallet</li>
              <li>2. Make sure you're connected to the local Anvil network (Chain ID: 31337)</li>
              <li>3. Deploy the contract using Foundry first</li>
              <li>4. Update the CONTRACT_ADDRESS in the code</li>
            </ul>
          </div>
        )}
      </div>
    </div>
  );
}

export default App; 