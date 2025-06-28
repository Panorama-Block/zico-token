import React, { useState, useEffect, useCallback } from 'react';
import { ethers } from 'ethers';
import toast, { Toaster } from 'react-hot-toast';
import { Button } from './components/ui/button';

import ZICO_TOKEN_ABI from './abis/ZicoToken.json';
import ZICO_STAKING_ABI from './abis/ZicoStaking.json';
import ZICO_RAFFLE_ABI from './abis/ZicoRaffle.json';
import TREASURY_VAULT_ABI from './abis/TreasuryVault.json';

const CONTRACT_ADDRESSES = {
  token: '0xYourZicoTokenAddress',
  staking: '0xYourZicoStakingAddress',
  raffle: '0xYourZicoRaffleAddress',
  vault: '0xYourTreasuryVaultAddress'
};

function App() {
  const [provider, setProvider] = useState(null);
  const [signer, setSigner] = useState(null);
  const [account, setAccount] = useState('');

  const [tokenContract, setTokenContract] = useState(null);
  const [stakingContract, setStakingContract] = useState(null);
  const [raffleContract, setRaffleContract] = useState(null);
  const [vaultContract, setVaultContract] = useState(null);

  const [balance, setBalance] = useState('0');
  const [staked, setStaked] = useState('0');
  const [rewards, setRewards] = useState('0');
  const [totalStaked, setTotalStaked] = useState('0');

  const connectWallet = async () => {
    if (!window.ethereum) return toast.error('Please install MetaMask');
    const _provider = new ethers.providers.Web3Provider(window.ethereum);
    await _provider.send('eth_requestAccounts', []);
    const _signer = _provider.getSigner();
    const _account = await _signer.getAddress();

    setProvider(_provider);
    setSigner(_signer);
    setAccount(_account);

    setTokenContract(new ethers.Contract(CONTRACT_ADDRESSES.token, ZICO_TOKEN_ABI, _signer));
    setStakingContract(new ethers.Contract(CONTRACT_ADDRESSES.staking, ZICO_STAKING_ABI, _signer));
    setRaffleContract(new ethers.Contract(CONTRACT_ADDRESSES.raffle, ZICO_RAFFLE_ABI, _signer));
    setVaultContract(new ethers.Contract(CONTRACT_ADDRESSES.vault, TREASURY_VAULT_ABI, _signer));

    toast.success('Wallet connected!');
  };

  const loadData = useCallback(async () => {
    if (!tokenContract || !stakingContract || !account) return;
    try {
      const [bal, stk, rew, totStk] = await Promise.all([
        tokenContract.balanceOf(account),
        stakingContract.getStakedAmount(account),
        stakingContract.getReward(account),
        stakingContract.totalStaked()
      ]);

      setBalance(ethers.utils.formatEther(bal));
      setStaked(ethers.utils.formatEther(stk));
      setRewards(ethers.utils.formatEther(rew));
      setTotalStaked(ethers.utils.formatEther(totStk));
    } catch (err) {
      toast.error('Error loading data');
      console.error(err);
    }
  }, [tokenContract, stakingContract, account]);

  useEffect(() => {
    loadData();
  }, [loadData]);

  const handleStake = async (amount) => {
    try {
      await tokenContract.approve(CONTRACT_ADDRESSES.staking, ethers.utils.parseEther(amount));
      const tx = await stakingContract.stake(ethers.utils.parseEther(amount));
      toast.promise(tx.wait(), {
        loading: 'Staking...',
        success: 'Staked!',
        error: 'Stake failed'
      });
      await tx.wait();
      loadData();
    } catch (err) {
      toast.error(err.message);
    }
  };

  const handleUnstake = async (amount) => {
    try {
      const tx = await stakingContract.unstake(ethers.utils.parseEther(amount));
      toast.promise(tx.wait(), {
        loading: 'Unstaking...',
        success: 'Unstaked!',
        error: 'Unstake failed'
      });
      await tx.wait();
      loadData();
    } catch (err) {
      toast.error(err.message);
    }
  };

  const handleClaim = async () => {
    try {
      const tx = await stakingContract.claimReward();
      toast.promise(tx.wait(), {
        loading: 'Claiming...',
        success: 'Claimed!',
        error: 'Claim failed'
      });
      await tx.wait();
      loadData();
    } catch (err) {
      toast.error(err.message);
    }
  };

  const handleDistribute = async () => {
    try {
      const tx = await vaultContract.distributeRewards();
      toast.promise(tx.wait(), {
        loading: 'Distributing...',
        success: 'Distributed!',
        error: 'Distribute failed'
      });
      await tx.wait();
    } catch (err) {
      toast.error(err.message);
    }
  };

  const handleLottery = async (amount) => {
    try {
      const tx = await raffleContract.requestRandomReward(ethers.utils.parseEther(amount));
      toast.promise(tx.wait(), {
        loading: 'Requesting...',
        success: 'Requested!',
        error: 'Failed'
      });
      await tx.wait();
    } catch (err) {
      toast.error(err.message);
    }
  };

  return (
    <div className="p-4">
      <Toaster />
      {!account ? (
        <Button onClick={connectWallet}>Connect Wallet</Button>
      ) : (
        <div>
          <div>Balance: {balance} ZICO</div>
          <div>Staked: {staked}</div>
          <div>Rewards: {rewards}</div>
          <div>Total Staked: {totalStaked}</div>
        </div>
      )}
    </div>
  );
}

export default App;
