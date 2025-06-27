import React, { useState, useEffect, useCallback } from 'react';
import { ethers } from 'ethers';
import toast, { Toaster } from 'react-hot-toast';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from './components/ui/card';
import { Button } from './components/ui/button';
import { Input } from './components/ui/input';

// ABI do contrato ZicoToken
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
  const [contract, setContract] = useState(null);
  const [account, setAccount] = useState('');
  const [balance, setBalance] = useState('0');
  const [stakedAmount, setStakedAmount] = useState('0');
  const [rewards, setRewards] = useState('0');
  const [totalStaked, setTotalStaked] = useState('0');
  const [isOwner, setIsOwner] = useState(false);
  const [isConnecting, setIsConnecting] = useState(false);
  const [activeTab, setActiveTab] = useState('swap');
  
  // Form states
  const [transferTo, setTransferTo] = useState('');
  const [transferAmount, setTransferAmount] = useState('');
  const [stakeAmount, setStakeAmount] = useState('');
  const [unstakeAmount, setUnstakeAmount] = useState('');
  const [crossChainDestination, setCrossChainDestination] = useState('43114');
  const [crossChainAmount, setCrossChainAmount] = useState('');
  const [lotteryAmount, setLotteryAmount] = useState('');

  // Contract address deployado
  const CONTRACT_ADDRESS = "0x5FbDB2315678afecb367f032d93F642f64180aa3";

  const loadData = useCallback(async () => {
    if (!contract || !account) return;
    
    try {
      const [balance, staked, userRewards, totalStaked, owner] = await Promise.all([
        contract.balanceOf(account),
        contract.stakes(account),
        contract.rewards(account),
        contract.totalStaked(),
        contract.owner()
      ]);
      
      setBalance(ethers.utils.formatEther(balance));
      setStakedAmount(ethers.utils.formatEther(staked));
      setRewards(ethers.utils.formatEther(userRewards));
      setTotalStaked(ethers.utils.formatEther(totalStaked));
      setIsOwner(account.toLowerCase() === owner.toLowerCase());
    } catch (error) {
      console.error('Error loading data:', error);
    }
  }, [contract, account]);

  useEffect(() => {
    loadData();
  }, [loadData]);

  const connectWallet = async () => {
    try {
      setIsConnecting(true);
      if (window.ethereum) {
        const provider = new ethers.providers.Web3Provider(window.ethereum);
        await provider.send("eth_requestAccounts", []);
        const signer = provider.getSigner();
        const account = await signer.getAddress();
        
        setAccount(account);
        
        const contract = new ethers.Contract(CONTRACT_ADDRESS, ZICO_TOKEN_ABI, signer);
        setContract(contract);
        
        toast.success('Wallet connected successfully!');
      } else {
        toast.error('Please install MetaMask!');
      }
    } catch (error) {
      toast.error('Error connecting wallet: ' + error.message);
    } finally {
      setIsConnecting(false);
    }
  };

  const handleTransfer = async () => {
    if (!transferTo || !transferAmount) return;
    try {
      const tx = await contract.transfer(transferTo, ethers.utils.parseEther(transferAmount));
      toast.promise(tx.wait(), {
        loading: 'Processing transfer...',
        success: 'Transfer completed!',
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
    if (!stakeAmount) return;
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
    if (!unstakeAmount) return;
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
    if (!crossChainAmount) return;
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
    if (!lotteryAmount) return;
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

  const formatNumber = (num) => {
    const number = parseFloat(num);
    if (number === 0) return '0';
    if (number < 0.01) return '<0.01';
    if (number >= 1000000) return (number / 1000000).toFixed(2) + 'M';
    if (number >= 1000) return (number / 1000).toFixed(2) + 'K';
    return number.toFixed(4);
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50 to-purple-50">
      <Toaster position="top-right" />
      
      {/* Header */}
      <div className="bg-white/80 backdrop-blur-md border-b border-white/20 sticky top-0 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
          <div className="flex justify-between items-center">
            <div className="flex items-center space-x-4">
              <div className="w-10 h-10 rounded-full bg-gradient-to-r from-purple-600 to-pink-600 flex items-center justify-center">
                <span className="text-white font-bold text-lg">Z</span>
              </div>
              <h1 className="text-2xl font-bold bg-gradient-to-r from-purple-600 to-pink-600 bg-clip-text text-transparent">
                Zico AI Token
              </h1>
            </div>
            
            {account ? (
              <div className="flex items-center space-x-4">
                <div className="hidden sm:flex items-center space-x-4 text-sm">
                  <span className="text-gray-600">{formatNumber(balance)} ZICO</span>
                  <span className="w-2 h-2 rounded-full bg-green-500"></span>
                </div>
                <div className="px-3 py-2 bg-gray-100 rounded-lg text-sm font-medium">
                  {account.slice(0, 6)}...{account.slice(-4)}
                </div>
              </div>
            ) : (
              <Button 
                onClick={connectWallet} 
                disabled={isConnecting}
                className="px-6 py-2"
              >
                {isConnecting ? 'Connecting...' : 'Connect Wallet'}
              </Button>
            )}
          </div>
        </div>
      </div>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Stats Overview */}
        <div className="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
          <Card className="bg-white/60 backdrop-blur-md border-white/20">
            <CardHeader className="pb-2">
              <CardDescription className="text-gray-600">Your Balance</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">{formatNumber(balance)}</div>
              <div className="text-sm text-gray-500">ZICO</div>
            </CardContent>
          </Card>
          
          <Card className="bg-white/60 backdrop-blur-md border-white/20">
            <CardHeader className="pb-2">
              <CardDescription className="text-gray-600">Staked</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold text-green-600">{formatNumber(stakedAmount)}</div>
              <div className="text-sm text-gray-500">ZICO</div>
            </CardContent>
          </Card>
          
          <Card className="bg-white/60 backdrop-blur-md border-white/20">
            <CardHeader className="pb-2">
              <CardDescription className="text-gray-600">Rewards</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold text-purple-600">{formatNumber(rewards)}</div>
              <div className="text-sm text-gray-500">ZICO</div>
            </CardContent>
          </Card>
          
          <Card className="bg-white/60 backdrop-blur-md border-white/20">
            <CardHeader className="pb-2">
              <CardDescription className="text-gray-600">Total Staked</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold text-blue-600">{formatNumber(totalStaked)}</div>
              <div className="text-sm text-gray-500">ZICO</div>
            </CardContent>
          </Card>
        </div>

        {/* Main Content */}
        <div className="max-w-2xl mx-auto">
          {/* Tab Navigation */}
          <div className="flex space-x-1 bg-white/60 backdrop-blur-md p-1 rounded-xl mb-6 border border-white/20">
            {[
              { id: 'swap', label: 'Transfer', icon: '💸' },
              { id: 'stake', label: 'Stake', icon: '🔒' },
              { id: 'bridge', label: 'Bridge', icon: '🌉' },
              ...(isOwner ? [{ id: 'admin', label: 'Admin', icon: '⚡' }] : [])
            ].map((tab) => (
              <button
                key={tab.id}
                onClick={() => setActiveTab(tab.id)}
                className={`flex-1 px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200 ${
                  activeTab === tab.id
                    ? 'bg-white shadow-md text-purple-600'
                    : 'text-gray-600 hover:text-gray-800'
                }`}
              >
                <span className="mr-2">{tab.icon}</span>
                {tab.label}
              </button>
            ))}
          </div>

          {/* Tab Content */}
          <Card className="bg-white/80 backdrop-blur-md border-white/30 shadow-xl">
            {activeTab === 'swap' && (
              <div>
                <CardHeader>
                  <CardTitle className="flex items-center">
                    <span className="mr-2">💸</span>
                    Transfer Tokens
                  </CardTitle>
                  <CardDescription>Send ZICO tokens to any address</CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Recipient Address
                    </label>
                    <Input
                      placeholder="0x..."
                      value={transferTo}
                      onChange={(e) => setTransferTo(e.target.value)}
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Amount
                    </label>
                    <div className="relative">
                      <Input
                        type="number"
                        placeholder="0.0"
                        value={transferAmount}
                        onChange={(e) => setTransferAmount(e.target.value)}
                        className="pr-16"
                      />
                      <div className="absolute right-3 top-3 text-sm font-medium text-gray-500">
                        ZICO
                      </div>
                    </div>
                  </div>
                  <Button
                    onClick={handleTransfer}
                    disabled={!transferTo || !transferAmount || !account}
                    className="w-full h-12 text-lg"
                  >
                    Transfer
                  </Button>
                </CardContent>
              </div>
            )}

            {activeTab === 'stake' && (
              <div>
                <CardHeader>
                  <CardTitle className="flex items-center">
                    <span className="mr-2">🔒</span>
                    Staking Pool
                  </CardTitle>
                  <CardDescription>Stake ZICO tokens to earn rewards</CardDescription>
                </CardHeader>
                <CardContent className="space-y-6">
                  <div className="grid grid-cols-2 gap-4">
                    <div>
                      <label className="block text-sm font-medium text-gray-700 mb-2">
                        Stake Amount
                      </label>
                      <div className="space-y-2">
                        <Input
                          type="number"
                          placeholder="0.0"
                          value={stakeAmount}
                          onChange={(e) => setStakeAmount(e.target.value)}
                        />
                        <Button
                          onClick={handleStake}
                          disabled={!stakeAmount || !account}
                          variant="secondary"
                          className="w-full"
                        >
                          Stake
                        </Button>
                      </div>
                    </div>
                    <div>
                      <label className="block text-sm font-medium text-gray-700 mb-2">
                        Unstake Amount
                      </label>
                      <div className="space-y-2">
                        <Input
                          type="number"
                          placeholder="0.0"
                          value={unstakeAmount}
                          onChange={(e) => setUnstakeAmount(e.target.value)}
                        />
                        <Button
                          onClick={handleUnstake}
                          disabled={!unstakeAmount || !account}
                          variant="outline"
                          className="w-full"
                        >
                          Unstake
                        </Button>
                      </div>
                    </div>
                  </div>
                  
                  <div className="pt-4 border-t">
                    <div className="flex justify-between items-center mb-4">
                      <span className="text-sm font-medium text-gray-700">
                        Claimable Rewards
                      </span>
                      <span className="text-lg font-bold text-purple-600">
                        {formatNumber(rewards)} ZICO
                      </span>
                    </div>
                    <Button
                      onClick={handleClaimRewards}
                      disabled={parseFloat(rewards) === 0 || !account}
                      className="w-full h-12"
                    >
                      Claim Rewards
                    </Button>
                  </div>
                </CardContent>
              </div>
            )}

            {activeTab === 'bridge' && (
              <div>
                <CardHeader>
                  <CardTitle className="flex items-center">
                    <span className="mr-2">🌉</span>
                    Cross-Chain Bridge
                  </CardTitle>
                  <CardDescription>Transfer tokens across different blockchains</CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Destination Chain
                    </label>
                    <select
                      value={crossChainDestination}
                      onChange={(e) => setCrossChainDestination(e.target.value)}
                      className="w-full h-12 px-4 bg-white border border-gray-200 rounded-xl focus:border-purple-500 focus:ring-2 focus:ring-purple-500/20"
                    >
                      <option value="43114">🔺 Avalanche</option>
                      <option value="137">🟣 Polygon</option>
                      <option value="42161">🔷 Arbitrum</option>
                    </select>
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Amount
                    </label>
                    <Input
                      type="number"
                      placeholder="0.0"
                      value={crossChainAmount}
                      onChange={(e) => setCrossChainAmount(e.target.value)}
                    />
                  </div>
                  <Button
                    onClick={handleCrossChain}
                    disabled={!crossChainAmount || !account}
                    className="w-full h-12 text-lg"
                  >
                    Bridge Tokens
                  </Button>
                </CardContent>
              </div>
            )}

            {activeTab === 'admin' && isOwner && (
              <div>
                <CardHeader>
                  <CardTitle className="flex items-center">
                    <span className="mr-2">⚡</span>
                    Admin Functions
                  </CardTitle>
                  <CardDescription>Manage rewards and lottery system</CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  <Button
                    onClick={handleDistributeRewards}
                    variant="secondary"
                    className="w-full h-12"
                  >
                    📊 Distribute Rewards to All Stakers
                  </Button>
                  
                  <div className="space-y-2">
                    <label className="block text-sm font-medium text-gray-700">
                      Lottery Reward Amount
                    </label>
                    <div className="flex space-x-2">
                      <Input
                        type="number"
                        placeholder="Reward amount"
                        value={lotteryAmount}
                        onChange={(e) => setLotteryAmount(e.target.value)}
                        className="flex-1"
                      />
                      <Button
                        onClick={handleLottery}
                        disabled={!lotteryAmount}
                        className="px-6"
                      >
                        🎲 Start Lottery
                      </Button>
                    </div>
                  </div>
                </CardContent>
              </div>
            )}
          </Card>

          {/* Instructions for new users */}
          {!account && (
            <Card className="mt-6 bg-gradient-to-r from-blue-50 to-purple-50 border-blue-200">
              <CardHeader>
                <CardTitle className="text-blue-800">🚀 Getting Started</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-2 text-blue-700">
                  <p>• Connect your MetaMask wallet</p>
                  <p>• Make sure you're on Anvil Local network (Chain ID: 31337)</p>
                  <p>• Import the test account for instant access to ZICO tokens</p>
                </div>
              </CardContent>
            </Card>
          )}
        </div>
      </div>
    </div>
  );
}

export default App; 