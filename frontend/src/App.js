import React, { useState, useEffect, useCallback } from 'react';
import { ethers } from 'ethers';
import toast, { Toaster } from 'react-hot-toast';
// import { Card, CardContent, CardDescription, CardHeader, CardTitle } from './components/ui/card';
import { Button } from './components/ui/button';
import { Input } from './components/ui/input';

// Import assets
import bridgeIcon from './assets/bridge.png';
import stakingIcon from './assets/staking.png';
import swapIcon from './assets/swap.png';
// import panoramaLogo from './assets/logo.png';
import panoramaLogoBig from './assets/logo-grande.png';

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
  const [activeTab, setActiveTab] = useState('transfer');
  
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
        loading: 'Requesting random reward...',
        success: 'Random reward requested!',
        error: 'Request failed!'
      });
      await tx.wait();
      loadData();
      setLotteryAmount('');
    } catch (error) {
      toast.error('Error: ' + error.message);
    }
  };

  const formatNumber = (num) => {
    const number = parseFloat(num);
    if (number >= 1000000) {
      return (number / 1000000).toFixed(2) + 'M';
    } else if (number >= 1000) {
      return (number / 1000).toFixed(2) + 'K';
    }
    return number.toFixed(4);
  };

  const chainNames = {
    '43114': 'Avalanche',
    '42161': 'Arbitrum',
    '137': 'Polygon'
  };

  const tabs = [
    { id: 'transfer', label: 'Transfer', icon: swapIcon },
    { id: 'stake', label: 'Stake', icon: stakingIcon },
    { id: 'bridge', label: 'Bridge', icon: bridgeIcon },
    { id: 'admin', label: 'Admin', icon: null }
  ];

  return (
    <div className="min-h-screen bg-black text-white relative overflow-hidden">
      {/* Background Elements */}
      <div className="grid-background"></div>
      <div className="panorama-bg">
        <img src={panoramaLogoBig} alt="Zico Token" />
      </div>

      {/* Header */}
      <header className="relative z-10 border-b border-cyan/20 bg-black/50 backdrop-blur-lg">
        <div className="container mx-auto px-6 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-3">
              <img src={panoramaLogoBig} alt="Zico Token" className="h-12 w-8" />
              <h1 className="text-2xl font-bold text-cyan">
  
              </h1>
            </div>
            
            {!account ? (
              <Button
                onClick={connectWallet}
                disabled={isConnecting}
                className="cyber-button px-8 py-3 rounded-lg font-semibold"
              >
                {isConnecting ? 'Connecting...' : 'Connect Wallet'}
              </Button>
            ) : (
              <div className="flex items-center space-x-4">
                <div className="text-sm text-cyan">
                  {account.substring(0, 6)}...{account.substring(38)}
                </div>
                <div className="h-3 w-3 bg-cyan rounded-full pulse-cyan"></div>
              </div>
            )}
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="relative z-10 container mx-auto px-6 py-8">
        {/* Stats Dashboard */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          <div className="stat-card">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-gray-400">Balance</p>
                <p className="text-2xl font-bold text-cyan">{account ? formatNumber(balance) : '0.00'}</p>
                <p className="text-xs text-gray-500">ZICO</p>
              </div>
              <div className="h-12 w-12 bg-cyan/10 rounded-xl flex items-center justify-center">
                <div className="h-6 w-6 bg-cyan rounded-full glow-animation"></div>
              </div>
            </div>
          </div>

          <div className="stat-card">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-gray-400">Staked</p>
                <p className="text-2xl font-bold text-cyan">{account ? formatNumber(stakedAmount) : '0.00'}</p>
                <p className="text-xs text-gray-500">ZICO</p>
              </div>
              <div className="h-12 w-12 bg-cyan/10 rounded-xl flex items-center justify-center">
                <img src={stakingIcon} alt="Staking" className="h-6 w-6" />
              </div>
            </div>
          </div>

          <div className="stat-card">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-gray-400">Rewards</p>
                <p className="text-2xl font-bold text-cyan">{account ? formatNumber(rewards) : '0.00'}</p>
                <p className="text-xs text-gray-500">ZICO</p>
              </div>
              <div className="h-12 w-12 bg-cyan/10 rounded-xl flex items-center justify-center">
                <div className="h-6 w-6 bg-gradient-cyber rounded-full"></div>
              </div>
            </div>
          </div>

          <div className="stat-card">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-gray-400">Total Staked</p>
                <p className="text-2xl font-bold text-cyan">{account ? formatNumber(totalStaked) : '0.00'}</p>
                <p className="text-xs text-gray-500">ZICO</p>
              </div>
              <div className="h-12 w-12 bg-cyan/10 rounded-xl flex items-center justify-center">
                <div className="h-6 w-6 bg-cyan/50 rounded-full pulse-cyan"></div>
              </div>
            </div>
          </div>
        </div>

        {/* Navigation Tabs */}
        <div className="flex flex-wrap gap-4 mb-8">
          {tabs.map((tab) => (
            <button
              key={tab.id}
              onClick={() => setActiveTab(tab.id)}
              className={`tab-button flex items-center space-x-2 ${
                activeTab === tab.id ? 'active' : ''
              }`}
              disabled={!account && tab.id !== 'transfer'}
            >
              {tab.icon && <img src={tab.icon} alt={tab.label} className="h-4 w-4" />}
              <span>{tab.label}</span>
            </button>
          ))}
        </div>

        {/* Tab Content */}
        <div className="glass-card p-8 rounded-xl">
          {!account && (
            <div className="text-center py-8 mb-6">
              <p className="text-gray-400 mb-4">Connect your wallet to access all features</p>
            </div>
          )}

          {activeTab === 'transfer' && (
            <div className="space-y-6">
              <div className="flex items-center space-x-3 mb-6">
                <img src={swapIcon} alt="Transfer" className="h-8 w-8" />
                <h2 className="text-2xl font-bold text-cyan">Transfer Tokens</h2>
              </div>
              
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div className="space-y-4">
                  <div>
                    <label className="block text-sm font-medium text-gray-300 mb-2">
                      Recipient Address
                    </label>
                    <Input
                      type="text"
                      placeholder="0x..."
                      value={transferTo}
                      onChange={(e) => setTransferTo(e.target.value)}
                      className="cyber-input w-full"
                      disabled={!account}
                    />
                  </div>
                  
                  <div>
                    <label className="block text-sm font-medium text-gray-300 mb-2">
                      Amount
                    </label>
                    <Input
                      type="number"
                      placeholder="0.0"
                      value={transferAmount}
                      onChange={(e) => setTransferAmount(e.target.value)}
                      className="cyber-input w-full"
                      disabled={!account}
                    />
                  </div>
                  
                  <Button
                    onClick={handleTransfer}
                    className="cyber-button w-full py-3"
                    disabled={!account || !transferTo || !transferAmount}
                  >
                    Transfer Tokens
                  </Button>
                </div>
                
                <div className="bg-black/30 rounded-lg p-6 border border-cyan/20">
                  <h3 className="text-lg font-semibold text-cyan mb-4">Transfer Info</h3>
                  <div className="space-y-2 text-sm">
                    <div className="flex justify-between">
                      <span className="text-gray-400">Available Balance:</span>
                      <span className="text-white">{account ? formatNumber(balance) : '0.00'} ZICO</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="text-gray-400">Network:</span>
                      <span className="text-white">Local Testnet</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          )}

          {activeTab === 'stake' && (
            <div className="space-y-6">
              <div className="flex items-center space-x-3 mb-6">
                <img src={stakingIcon} alt="Staking" className="h-8 w-8" />
                <h2 className="text-2xl font-bold text-cyan">Staking Portal</h2>
              </div>
              
              <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
                {/* Stake */}
                <div className="space-y-4">
                  <h3 className="text-lg font-semibold text-cyan">Stake Tokens</h3>
                  <Input
                    type="number"
                    placeholder="Amount to stake"
                    value={stakeAmount}
                    onChange={(e) => setStakeAmount(e.target.value)}
                    className="cyber-input w-full"
                    disabled={!account}
                  />
                  <Button
                    onClick={handleStake}
                    className="cyber-button w-full"
                    disabled={!account || !stakeAmount}
                  >
                    Stake
                  </Button>
                </div>

                {/* Unstake */}
                <div className="space-y-4">
                  <h3 className="text-lg font-semibold text-cyan">Unstake Tokens</h3>
                  <Input
                    type="number"
                    placeholder="Amount to unstake"
                    value={unstakeAmount}
                    onChange={(e) => setUnstakeAmount(e.target.value)}
                    className="cyber-input w-full"
                    disabled={!account}
                  />
                  <Button
                    onClick={handleUnstake}
                    className="cyber-button w-full"
                    disabled={!account || !unstakeAmount}
                  >
                    Unstake
                  </Button>
                </div>

                {/* Rewards */}
                <div className="space-y-4">
                  <h3 className="text-lg font-semibold text-cyan">Claim Rewards</h3>
                  <div className="bg-black/30 rounded-lg p-4 border border-cyan/20">
                    <div className="text-center">
                      <p className="text-2xl font-bold text-cyan">{account ? formatNumber(rewards) : '0.00'}</p>
                      <p className="text-sm text-gray-400">Available Rewards</p>
                    </div>
                  </div>
                  <Button
                    onClick={handleClaimRewards}
                    className="cyber-button w-full"
                    disabled={!account || parseFloat(rewards) === 0}
                  >
                    Claim Rewards
                  </Button>
                </div>
              </div>
            </div>
          )}

          {activeTab === 'bridge' && (
            <div className="space-y-6">
              <div className="flex items-center space-x-3 mb-6">
                <img src={bridgeIcon} alt="Bridge" className="h-8 w-8" />
                <h2 className="text-2xl font-bold text-cyan">Cross-Chain Bridge</h2>
              </div>
              
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div className="space-y-4">
                  <div>
                    <label className="block text-sm font-medium text-gray-300 mb-2">
                      Destination Chain
                    </label>
                    <select
                      value={crossChainDestination}
                      onChange={(e) => setCrossChainDestination(e.target.value)}
                      className="cyber-input w-full"
                      disabled={!account}
                    >
                      <option value="43114">Avalanche</option>
                      <option value="42161">Arbitrum</option>
                      <option value="137">Polygon</option>
                    </select>
                  </div>
                  
                  <div>
                    <label className="block text-sm font-medium text-gray-300 mb-2">
                      Amount
                    </label>
                    <Input
                      type="number"
                      placeholder="0.0"
                      value={crossChainAmount}
                      onChange={(e) => setCrossChainAmount(e.target.value)}
                      className="cyber-input w-full"
                      disabled={!account}
                    />
                  </div>
                  
                  <Button
                    onClick={handleCrossChain}
                    className="cyber-button w-full py-3"
                    disabled={!account || !crossChainAmount}
                  >
                    Bridge to {chainNames[crossChainDestination]}
                  </Button>
                </div>
                
                <div className="bg-black/30 rounded-lg p-6 border border-cyan/20">
                  <h3 className="text-lg font-semibold text-cyan mb-4">Bridge Info</h3>
                  <div className="space-y-3 text-sm">
                    <div className="flex justify-between">
                      <span className="text-gray-400">From:</span>
                      <span className="text-white">Local Testnet</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="text-gray-400">To:</span>
                      <span className="text-white">{chainNames[crossChainDestination]}</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="text-gray-400">Estimated Time:</span>
                      <span className="text-white">~5-10 minutes</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          )}

          {activeTab === 'admin' && account && isOwner && (
            <div className="space-y-6">
              <h2 className="text-2xl font-bold text-cyan">Admin Functions</h2>
              
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div className="space-y-4">
                  <h3 className="text-lg font-semibold text-cyan">Reward Distribution</h3>
                  <Button
                    onClick={handleDistributeRewards}
                    className="cyber-button w-full"
                    disabled={!account}
                  >
                    Distribute Rewards
                  </Button>
                </div>

                <div className="space-y-4">
                  <h3 className="text-lg font-semibold text-cyan">Random Lottery</h3>
                  <Input
                    type="number"
                    placeholder="Reward amount"
                    value={lotteryAmount}
                    onChange={(e) => setLotteryAmount(e.target.value)}
                    className="cyber-input w-full"
                    disabled={!account}
                  />
                  <Button
                    onClick={handleLottery}
                    className="cyber-button w-full"
                    disabled={!account || !lotteryAmount}
                  >
                    Start Random Lottery
                  </Button>
                </div>
              </div>
            </div>
          )}

          {activeTab === 'admin' && account && !isOwner && (
            <div className="text-center py-8">
              <p className="text-gray-400">Admin access required</p>
            </div>
          )}

          {activeTab === 'admin' && !account && (
            <div className="text-center py-8">
              <p className="text-gray-400">Connect wallet to access admin functions</p>
            </div>
          )}
        </div>
      </main>

      <Toaster
        position="bottom-right"
        toastOptions={{
          style: {
            background: 'rgba(26, 26, 26, 0.9)',
            color: '#ffffff',
            border: '1px solid rgba(0, 255, 255, 0.3)',
            borderRadius: '8px',
          },
        }}
      />
    </div>
  );
}

export default App; 